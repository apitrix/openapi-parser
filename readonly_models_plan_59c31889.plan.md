---
name: Readonly Models Plan
overview: Convert all model struct fields to private with public getter methods, making models readonly after construction. Implement background ref resolution so Parse() returns immediately while refs resolve concurrently.
todos:
  - id: pilot-info
    content: "Pilot: Convert Info model in openapi30 to private fields + getters, move its parser into the model package, update tests"
    status: pending
  - id: shared-utils
    content: Move parsers/internal/shared/ utilities to a location importable by model packages
    status: pending
  - id: convert-openapi30
    content: Convert all openapi30 models to readonly + move all openapi30x parsers into models/openapi30/
    status: pending
  - id: convert-openapi31
    content: Convert all openapi31 models to readonly + move all openapi31x parsers into models/openapi31/
    status: pending
  - id: convert-openapi20
    content: Convert all openapi20 models to readonly + move all openapi20 parsers into models/openapi20/
    status: pending
  - id: conformance-tests
    content: Rewrite conformance tests for the new getter-method-based models
    status: pending
  - id: update-tests
    content: Update all integration and parse tests to use getter methods instead of field access
    status: pending
  - id: background-refs
    content: Implement background ref resolution — Parse() returns model immediately, refs resolve concurrently, Value() blocks only if not yet resolved
    status: pending
  - id: cleanup
    content: Remove empty parsers/ directory, update README and IMPLEMENTATION_DECISIONS.md
    status: pending
isProject: false
---

# Readonly Models with Getter Methods

## Current State

Models are plain structs with **public fields** across 3 API versions (~75 structs, ~350+ fields):

```go
// models/openapi30/info.go — current
type Info struct {
    Node
    Title          string   `json:"title" yaml:"title"`
    Description    string   `json:"description,omitempty" yaml:"description,omitempty"`
    Contact        *Contact `json:"contact,omitempty" yaml:"contact,omitempty"`
    // ...
}
```

Parsers (separate package) set fields directly: `info.Title = shared.NodeGetString(node, "title")`
Tests access fields directly: `assert.Equal(t, "Test API", result.Document.Info.Title)`

## Core Challenge

In Go, if fields are private (lowercase), code in a **different package** cannot read or write them. The models live in `models/openapi30` and the parsers live in `parsers/openapi30x` -- two different packages. Making fields private means parsers can no longer set them.

## Recommended Approach: Merge Parsers Into Model Packages

Move parsing code into the same package as models. This is the **only approach that avoids duplication** and gives true readonly without boilerplate.

### Before (current structure)

```
models/openapi30/schema.go          <- struct definition
parsers/openapi30x/schema.go        <- parsing logic (different package)
parsers/openapi30x/schema_allof.go  <- sub-parsers
```

### After (merged structure)

```
models/openapi30/schema.go           <- private fields + getter methods
models/openapi30/schema_parse.go     <- parsing logic (same package, writes private fields)
models/openapi30/schema_allof.go     <- sub-parsers (same package)
models/openapi30/parse_context.go    <- ParseContext, helpers
```

### What the model files look like after

```go
// models/openapi30/info.go — after
type Info struct {
    node           Node     // private embedded
    title          string
    description    string
    termsOfService string
    contact        *Contact
    license        *License
    version        string
}

func (i *Info) Title() string          { return i.title }
func (i *Info) Description() string    { return i.description }
func (i *Info) TermsOfService() string { return i.termsOfService }
func (i *Info) Contact() *Contact      { return i.contact }
func (i *Info) License() *License      { return i.license }
func (i *Info) Version() string        { return i.version }
func (i *Info) Node() Node             { return i.node }
```

### What the parser files look like after

```go
// models/openapi30/info_parse.go — same package, can set private fields
func parseInfo(node *yaml.Node, ctx *parseContext) *Info {
    info := &Info{}
    info.title = nodeGetString(node, "title")
    info.description = nodeGetString(node, "description")
    // ...
    return info
}
```

### Impact on public API

Consumer code changes from field access to method calls:

```go
// Before
doc.Info.Title
doc.Components.Schemas["Pet"].Value.Properties["name"]

// After
doc.Info().Title()
doc.Components().Schemas()["Pet"].Value().Properties()["name"]
```

### Conformance tests need updating

The current conformance test uses **reflection on JSON struct tags** to validate field coverage against the OpenAPI JSON schema. With private fields and no JSON tags, this needs a new approach:

- Option A: Maintain a manual mapping of getter method names to JSON property names
- Option B: Keep a parallel set of "spec field names" as a constant list validated against the schema
- Option C: Use reflection on method names with a naming convention

### Impact on JSON serialization

Currently models have `json:"..."` tags for serialization. With private fields, `encoding/json` cannot marshal/unmarshal them. If JSON output is needed in the future:

- Implement `json.Marshaler` on each model, or
- Use a separate serialization layer

Since all current tags use `json:"-"` for library metadata and the models aren't deserialized via `encoding/json` (the parser uses `yaml.Node` directly), this isn't a breaking change for current usage.

---

## Background Ref Resolution

Instead of the current eager resolve phase that blocks `Parse()` from returning, refs resolve **concurrently in the background**. The caller gets the model immediately and can inspect all non-ref properties. When `.Value()` is called on a ref, it blocks only if that ref hasn't finished resolving yet.

### Current flow (blocking)

```
Parse() ──► build model tree ──► Resolve() walks entire tree ──► return model
                                 (blocks until ALL refs done)
```

### New flow (concurrent)

```
Parse() ──► build model tree ──► return model immediately
                 │
                 └──► background goroutine: resolve refs concurrently
                      (each ref gets a done channel)
```

### Implementation

Resolution infrastructure lives in `Trix`, keeping the model surface clean. The `Trix` struct in `models/shared/meta.go` gains a ref resolution field:

```go
// models/shared/meta.go
type RefState struct {
    Done chan struct{} // closed when resolution completes
    Err  error        // resolution error, if any
}

type Trix struct {
    Source NodeSource    // source location info
    Errors []ParseError // parsing errors
    Ref    *RefState    // nil for non-ref types; set by parser for $ref nodes
}
```

The Ref model types stay clean — only spec-defined fields + Node:

```go
type SchemaRef struct {
    node     Node
    ref      string
    value    *Schema
    circular bool
}

func (r *SchemaRef) Ref() string { return r.ref }

func (r *SchemaRef) Value() *Schema {
    if rs := r.node.Trix.Ref; rs != nil {
        <-rs.Done // block until this specific ref is resolved
    }
    return r.value
}

func (r *SchemaRef) Circular() bool {
    if rs := r.node.Trix.Ref; rs != nil {
        <-rs.Done
    }
    return r.circular
}
```

Resolution errors are accessed through `Trix` consistently with other error patterns:

```go
ref.Trix().Ref.Err  // per-ref resolution error (after waiting)
ref.Trix().Errors   // parse errors (already existing pattern)
```

The parser kicks off background resolution before returning:

```go
func Parse(data []byte) (*ParseResult, error) {
    doc, root, err := parseDocument(data)
    if err != nil {
        return nil, err
    }

    // Start background ref resolution
    go resolveAllRefs(doc, root, basePath)

    return &ParseResult{Document: doc}, nil
}
```

### Benefits

- `Parse()` returns immediately — caller can inspect `Info()`, `Paths()`, operation metadata, etc. without waiting for ref resolution
- Refs resolve concurrently — external files load in parallel
- `.Value()` blocks only on the specific ref being accessed, not all refs
- Model structs stay clean — resolution state lives in `Trix.Ref`, consistent with the "library infra in Trix" principle
- `Trix.Ref` is `nil` for non-ref types — zero overhead for regular models
- Only affects ~10 Ref types, not all ~75 model types
- Natural fit with the getter-method readonly architecture

### Considerations

- **Circular detection** still works: the `resolving` map and `RefResolver.visiting` map operate within the background goroutine
- **Thread safety**: Each ref's `Done` channel provides safe synchronization. The `RefResolver` file cache may need a mutex if multiple goroutines resolve external files in parallel (or resolve sequentially within the background goroutine)
- **Error model**: Resolution errors are per-ref on `Trix.Ref.Err`. The caller can also call a top-level `result.Wait()` to block until all resolution is done and collect all errors
- **Opt-out**: Provide a `ParseSync()` or option flag for callers who want the old blocking behavior

---

## Scale Estimate


| Area                         | Count      | Effort |
| ---------------------------- | ---------- | ------ |
| Model structs to convert     | ~75        | High   |
| Total fields needing getters | ~350+      | High   |
| Parser files to move/update  | ~300+      | High   |
| Test files to update         | ~50+       | Medium |
| Conformance test rewrite     | 3 versions | Low    |


Recommend piloting with **one small model** (`Info` in openapi30) to validate the pattern, then systematically converting all models version by version.

## Execution Order

1. Pilot: Convert `Info` model (openapi30) -- validate the pattern end-to-end
2. Move shared parser utilities (`parsers/internal/shared/`) into a shared location accessible from model packages
3. Convert all openapi30 models + move their parsers into `models/openapi30/`
4. Implement background ref resolution for openapi30 Ref types (done channels, async resolve goroutine, `Wait()`)
5. Replicate for openapi31
6. Replicate for openapi20
7. Update/rewrite conformance tests
8. Update all integration/parse tests
9. Remove empty `parsers/` directory
10. Document ref resolution behavior in @docs as ref_resolution.md

