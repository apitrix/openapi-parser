# Implementation Decisions

## Design Patterns

### 1. Simple Properties Inline
Simple scalar fields are parsed directly in the parent parser:
```go
info.Title = nodeGetString(node, "title")
info.Version = nodeGetString(node, "version")
```

### 2. Complex Properties Delegated
Complex nested objects get separate files following naming convention `{parent}_{property}.go`:
- `info.go` → delegates to `info_contact.go`, `info_license.go`
- `operation.go` → delegates to `operation_parameters.go`, `operation_requestbody.go`
- `schema.go` → delegates to `schema_properties.go`, `schema_allof.go`

### 3. Reference Handling
`$ref` is handled by ref parsers in `ref_{type}.go` files:
```go
// ref_schema.go
if nodeHasRef(node) {
    ref.Ref = nodeGetRef(node)
    return ref, nil
}
ref.Value, err = parseSchema(node, ctx)
```

### 4. Shared Parsers
Common types used across multiple contexts use `shared_` prefix:
- `shared_responses.go` - Responses used in operations
- `shared_securityrequirement.go` - Security requirements

---

## Single-Pass Mapping Iteration

**Problem:** The original implementation used a two-phase pattern for iterating over YAML mapping nodes:

```go
// Old pattern - O(keys × scan) complexity
for _, key := range nodeKeys(node) {
    value := nodeGetValue(node, key)  // Linear scan for each key
    // process value...
}
```

This approach required:
1. First pass: scan all node content to extract keys
2. For each key: re-scan content to find the corresponding value

With *n* keys, this resulted in O(n²) scans of the node content.

**Solution:** Introduced `nodeMapPairs()` using Go 1.23+ range-over-func iterators:

```go
// New pattern - O(n) single-pass iteration
for key, value := range nodeMapPairs(node) {
    // process value directly...
}
```

**Benefits:**
| Aspect | Before | After |
|--------|--------|-------|
| Time complexity | O(n²) | O(n) |
| Memory allocations | 1 slice per map | Zero allocations |
| Code clarity | 2 lines per loop | 1 line per loop |

**Trade-off:** Requires Go 1.23+ (uses `iter.Seq2` from the `iter` package).

**Files affected:** All parsers that iterate over YAML mappings (37 total). The `nodeMapPairs` function is defined in `node_helpers.go` for each parser package.
