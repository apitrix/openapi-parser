---
name: Readonly Models Plan
overview: Convert all model struct fields to private with public getter methods, making models readonly after construction. Parsers stay in place and use constructors. Implement background ref resolution so Parse() returns immediately while refs resolve concurrently.
todos:
  - id: pilot-info
    content: "Pilot: Convert Info/Contact/License in openapi30 to private fields + getters + constructors, update parser to use constructors, update tests"
    status: in-progress
  - id: convert-openapi30
    content: Convert all remaining openapi30 models to readonly pattern
    status: pending
  - id: convert-openapi31
    content: Convert all openapi31 models to readonly pattern
    status: pending
  - id: convert-openapi20
    content: Convert all openapi20 models to readonly pattern
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
    content: Update README and IMPLEMENTATION_DECISIONS.md
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

## Approach: Constructors + Getters (Parsers Stay In Place)

Models use **private fields + getter methods + constructors**. Parsers stay in `parsers/` and use constructors to create immutable model objects. No parser files move.

### Structure (unchanged)

```
models/openapi30/info.go             <- private fields + getters + constructor
models/openapi30/contact.go          <- private fields + getters + constructor
parsers/openapi30x/info.go           <- uses constructor (stays here)
parsers/openapi30x/info_contact.go   <- uses constructor (stays here)
```

### What the model files look like

```go
// models/openapi30/info.go
type Info struct {
    Node           // public embedded — VendorExtensions + Trix stay accessible
    title          string
    description    string
    termsOfService string
    contact        *Contact
    license        *License
    version        string
}

// Getters
func (i *Info) Title() string          { return i.title }
func (i *Info) Description() string    { return i.description }
func (i *Info) TermsOfService() string { return i.termsOfService }
func (i *Info) Contact() *Contact      { return i.contact }
func (i *Info) License() *License      { return i.license }
func (i *Info) Version() string        { return i.version }

// Constructor — used by parsers
func NewInfo(title, description, termsOfService, version string, contact *Contact, license *License) *Info {
    return &Info{
        title: title, description: description, termsOfService: termsOfService,
        version: version, contact: contact, license: license,
    }
}
```

### What the parser looks like (stays in parsers/openapi30x/)

```go
// parsers/openapi30x/info.go — no change in location
func parseOpenAPIInfo(node *yaml.Node, ctx *ParseContext) (*openapi30models.Info, error) {
    if node == nil {
        return nil, ctx.errorf("info is required")
    }

    // Parse sub-objects first
    contact, err := parseInfoContact(node, ctx)
    // collect error...
    license, err := parseInfoLicense(node, ctx)
    // collect error...

    // Create via constructor
    info := openapi30models.NewInfo(
        nodeGetString(node, "title"),
        nodeGetString(node, "description"),
        nodeGetString(node, "termsOfService"),
        nodeGetString(node, "version"),
        contact,
        license,
    )

    // Node-level fields (VendorExtensions + Trix) are still public via embedding
    info.VendorExtensions = parseNodeExtensions(node)
    info.Trix.Source = ctx.nodeSource(node)
    info.Trix.Errors = append(info.Trix.Errors, errors...)

    return info, nil
}
```

### Key design decisions

- **`Node` stays embedded and public** — `VendorExtensions` and `Trix` are library metadata, not spec fields. They remain directly accessible.
- **Constructors accept all spec fields** — parsers pass everything at construction time.
- **SetProperty() methods** if post-construction mutation is needed — for example, if a parser needs to set a field after the initial construction (rare, but possible for complex types).

### Impact on public API

Consumer code changes from field access to method calls:

```go
// Before
doc.Info.Title
doc.Info.Contact.Email

// After  
doc.Info.Title()
doc.Info.Contact().Email()
```

### Conformance tests need updating

The current conformance test uses **reflection on JSON struct tags** to validate field coverage against the OpenAPI JSON schema. With private fields and no JSON tags, this needs a new approach:

- Option A: Maintain a manual mapping of getter method names to JSON property names
- Option B: Keep a parallel set of "spec field names" as a constant list validated against the schema
- Option C: Use reflection on method names with a naming convention

### Impact on JSON serialization

Currently models have `json:"..."` tags for serialization. With private fields, `encoding/json` cannot marshal/unmarshal them. Since the parser uses `yaml.Node` directly (not `encoding/json`), this isn't a breaking change.

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

Resolution state lives directly on each `*Ref` type as **private fields** — not in `Trix`. This keeps `Trix` clean as a public-facing interface with zero internal plumbing.

```go
type SchemaRef struct {
    node     Node
    ref      string
    value    *Schema
    circular bool
    // resolution state — private, set by parser, invisible to consumers
    done     chan struct{} // closed when resolution completes
    err      error        // resolution error, if any
}

func (r *SchemaRef) Ref() string { return r.ref }

func (r *SchemaRef) Value() *Schema {
    if r.done != nil {
        <-r.done // block until this specific ref is resolved
    }
    return r.value
}

func (r *SchemaRef) Circular() bool {
    if r.done != nil {
        <-r.done
    }
    return r.circular
}

func (r *SchemaRef) ResolveErr() error {
    if r.done != nil {
        <-r.done
    }
    return r.err
}
```

Benefits of this approach over an external map or Trix field:
- No map lookup or mutex on every `.Value()` call
- State lives where it belongs (on the ref itself)
- GC-friendly (ref dies → state dies)
- `Trix` stays 100% clean — only `Source` and `Errors`, nothing internal

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
- Resolution state is private on each ref — `Trix` stays a clean public-facing interface
- `done` is `nil` for inline (non-`$ref`) nodes — zero overhead, no blocking
- Only affects ~10 Ref types, not all ~75 model types

### Considerations

- **Circular detection** still works: the `resolving` map and `RefResolver.visiting` map operate within the background goroutine
- **Thread safety**: Each ref's `done` channel provides safe synchronization. The `RefResolver` file cache may need a mutex if multiple goroutines resolve external files in parallel (or resolve sequentially within the background goroutine)
- **Error model**: Resolution errors are per-ref via `ref.ResolveErr()`. The caller can also call a top-level `result.Wait()` to block until all resolution is done and collect all errors
- **Opt-out**: Provide a `ParseSync()` or option flag for callers who want the old blocking behavior

---

## Scale Estimate


| Area                         | Count      | Effort |
| ---------------------------- | ---------- | ------ |
| Model structs to convert     | ~75        | High   |
| Total fields needing getters | ~350+      | High   |
| Parser files to update       | ~300+      | Medium |
| Test files to update         | ~50+       | Medium |
| Conformance test rewrite     | 3 versions | Low    |


Recommend piloting with **one small model** (`Info` in openapi30) to validate the pattern, then systematically converting all models version by version.

## Execution Order

1. Pilot: Convert `Info` model (openapi30) -- validate the pattern end-to-end
4. Implement background ref resolution for openapi30 Ref types (done channels, async resolve goroutine, `Wait()`)
5. Replicate for openapi31
6. Replicate for openapi20
7. Update/rewrite conformance tests
8. Update all integration/parse tests
9. Document ref resolution behavior in @docs as ref_resolution.md


