# Unknown Field Detection

The parser detects fields in your OpenAPI document that are not part of the OpenAPI specification. This helps identify typos, unsupported properties, or misplaced fields.

All parse functions (`Parse`, `ParseReader`, `ParseFile`) always return unknown fields in the result — there is no separate API for unknown field detection.

## Quick Start

```go
import "openapi-parser/parsers/openapi30x"

result, err := openapi30x.Parse(data)
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

### ParseResult

All parse functions return `*ParseResult`:

```go
type ParseResult struct {
    Document      *OpenAPI       // The parsed OpenAPI specification
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
- Valid OpenAPI fields in their correct context

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
