# Unknown Field Detection

The parser can detect fields in your OpenAPI document that are not part of the OpenAPI 3.0 specification. This helps identify typos, unsupported properties, or misplaced fields.

## Quick Start

```go
import "openapi-parser/parsers/openapi30"

// Parse with unknown field detection
result, err := openapi30.ParseWithUnknownFields(data)
if err != nil {
    log.Fatal(err)
}

// Use the parsed document
doc := result.Document

// Check for unknown fields
if len(result.UnknownFields) > 0 {
    for _, f := range result.UnknownFields {
        log.Printf("Warning: unknown field '%s' at %s (line %d, col %d)",
            f.Key, f.Path, f.Line, f.Column)
    }
}
```

## API Reference

### ParseWithUnknownFields

```go
func ParseWithUnknownFields(data []byte) (*ParseResult, error)
```

Parses an OpenAPI 3.0 specification and returns both the parsed document and any unknown fields detected.

### ParseReaderWithUnknownFields

```go
func ParseReaderWithUnknownFields(r io.Reader) (*ParseResult, error)
```

Same as `ParseWithUnknownFields` but reads from an `io.Reader`.

### ParseResult

```go
type ParseResult struct {
    Document      *v30.OpenAPI   // The parsed OpenAPI specification
    UnknownFields []UnknownField // Fields not recognized as valid OpenAPI fields
}
```

### UnknownField

```go
type UnknownField struct {
    Path   string // JSON path, e.g., "paths./pets.get.responses.200"
    Key    string // The unknown field name
    Line   int    // Source line number (1-based)
    Column int    // Source column number (1-based)
}
```

## What Gets Detected

**Detected as unknown:**
- Typos in field names (e.g., `summery` instead of `summary`)
- Fields in the wrong context (e.g., `operationId` at the path level)
- Fields from different OpenAPI versions

**NOT detected as unknown:**
- Valid extension fields (`x-*`) - these are allowed by the spec
- Valid OpenAPI 3.0 fields in their correct context

## Examples

### Detecting a Typo

```yaml
openapi: "3.0.3"
info:
  title: "My API"
  verison: "1.0.0"  # Typo! Should be "version"
paths: {}
```

Output:
```
Warning: unknown field 'verison' at info (line 4, col 3)
```

### Detecting Misplaced Fields

```yaml
openapi: "3.0.3"
info:
  title: "My API"
  version: "1.0.0"
paths:
  /pets:
    operationId: "getPets"  # Wrong! operationId belongs inside get/post/etc
    get:
      responses:
        "200":
          description: "OK"
```

Output:
```
Warning: unknown field 'operationId' at paths./pets (line 7, col 5)
```

### Extensions Are Valid

```yaml
openapi: "3.0.3"
info:
  title: "My API"
  version: "1.0.0"
  x-custom-field: "This is allowed"  # No warning
paths: {}
```

No warnings - extension fields starting with `x-` are valid per the OpenAPI spec.

## Backward Compatibility

The existing `Parse()` function works exactly as before:

```go
// This still works - no breaking changes
doc, err := openapi30.Parse(data)
```

The only difference is that `Parse()` internally calls `ParseWithUnknownFields()` and discards the unknown field information.
