# OpenAPI 3.0 Parser

This package parses OpenAPI 3.0.x specifications into strongly-typed Go models.

## Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                         parse.go                                 │
│  Parse() → ParseWithUnknownFields() → parseOpenAPI()            │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                       openapi.go                                 │
│  Root document parser - delegates to component parsers          │
└─────────────────────────────────────────────────────────────────┘
          │           │           │           │
          ▼           ▼           ▼           ▼
    ┌─────────┐ ┌─────────┐ ┌─────────┐ ┌─────────┐
    │ info.go │ │pathitem │ │component│ │ server  │
    └─────────┘ └─────────┘ └─────────┘ └─────────┘
         │           │           │
         ▼           ▼           ▼
    ┌─────────┐ ┌─────────┐ ┌─────────┐
    │ contact │ │operation│ │ schemas │
    │ license │ │requestb │ │ params  │
    └─────────┘ └─────────┘ └─────────┘
```

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
- `schema.go` → delegates to `schema_properties.go`, `schema_allof.go`, etc.

### 3. Reference Handling
`$ref` is handled by ref parsers in `ref_{type}.go` files:
```go
// ref_schema.go
if nodeHasRef(node) {
    ref.Ref = nodeGetRef(node)
    return ref, nil
}
// Parse inline schema
ref.Value, err = parseSchema(node, ctx)
```

### 4. Shared Parsers
Common types used across multiple contexts use `shared_` prefix:
- `shared_responses.go` - Responses used in operations
- `shared_securityrequirement.go` - Security requirements

## Core Components

| File | Purpose |
|------|---------|
| `parse.go` | Entry points: `Parse()`, `ParseWithUnknownFields()` |
| `context.go` | `ParseContext` - tracks JSON path, caches components |
| `node_helpers.go` | yaml.Node utilities for extracting values |
| `known_fields.go` | Valid field lists for unknown field detection |
| `errors.go` | `ParseError` with path and location info |
| `unknown_fields.go` | Detection of unrecognized fields |

## File Organization

```
parsers/openapi30/
├── parse.go              # Entry points
├── context.go            # ParseContext
├── openapi.go            # Root parser
├── info.go               # Info parser
├── info_contact.go       # Info.Contact parser
├── info_license.go       # Info.License parser
├── pathitem.go           # PathItem parser
├── operation.go          # Operation parser
├── operation_*.go        # Operation sub-properties
├── parameter.go          # Parameter parser
├── schema.go             # Schema parser
├── schema_*.go           # Schema sub-properties
├── components.go         # Components parser
├── components_*.go       # Components sub-sections
├── ref_*.go              # Reference parsers
├── shared_*.go           # Shared parsers
├── *_test.go             # Unit tests
├── helpers.go            # Map helpers
├── node_helpers.go       # yaml.Node helpers
├── known_fields.go       # Known field definitions
├── unknown_fields.go     # Unknown field detection
└── errors.go             # Error types
```

## Usage

```go
import "openapi-parser/parsers/openapi30"

doc, err := openapi30.Parse(data)
if err != nil {
    log.Fatal(err)
}

// With unknown field detection
result, err := openapi30.ParseWithUnknownFields(data)
for _, f := range result.UnknownFields {
    log.Printf("Unknown: %s at %s", f.Name, f.Path)
}
```

## Testing

Tests follow the **Arrange-Act-Assert (AAA)** pattern:
```go
func TestParseInfo_Contact(t *testing.T) {
    // Arrange
    yaml := `openapi: "3.0.3"...`
    
    // Act
    doc, err := Parse([]byte(yaml))
    
    // Assert
    require.NoError(t, err)
    assert.Equal(t, "expected", doc.Info.Contact.Name)
}
```

Run tests:
```bash
go test -v ./parsers/openapi30/...
```
