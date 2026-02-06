# OpenAPI Parsers

This package contains parsers for OpenAPI specifications:

| Package | Specification | Go Import |
|---------|---------------|-----------|
| `openapi30` | OpenAPI 3.0.x | `openapi-parser/parsers/openapi30` |
| `openapi20` | OpenAPI 2.0 (Swagger) | `openapi-parser/parsers/openapi20` |

Both parsers share identical architecture and design patterns.

## Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                         parse.go                                 │
│  Parse() → ParseWithUnknownFields() → parseOpenAPI/Swagger()    │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                    openapi.go / swagger.go                       │
│  Root document parser - delegates to component parsers          │
└─────────────────────────────────────────────────────────────────┘
          │           │           │           │
          ▼           ▼           ▼           ▼
    ┌─────────┐ ┌─────────┐ ┌─────────┐ ┌─────────┐
    │ info.go │ │pathitem │ │component│ │ server  │
    └─────────┘ └─────────┘ └─────────┘ └─────────┘
```

## Core Components

| File | Purpose |
|------|---------|
| `parse.go` | Entry points: `Parse()`, `ParseWithUnknownFields()` |
| `context.go` | `ParseContext` - tracks JSON path, caches components |
| `node_helpers.go` | yaml.Node utilities for extracting values |
| `known_fields.go` | Valid field lists for unknown field detection |
| `errors.go` | `ParseError` with path and location info |

## Usage

```go
// OpenAPI 3.0
import "openapi-parser/parsers/openapi30"
doc, err := openapi30.Parse(data)

// OpenAPI 2.0 (Swagger)
import "openapi-parser/parsers/openapi20"
doc, err := openapi20.Parse(data)

// With unknown field detection
result, err := openapi30.ParseWithUnknownFields(data)
for _, f := range result.UnknownFields {
    log.Printf("Unknown: %s at %s", f.Name, f.Path)
}
```

## Testing

Tests follow the **Arrange-Act-Assert (AAA)** pattern.

```bash
# Run all parser tests
go test -v ./parsers/...

# Run specific parser tests
go test -v ./parsers/openapi30/...
go test -v ./parsers/openapi20/...
```

## Implementation Details

See [IMPLEMENTATION_DECISIONS.md](IMPLEMENTATION_DECISIONS.md) for design patterns and technical optimizations.
