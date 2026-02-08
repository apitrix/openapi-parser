# Reference Resolver — Deep Dive

This document describes how `$ref` resolution works end-to-end in the OpenAPI parser, covering the shared `RefResolver`, per-parser resolve walks, circular reference detection, and external file handling.

## Overview

Reference resolution is a **post-parse phase**. After the YAML tree is parsed into typed Go structs, a separate walk resolves every `*Ref` type that has a non-empty `Ref` string and a `nil` `Value`.

```
YAML bytes
   │
   ▼
┌──────────────────┐
│  Parse Phase     │  Parse YAML → populate model structs
│  (parse.go)      │  $ref stored as string, Value = nil
└──────────────────┘
   │
   ▼
┌──────────────────┐
│  Resolve Phase   │  Walk model tree → resolve $ref values
│  (resolve.go)    │  Load targets, parse into Value, detect cycles
└──────────────────┘
   │
   ▼
  Fully resolved model with Circular flags set
```

## Entry Point

Each parser has a top-level `Resolve` function and a `ParseFile` function that triggers the full pipeline:

```go
// parsers/openapi30x/resolve.go
func Resolve(doc *OpenAPI, root *yaml.Node, basePath string) error {
    r := shared.NewRefResolver(basePath, root)
    resolving := make(map[string]bool)
    return resolveDocument(doc, r, resolving)
}

// parsers/openapi30x/parse.go
func ParseFile(filePath string) (*OpenAPI, error) {
    doc, _ := parseOpenAPI(docNode, ctx)   // parse phase
    Resolve(doc, docNode, basePath)         // resolve phase
    return doc, nil
}
```

---

## The `RefResolver` (shared)

Located in `parsers/internal/shared/resolver.go`, the `RefResolver` handles:

### Fields

| Field | Type | Purpose |
|-------|------|---------|
| `BasePath` | `string` | Directory of the root document for relative path resolution |
| `Root` | `*yaml.Node` | YAML tree of the root document |
| `Fs` | `afero.Fs` | Filesystem for reading external files |
| `fileCache` | `map[string]*yaml.Node` | Cached parsed external files |
| `visiting` | `map[string]bool` | Currently-resolving refs (YAML-level cycle detection) |

### `Resolve(ref string) → (*ResolveResult, error)`

This is the core method. It takes a `$ref` string and returns the target `yaml.Node`:

```
Resolve("#/components/schemas/Pet")
   │
   ├─ canonicalize(ref) → "#/components/schemas/Pet"
   │
   ├─ Check visiting map → if present, return Circular=true
   │
   ├─ Mark as visiting, defer cleanup
   │
   ├─ SplitRef(ref) → (filePath="", pointer="/components/schemas/Pet")
   │
   ├─ filePath == "" → use Root as target document
   │  (otherwise: loadFile(filePath) → load, parse, cache external file)
   │
   └─ ResolveJSONPointer(root, "/components/schemas/Pet")
      │
      └─ Walk mapping nodes by key segments → return target node
```

### `SplitRef(ref string) → (filePath, pointer)`

Breaks a `$ref` into its two components:

| Input | `filePath` | `pointer` |
|-------|-----------|-----------|
| `#/components/schemas/Pet` | `""` | `/components/schemas/Pet` |
| `./schemas/pet.yaml` | `./schemas/pet.yaml` | `""` |
| `./common.yaml#/definitions/Error` | `./common.yaml` | `/definitions/Error` |

### `ResolveJSONPointer(root, pointer)`

Implements [RFC 6901](https://datatracker.ietf.org/doc/html/rfc6901). Walks the YAML tree by splitting the pointer into `/`-separated segments, handling:
- **Mapping nodes** — key lookup
- **Sequence nodes** — integer index
- **RFC 6901 escaping** — `~0` → `~`, `~1` → `/`
- **URL encoding** — `%2F` → `/`

### `loadFile(filePath)`

Loads and caches external YAML/JSON files:

1. Resolve relative path against `BasePath`
2. Check `fileCache` — return cached if present
3. Read file via `afero.Fs`
4. Parse YAML, unwrap document node
5. Store in `fileCache`, return

---

## Per-Parser Resolve Walk

Each parser version has a `resolve.go` that walks the parsed model tree. The walk is structured as a set of mutually recursive functions:

```
resolveDocument
├── resolveComponents
│   ├── resolveSchemaRef     (for each schema)
│   ├── resolveResponseRef   (for each response)
│   ├── resolveParameterRef  (for each parameter)
│   ├── resolveExampleRef    (for each example)
│   ├── resolveRequestBodyRef
│   ├── resolveHeaderRef
│   ├── resolveSecuritySchemeRef
│   ├── resolveLinkRef
│   └── resolveCallbackRef
├── resolvePathItem          (for each path)
│   └── resolveOperation     (for each HTTP method)
│       ├── resolveParameterRef
│       ├── resolveRequestBodyRef
│       └── resolveResponseRef
└── ...
```

### Individual Ref Resolver Pattern

Every `resolve*Ref` function follows the same pattern:

```go
func resolveSchemaRef(ref *SchemaRef, r *RefResolver, resolving map[string]bool) error {
    // 1. Skip if nil, no ref, or already resolved
    if ref == nil || ref.Ref == "" || ref.Value != nil {
        return nil
    }

    // 2. Check model-level cycle detection
    if resolving[ref.Ref] {
        ref.Circular = true
        return nil
    }

    // 3. Resolve at YAML level
    result, err := r.Resolve(ref.Ref)
    if err != nil { return err }

    // 4. Check YAML-level circular detection
    if result.Circular {
        ref.Circular = true
        return nil
    }

    // 5. Parse the resolved YAML node into a typed Value
    ref.Value, err = parseSchema(result.Node, ctx)
    if err != nil { return err }

    // 6. Mark as resolving and walk children
    resolving[ref.Ref] = true
    defer delete(resolving, ref.Ref)
    return resolveSchema(ref.Value, r, resolving)
}
```

---

## Circular Reference Detection

Circular references are detected at two complementary levels to prevent infinite recursion.

### Level 1: YAML-Level (`RefResolver.visiting`)

The `visiting` map in `RefResolver` tracks which canonical refs are currently being resolved within a single `Resolve()` call chain. This uses `defer delete` for stack-like cleanup:

```go
func (r *RefResolver) Resolve(ref string) (*ResolveResult, error) {
    canonicalRef := r.canonicalize(ref)
    if r.visiting[canonicalRef] {
        return &ResolveResult{Circular: true}, nil  // cycle!
    }
    r.visiting[canonicalRef] = true
    defer func() { delete(r.visiting, canonicalRef) }()
    // ... resolve ...
}
```

**Scope:** Single `Resolve()` invocation stack. After `Resolve()` returns, the ref is removed from `visiting`, allowing the same ref to be resolved again in a different context.

### Level 2: Model-Level (`resolving` Map)

A `resolving map[string]bool` is threaded through all per-parser resolve functions. It tracks which `$ref` strings are currently being walked in the model tree.

#### Pre-Registration

Before walking the children of a top-level component, its canonical `$ref` path is pre-registered in the `resolving` map. This ensures that a schema's immediate self-reference is caught on the **first encounter**:

```go
// In resolveComponents:
for name, ref := range c.Schemas {
    canonicalRef := "#/components/schemas/" + name
    resolving[canonicalRef] = true         // ← pre-register
    resolveSchemaRef(ref, r, resolving)    // walk children
    delete(resolving, canonicalRef)        // cleanup
}
```

Without pre-registration, a schema like `TreeNode` referencing `#/definitions/TreeNode` in its `items` would not be detected until a second recursive attempt — by which point the parser may have already entered infinite recursion.

#### Why Two Levels?

| Scenario | Caught By |
|----------|-----------|
| `$ref: "#/schemas/A"` → A has `$ref: "#/schemas/A"` | Level 2 (pre-registration) |
| `$ref: "#/schemas/A"` → A → B → A | Level 2 (resolving map) |
| `$ref: "./a.yaml"` → a.yaml has `$ref: "./a.yaml"` | Level 1 (visiting map) |
| `$ref: "./a.yaml#/X"` → X has `$ref: "./b.yaml#/Y"` → Y has `$ref: "./a.yaml#/X"` | Level 1 (visiting map) |

### Circular Flag

When a cycle is detected, the `Circular` field on the `*Ref` type is set to `true`, and the `Value` is left as `nil`. Consumers of the parsed model can check this flag:

```go
if schemaRef.Circular {
    // This is a self-reference; don't traverse Value
}
```

---

## External File Resolution

External `$ref` values (e.g., `./schemas/pet.yaml` or `./common.yaml#/definitions/Error`) are handled transparently:

1. **`SplitRef`** separates the file path from the JSON pointer
2. **`loadFile`** reads the file via the configured `afero.Fs`, parses it, and caches the result
3. **`ResolveJSONPointer`** navigates to the target within the loaded file

### File Caching

Each unique file path is loaded and parsed only once. Subsequent references to the same file use the cached `*yaml.Node`:

```go
type RefResolver struct {
    fileCache map[string]*yaml.Node  // keyed by absolute path
}
```

### Filesystem Abstraction

The `RefResolver` uses `afero.Fs` for all file I/O:

```go
// Production
resolver := shared.NewRefResolver(basePath, root)          // uses afero.OsFs

// Testing
memFs := afero.NewMemMapFs()
afero.WriteFile(memFs, "/pet.yaml", data, 0644)
resolver := shared.NewRefResolverWithFs(basePath, root, memFs)
```

This makes tests fast, deterministic, and parallel-safe — no temp files or OS cleanup needed.

---

## Canonicalization

To compare refs consistently (e.g., `./pet.yaml` vs `/abs/path/pet.yaml`), the resolver canonicalizes all ref strings before cycle-detection checks:

```go
func (r *RefResolver) canonicalize(ref string) string {
    filePath, pointer := SplitRef(ref)
    if filePath == "" {
        return "#" + pointer          // local: "#/components/schemas/Pet"
    }
    absPath := filepath.Join(r.BasePath, filePath)
    absPath = filepath.Clean(absPath)
    return absPath + "#" + pointer    // external: "/abs/path/pet.yaml#/Pet"
}
```

This ensures that `./pet.yaml` and `../dir/pet.yaml` resolve to the same canonical key when they point to the same file.
