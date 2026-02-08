# Implementation Decisions

## Design Patterns

### 1. Trix Namespace — Separating Spec from Library

**Problem:** Model structs need to carry library-level metadata (source locations, raw data) and vendor extensions alongside spec-defined fields. Mixing them pollutes the struct surface and makes it unclear what's from the OpenAPI spec vs. what's library infrastructure.

**Decision:** Every model embeds `Node`, which contains exactly two non-spec concerns:

```go
type Node struct {
    VendorExtensions map[string]interface{} `json:"-" yaml:"-"`
    Trix             Trix                   `json:"-" yaml:"-"`
}

type Trix struct {
    Source NodeSource   `json:"-" yaml:"-"` // line/column, raw data
    Errors []ParseError `json:"-" yaml:"-"` // recoverable parsing errors
}
```

- **`VendorExtensions`** — Holds `x-*` extension fields defined by the user's spec. These are part of the OpenAPI standard but not part of any specific object schema, so they live on `Node` rather than on each struct.
- **`Trix`** — Named after the library (apitrix). Contains all library-provided metadata and functionality. Holds `Source` (line/column numbers and raw parsed data) and `Errors` (recoverable parsing errors encountered while parsing this node's children). Future capabilities (validation, conversion, parent traversal) will be added here.

**Benefits:**
- A developer sees `schema.Title` (spec) vs. `schema.Trix.Source.Start.Line` (library) — instantly clear which is which
- Adding new library features only expands `Trix`, never the model struct itself
- All `Trix` fields are `json:"-"`, so serialization stays clean
- Each OpenAPI version (v2.0, v3.0, v3.1) has its own `Trix` that can diverge as needed

### 1b. Error Collection on Nodes

**Problem:** A "fail-fast" parser returns on the first error, forcing users to fix one issue at a time. For linting, IDE support, and batch validation, it's better to surface all issues at once.

**Decision:** Recoverable errors encountered while parsing a node's children are **collected on that node's `Trix.Errors`** instead of causing the parser to return early:

```go
// Before (fail-fast)
op.Responses, err = parseResponses(node, ctx)
if err != nil {
    return nil, err
}

// After (error collection)
op.Responses, err = parseResponses(node, ctx)
if err != nil {
    op.Trix.Errors = append(op.Trix.Errors, toParseError(err))
}
```

`Parse()` always returns a (possibly partial) document. Errors are accessible per-node:

```go
doc, err := openapi30x.Parse(data)
for _, e := range doc.Trix.Errors {
    log.Printf("root-level error: %s", e.Message)
}
for _, e := range doc.Info.Trix.Errors {
    log.Printf("info error: %s", e.Message)
}
```

**Fatal errors** that prevent *any* parsing still return immediately:
- Version validation (`openapi`/`swagger` field missing or unsupported)
- Malformed YAML/JSON

### 2. Simple Properties Inline
Simple scalar fields are parsed directly in the parent parser:
```go
info.Title = shared.NodeGetString(node, "title")
info.Version = shared.NodeGetString(node, "version")
```

### 3. Complex Properties Delegated
Complex nested objects get separate files following naming convention `{parent}_{property}.go`:
- `info.go` → delegates to `info_contact.go`, `info_license.go`
- `operation.go` → delegates to `operation_parameters.go`, `operation_requestbody.go`
- `schema.go` → delegates to `schema_properties.go`, `schema_allof.go`

### 4. Reference Handling
`$ref` is handled by ref parsers in `ref_{type}.go` files:
```go
// ref_schema.go
if shared.NodeHasRef(node) {
    ref.Ref = shared.NodeGetRef(node)
    return ref, nil
}
ref.Value, err = parseSchema(node, ctx)
```

### 5. Shared Parsers
Common types used across multiple contexts use `shared_` prefix:
- `shared_responses.go` — Responses used in operations
- `shared_securityrequirement.go` — Security requirements

### 6. Shared Internal Package
All three parsers (`openapi20`, `openapi30x`, `openapi31x`) share an `internal/shared` package that contains:
- Node helpers (`node.go`) — extract typed values from `yaml.Node`
- Map utilities (`maputil.go`) — extract typed values from `map[string]interface{}`
- Error types (`errors.go`) — `ParseError` with path and source location
- Unknown field detection (`unknown.go`) — `DetectUnknownNodeFields`
- Field sets (`set.go`) — `ToSet` for O(1) lookup
- Reference resolver (`resolver.go`) — `RefResolver` for `$ref` resolution

---

## Two-Phase Parsing: Parse then Resolve

**Problem:** `$ref` references point to other parts of the document (or external files). During initial parsing, those targets may not have been parsed yet.

**Decision:** Split parsing into two phases:

1. **Parse phase** (`parse.go → openapi.go`) — Walk the YAML tree and populate model structs. `$ref` values are stored as strings in `Ref` fields; `Value` remains `nil`.

2. **Resolve phase** (`resolve.go`) — Walk the populated model tree. For every `*Ref` type with a non-empty `Ref` and `nil` `Value`, call `RefResolver.Resolve()` to get the target YAML node, parse it, and populate `Value`.

```go
// ParseFile triggers both phases:
doc, err := parseOpenAPI(docNode, ctx)   // Phase 1: parse
err = Resolve(doc, docNode, basePath)     // Phase 2: resolve
```

**Benefits:**
- Forward references work naturally
- External file references are loaded on demand and cached
- Circular references are detected and flagged instead of causing infinite loops

---

## Two-Level Circular Reference Detection

**Problem:** Naive `$ref` resolution causes infinite recursion on self-referencing schemas like:
```yaml
definitions:
  TreeNode:
    type: object
    properties:
      children:
        type: array
        items:
          $ref: '#/definitions/TreeNode'
```

**Decision:** Use two complementary layers of cycle detection:

### Level 1: YAML Resolver (`RefResolver.visiting`)
The `RefResolver` maintains a `visiting` map that tracks which canonical refs are currently being resolved within a single `Resolve()` call. Uses `defer delete` for stack-like cleanup:
```go
func (r *RefResolver) Resolve(ref string) (*ResolveResult, error) {
    canonicalRef := r.canonicalize(ref)
    if r.visiting[canonicalRef] {
        return &ResolveResult{Circular: true}, nil
    }
    r.visiting[canonicalRef] = true
    defer func() { delete(r.visiting, canonicalRef) }()
    // ... resolve ...
}
```

### Level 2: Model-Level `resolving` Map
A `resolving map[string]bool` is threaded through all per-parser resolve functions. Before walking a top-level component definition, its canonical `$ref` path is pre-registered:
```go
for name, ref := range c.Schemas {
    canonicalRef := "#/components/schemas/" + name
    resolving[canonicalRef] = true        // pre-register
    resolveSchemaRef(ref, r, resolving)   // walk children
    delete(resolving, canonicalRef)       // cleanup
}
```

Individual ref resolvers check `resolving[ref.Ref]` before recursing:
```go
func resolveSchemaRef(ref *SchemaRef, r *RefResolver, resolving map[string]bool) error {
    if resolving[ref.Ref] {
        ref.Circular = true
        return nil                         // stop recursion
    }
    // ... resolve normally ...
}
```

**Why two levels?**
- Level 1 catches cycles during YAML-level pointer traversal (e.g., `$ref` chains through external files)
- Level 2 catches cycles during model-tree walking (e.g., schema A → schema B → schema A)
- Pre-registration ensures first-encounter self-references are caught immediately

---

## Single-Pass Mapping Iteration

**Problem:** The original implementation used a two-phase pattern for iterating over YAML mapping nodes:

```go
// Old pattern - O(keys × scan) complexity
for _, key := range nodeKeys(node) {
    value := nodeGetValue(node, key)  // Linear scan for each key
}
```

**Solution:** Introduced `NodeMapPairs()` using Go 1.23+ range-over-func iterators:

```go
// New pattern - O(n) single-pass iteration
for key, value := range shared.NodeMapPairs(node) {
    // process value directly
}
```

| Aspect | Before | After |
|--------|--------|-------|
| Time complexity | O(n²) | O(n) |
| Memory allocations | 1 slice per map | Zero |
| Code clarity | 2 lines per loop | 1 line per loop |

**Trade-off:** Requires Go 1.23+ (uses `iter.Seq2`).

---

## Precomputed Known-Field Sets

**Problem:** Unknown field detection built a `map[string]bool` from a `[]string` slice on every call:

```go
// Old pattern - O(n) map construction per node
func detectUnknownNodeFields(node *yaml.Node, knownFields []string, path string) {
    known := make(map[string]bool, len(knownFields))
    for _, f := range knownFields { known[f] = true }
}
```

**Solution:** Precompute `map[string]struct{}` sets at init time in `known_fields.go`:

```go
// Precomputed once at init
var schemaKnownFieldsSet = shared.ToSet(schemaKnownFields)

// O(1) lookup, zero allocations per call
shared.DetectUnknownNodeFields(node, schemaKnownFieldsSet, path)
```

| Aspect | Before | After |
|--------|--------|-------|
| Map construction | Per node | Once at init |
| Memory per call | 1 map allocation | Zero |
| Lookup | O(1) after O(n) build | O(1) |

---

## Filesystem Abstraction with Afero

**Problem:** Testing external `$ref` resolution requires reading files from disk, making tests depend on real filesystem state.

**Decision:** Use [afero](https://github.com/spf13/afero) for filesystem abstraction:

```go
// Production: real filesystem
resolver := shared.NewRefResolver(basePath, root)

// Testing: in-memory filesystem
memFs := afero.NewMemMapFs()
afero.WriteFile(memFs, "/specs/pet.yaml", petData, 0644)
resolver := shared.NewRefResolverWithFs(basePath, root, memFs)
```

**Benefits:**
- Tests run in-memory with no cleanup needed
- Tests are deterministic and parallel-safe
- No temp files or OS-level filesystem operations
