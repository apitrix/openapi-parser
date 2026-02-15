# Unknown Field Detection

The parser detects fields in your OpenAPI document that are not part of the OpenAPI specification. This helps identify typos, unsupported properties, or misplaced fields.

All parse functions (`Parse`, `ParseReader`, `ParseFile`) always return unknown fields in the result — there is no separate API for unknown field detection.

## Quick Start

Unknown fields appear in `result.Errors` with `Kind == "unknown_field"`:

```go
import "github.com/apitrix/openapi-parser/parsers/openapi30x"

result, err := openapi30x.Parse(data)
if err != nil {
    log.Fatal(err)
}

for _, e := range result.Errors {
    if e.Kind == "unknown_field" {
        log.Printf("Unknown field at %v (line %d): %s", e.Path, e.Line, e.Message)
    }
}
```

## API Reference

### ParseResult.Errors

All parse functions return `*ParseResult` with `Errors []*ParseError`. Unknown fields are reported as `ParseError` with `Kind: "unknown_field"`:

```go
type ParseError struct {
    Path    []string // JSON path
    Message string   // Error message
    Kind    string   // "error", "unknown_field", or "resolve_error"
    Line    int      // Source line (1-based)
    Column  int      // Source column (1-based)
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
