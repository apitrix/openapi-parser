# Reference Resolution

This document describes what types of `$ref` references are supported, how to use them through the public API, and how resolution is implemented internally.

## Supported Reference Types

The parser resolves three categories of `$ref` strings:

| Type | Pattern | Example |
|------|---------|---------|
| **Local** (JSON Pointer) | `#/path/to/node` | `#/components/schemas/Pet` |
| **External File** | `./file.yaml` or `./file.yaml#/pointer` | `./models.yaml#/definitions/Tag` |
| **Remote URL** | `https://host/file.yaml` or `https://host/file.yaml#/pointer` | `https://example.com/pet.yaml` |

All three types can include an optional JSON Pointer fragment (`#/...`) to resolve a specific node within the target document. Without a fragment, the entire document root is returned.

### Local References

Local refs resolve within the root document's YAML tree using [RFC 6901](https://datatracker.ietf.org/doc/html/rfc6901) JSON Pointers. They always start with `#`:

```yaml
components:
  schemas:
    Pet:
      type: object
      properties:
        name:
          type: string
    PetList:
      type: array
      items:
        $ref: '#/components/schemas/Pet'   # ← local ref
```

JSON Pointer features supported:
- Mapping node key lookup (`/components/schemas/Pet`)
- Sequence node index access (`/items/0`)
- RFC 6901 escaping: `~0` → `~`, `~1` → `/`
- URL-encoded characters (e.g. `%2F` → `/`)

### External File References

External file refs load YAML/JSON documents from the local filesystem, relative to the root document's directory:

```yaml
components:
  schemas:
    Tag:
      $ref: './models.yaml#/Tag'           # file + pointer
    Error:
      $ref: '../common/error.yaml'         # file only (returns document root)
```

Files are read via `afero.Fs`, enabling in-memory filesystems for testing:

```go
fs := afero.NewMemMapFs()
afero.WriteFile(fs, "/base/models.yaml", data, 0644)
resolver := shared.NewRefResolverWithFs("/base", root, fs)
```

### Remote URL References

Remote refs fetch documents over HTTP/HTTPS:

```yaml
components:
  schemas:
    Pet:
      $ref: 'https://example.com/schemas/pet.yaml'
    Error:
      $ref: 'https://example.com/common.yaml#/definitions/Error'
```

Remote fetching uses a configurable `*http.Client` (default: 30-second timeout):

```go
resolver := shared.NewRefResolver(basePath, root)
resolver.HTTPClient = &http.Client{
    Transport: &authTransport{token: "bearer xyz"},
}
```

### Document Caching

Both local files and remote URLs are cached after first load. Repeated refs to the same source reuse the cached YAML tree:

```
fileCache map[string]*yaml.Node   // keyed by absolute path or full URL
```

---

## Public API

### Entry Points

Each parser version (`openapi20`, `openapi30x`, `openapi31x`) provides three entry points:

| Function | Resolves Refs? | Description |
|----------|:-:|-------------|
| `Parse(data []byte, cfg...)` | ❌ | Parse only — no resolution, no blocking |
| `ParseReader(r io.Reader, cfg...)` | ❌ | Same as Parse, from an `io.Reader` |
| `ParseFile(pathOrURL string, cfg...)` | ✅ | Parse + background resolve from file or URL |

`ParseFile` auto-detects whether the input is a local file path or an HTTP/HTTPS URL.

### ParseConfig

Resolution behavior is controlled by `shared.ParseConfig`:

```go
type ParseConfig struct {
    DetectUnknownFields bool   // Report unrecognized fields
    ResolveInternalRefs bool   // Resolve local #/... refs
    ResolveExternalRefs bool   // Resolve file and URL refs
}
```

| Preset | Internal | External | Unknown Fields |
|--------|:--------:|:--------:|:--------------:|
| `shared.All()` (default) | ✅ | ✅ | ✅ |
| `shared.None()` | ❌ | ❌ | ❌ |

```go
// Default — all features enabled
result, _ := openapi30x.ParseFile("api.yaml")

// Only internal refs
result, _ := openapi30x.ParseFile("api.yaml", &shared.ParseConfig{
    ResolveInternalRefs: true,
})
```

### Accessing Resolved Values

When `ParseFile` is used, resolution runs in a **background goroutine**. The generic ref types (`shared.Ref[T]` and `shared.RefWithMeta[T]`) provide blocking and non-blocking accessors:

| Method | Blocks? | Purpose |
|--------|:-------:|---------|
| `Value()` | ✅ | Returns resolved model; blocks until resolution completes |
| `Circular()` | ✅ | Returns circular flag; blocks until resolution completes |
| `ResolveErr()` | ✅ | Returns resolution error; blocks until done |
| `RawValue()` | ❌ | Returns value immediately (for resolver use) |

```go
result, _ := openapi30x.ParseFile("api.yaml")

// Pattern 1: Auto-blocking — blocks only when you access the ref
petType := result.Document.Components().Schemas()["Pet"].Value().Type()

// Pattern 2: Wait for all resolution first
result.Wait()
pet := result.Document.Components().Schemas()["Pet"].Value()

// Pattern 3: Parse without resolution — Value() returns nil immediately
result, _ := openapi30x.Parse([]byte(data))
ref := result.Document.Components().Schemas()["Pet"]
ref.Value()  // nil (no done channel, no blocking)
```

### Generic Ref Types

All ref wrappers use two generic types defined in `models/shared/ref.go`:

| Generic Type | Used By | Extra Fields |
|---|---|---|
| `shared.Ref[T]` | OpenAPI 2.0 and 3.0 | — |
| `shared.RefWithMeta[T]` | OpenAPI 3.1 | `Summary`, `Description` |

Referenceable model types vary by OpenAPI version:

| Model Type `T` | 2.0 | 3.0 | 3.1 |
|----------------|:---:|:---:|:---:|
| `Schema` | ✅ | ✅ | ✅ |
| `Parameter` | ✅ | ✅ | ✅ |
| `Response` | ✅ | ✅ | ✅ |
| `Callback` | — | ✅ | ✅ |
| `Example` | — | ✅ | ✅ |
| `Header` | — | ✅ | ✅ |
| `Link` | — | ✅ | ✅ |
| `PathItem` | — | ✅ | ✅ |
| `RequestBody` | — | ✅ | ✅ |
| `SecurityScheme` | — | ✅ | ✅ |

The generic struct (shown for `Ref[T]`; `RefWithMeta[T]` adds `Summary` and `Description`):

```go
type Ref[T any] struct {
    Node                           // embedded — VendorExtensions, Trix
    Ref      string                // the $ref string
    value    *T                    // resolved model (private)
    circular bool                  // circular reference detected
    done     chan struct{}          // closed when resolution completes
    err      error                 // resolution error
}
```

Constructors: `shared.NewRef[T](ref)` and `shared.NewRefWithMeta[T](ref)`.

### Utility Functions

The `shared` package exposes ref classification helpers:

```go
shared.IsLocalRef(ref string) bool     // starts with #
shared.IsExternalRef(ref string) bool  // has a file path component
shared.IsRemoteRef(ref string) bool    // starts with http:// or https://
shared.SplitRef(ref string) (filePath, pointer string)
```

---

## Implementation

### Architecture

Resolution is a post-parse phase. The parser first builds the full model tree with `$ref` strings stored as-is and `Value` fields set to `nil`. A separate resolve walk then populates each ref:

```
YAML bytes → Parse Phase → Resolve Phase → Fully resolved model
               (parse.go)    (resolve.go)
```

### RefResolver (`parsers/shared/resolver.go`)

The shared `RefResolver` is the core engine. It handles all three ref types and provides:

- **`Resolve(ref string) → (*ResolveResult, error)`** — resolves any `$ref` string to its target `yaml.Node`
- **Cycle detection** — via a `visiting` map that tracks refs currently being resolved
- **Document caching** — via `fileCache` to avoid re-reading files or re-fetching URLs
- **Canonicalization** — normalizes ref strings (e.g. `./pet.yaml` → `/abs/path/pet.yaml`) for consistent cycle detection keys

Resolution flow for a single ref:

```
Resolve("./models.yaml#/Tag")
  ├─ canonicalize → "/abs/models.yaml#/Tag"
  ├─ Check visiting map → cycle? return Circular=true
  ├─ Mark visiting, defer cleanup
  ├─ SplitRef → (filePath="./models.yaml", pointer="/Tag")
  ├─ loadFile("./models.yaml") → read, parse, cache
  └─ ResolveJSONPointer(root, "/Tag") → target node
```

### Per-Parser Resolve Walk (`resolve.go`)

Each parser version has a `resolve.go` with mutually recursive functions that walk the model tree:

```
resolveDocument
├── resolveComponents
│   ├── resolveSchemaRef      (for each schema)
│   ├── resolveResponseRef    (for each response)
│   ├── resolveParameterRef   (for each parameter)
│   └── ... (all ref types)
├── resolvePathItem           (for each path)
│   └── resolveOperation      (for each HTTP method)
│       ├── resolveParameterRef
│       ├── resolveRequestBodyRef
│       └── resolveResponseRef
└── ...
```

Each `resolve*Ref` function follows the same pattern:

1. Skip if `nil`, no ref, or already resolved
2. Check model-level cycle detection (`resolving` map)
3. Call `r.Resolve(ref.Ref)` to get the target YAML node
4. Check YAML-level circular detection
5. Parse the resolved YAML node into a typed model
6. Mark as resolving, walk children recursively, cleanup

### Circular Reference Detection

Cycles are detected at **two complementary levels**:

| Level | Mechanism | Catches |
|-------|-----------|---------|
| **YAML-level** | `RefResolver.visiting` map | Cross-file cycles (`a.yaml → b.yaml → a.yaml`) |
| **Model-level** | `resolving` map (per-parser) | Same-document cycles (`Schema A → Schema A`) |

The model-level map also uses **pre-registration** — before resolving a component, its canonical path (e.g. `#/components/schemas/TreeNode`) is added to the map. This catches immediate self-references on first encounter.

When a cycle is detected:
- `Circular()` returns `true`
- `Value()` returns `nil`
- No infinite recursion occurs

### Background Resolution

When `ParseFile` is called with resolution enabled:

1. **Parse** — build the document model (synchronous)
2. **Init done channels** — `initRefDoneChannels()` walks all refs and creates `done` channels (synchronous, before goroutine)
3. **Spawn resolver** — `go Resolve(...)` runs in background
4. **Return immediately** — caller gets `ParseResult` right away

Consumers block only when accessing a specific ref's `Value()`. The per-ref `done` channel is closed by `MarkDone()` after the resolver processes that ref.

```go
func parseAndResolve(data []byte, basePath string, cfg *ParseConfig) (*ParseResult, error) {
    doc, _ := parseOpenAPI(docNode, ctx)     // synchronous parse
    initRefDoneChannels(doc)                  // init channels BEFORE goroutine
    result.done = make(chan struct{})
    go func() {
        defer close(result.done)
        Resolve(doc, docNode, basePath)       // background resolve
    }()
    return result, nil                        // immediate return
}
```

`ParseResult.Wait()` blocks until **all** refs are resolved:

```go
result, _ := openapi30x.ParseFile("api.yaml")
result.Wait()  // block until everything is done
```

### Error Handling

Resolution errors are stored **per-ref**, not globally:

```go
ref := result.Document.Components().Schemas()["MissingSchema"]
if err := ref.ResolveErr(); err != nil {
    // e.g. "failed to resolve external ref \"missing.yaml\": file not found"
}
```

### Key Source Files

| File | Purpose |
|------|---------|
| `parsers/shared/resolver.go` | Core `RefResolver` — `Resolve()`, `loadFile()`, `loadURL()`, `canonicalize()` |
| `parsers/shared/config.go` | `ParseConfig` — `ResolveInternalRefs`, `ResolveExternalRefs` |
| `parsers/shared/fetch.go` | `FetchURL()` — used by `ParseFile` for remote entry points |
| `parsers/{version}/resolve.go` | Per-version resolve walk (tree traversal + ref resolution) |
| `parsers/{version}/parse.go` | Entry points (`Parse`, `ParseFile`), background goroutine setup |
| `models/shared/ref.go` | Generic `Ref[T]` and `RefWithMeta[T]` with blocking `Value()`, `Circular()`, `MarkDone()` |

---

## Resolved Reference Types

The following reference mechanisms are now **fully resolved** by the parser. Resolved values are stored in the `Trix` library metadata (not part of the OpenAPI spec itself).

### 1. `$anchor` References (OpenAPI 3.1 only)

The resolver builds an anchor index for the root document and external files. Refs using `#anchorName` syntax resolve to the schema node with the matching `$anchor`:

```go
// $ref: '#pet' resolves to the schema with $anchor: pet
ref.Value() // the resolved schema
```

### 2. `$dynamicRef` / `$dynamicAnchor` (OpenAPI 3.1 only)

The resolver scans for `$dynamicAnchor` declarations and resolves `$dynamicRef` statically to the first matching anchor. The resolved schema is stored in `Trix.ResolvedDynamicRef`:

```go
schema.Trix.ResolvedDynamicRef         // *shared.RefWithMeta[Schema]
schema.Trix.ResolvedDynamicRef.Value() // the schema object
```

> [!NOTE]
> In a parser context (no validation scope), `$dynamicRef` is resolved statically to the first matching `$dynamicAnchor` in the document tree. This is correct for code-generation and tooling use cases.

### 3. `operationRef` in Link Objects (3.0 and 3.1)

After all `$ref` resolution completes, the resolver parses `operationRef` JSON pointers (e.g. `#/paths/~1users~1{id}/get`) and wires them to the already-parsed `Operation` objects. The result is stored in `Trix.ResolvedOperation`:

```go
link.Trix.ResolvedOperation               // *Operation — the target operation
link.Trix.ResolvedOperation.OperationID()  // e.g. "getUser"
```

### 4. `discriminator.mapping` Values (3.0 and 3.1)

Mapping values (both bare schema names like `"Dog"` and explicit refs like `"#/components/schemas/Dog"`) are resolved to schema refs. The result is stored in `Trix.ResolvedMapping` on the Discriminator:

```go
disc.Trix.ResolvedMapping             // map[string]*shared.Ref[Schema] (3.0) or *shared.RefWithMeta[Schema] (3.1)
disc.Trix.ResolvedMapping["dog"].Value() // the Dog schema
```

Bare names are automatically expanded: `"Dog"` → `#/components/schemas/Dog`.

### 5. `$ref` with Sibling Properties (3.0 vs 3.1)

OpenAPI 3.0: `$ref` overrides all sibling properties. OpenAPI 3.1: `summary` and `description` siblings are preserved on `shared.RefWithMeta[T]` (all 3.1 ref types).

### Summary Table

| Feature | Parsed? | Resolved? | Access | Versions |
|---------|:-------:|:---------:|--------|----------|
| `$ref` (standard) | ✅ | ✅ | `ref.Value()` | All |
| External file `$ref` | ✅ | ✅ | `ref.Value()` | All |
| Remote URL `$ref` | ✅ | ✅ | `ref.Value()` | All |
| `$ref` + siblings | ✅ | ✅ | `ref.Summary`, `ref.Description` | 3.1 |
| `$anchor` | ✅ | ✅ | `ref.Value()` (via `#anchorName` ref) | 3.1 |
| `$dynamicRef` | ✅ | ✅ | `schema.Trix.ResolvedDynamicRef` | 3.1 |
| `operationRef` (Link) | ✅ | ✅ | `link.Trix.ResolvedOperation` | 3.0, 3.1 |
| `discriminator.mapping` | ✅ | ✅ | `disc.Trix.ResolvedMapping` | 3.0, 3.1 |

