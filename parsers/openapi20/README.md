# OpenAPI 2.0 (Swagger) Parser

This package parses Swagger 2.0 specifications into strongly-typed Go models.

## Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                         parse.go                                 │
│  Parse() → ParseWithUnknownFields() → parseSwagger()            │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                       swagger.go                                 │
│  Root document parser - delegates to component parsers          │
└─────────────────────────────────────────────────────────────────┘
          │           │           │           │
          ▼           ▼           ▼           ▼
    ┌─────────┐ ┌─────────┐ ┌─────────┐ ┌─────────┐
    │ info.go │ │pathitem │ │ schema  │ │security │
    └─────────┘ └─────────┘ └─────────┘ └─────────┘
         │           │
         ▼           ▼
    ┌─────────┐ ┌─────────┐
    │ contact │ │operation│
    │ license │ │parameter│
    └─────────┘ └─────────┘
```

## Design Patterns

### 1. Simple Properties Inline
Simple scalar fields are parsed directly in the parent parser:
```go
info.Title = nodeGetString(node, "title")
info.Version = nodeGetString(node, "version")
```

### 2. Complex Properties Delegated
Complex nested objects get separate files:
- `info.go` → delegates to `info_contact.go`, `info_license.go`
- `schema.go` → handles properties, additionalProperties, allOf

### 3. Reference Handling
`$ref` is handled by ref parsers that check for reference or parse inline:
```go
// ref_schema.go
if nodeHasRef(node) {
    ref.Ref = nodeGetRef(node)
    return ref, nil
}
// Parse inline schema
ref.Value, err = parseSchema(node, ctx)
```

## Core Components

| File | Purpose |
|------|---------|
| `parse.go` | Entry points: `Parse()`, `ParseWithUnknownFields()` |
| `context.go` | `ParseContext` - tracks JSON path, caches definitions |
| `node_helpers.go` | yaml.Node utilities for extracting values |
| `known_fields.go` | Valid field lists for unknown detection |
| `errors.go` | `ParseError` with path and location |

## Usage

```go
import "openapi-parser/parsers/openapi20"

doc, err := openapi20.Parse(data)
if err != nil {
    log.Fatal(err)
}

// With unknown field detection
result, err := openapi20.ParseWithUnknownFields(data)
for _, f := range result.UnknownFields {
    log.Printf("Unknown: %s at %s", f.Name, f.Path)
}
```

## Testing

Tests follow the **Arrange-Act-Assert (AAA)** pattern:
```go
func TestParseInfo_Contact(t *testing.T) {
    // Arrange
    yaml := `swagger: "2.0"...`
    
    // Act
    doc, err := Parse([]byte(yaml))
    
    // Assert
    require.NoError(t, err)
    assert.Equal(t, "expected", doc.Info.Contact.Name)
}
```

Run tests:
```bash
go test -v ./parsers/openapi20/...
```
