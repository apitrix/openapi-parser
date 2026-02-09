# OpenAPI Parsers

This package contains parsers for OpenAPI specifications:

| Package | Specification | Go Import |
|---------|---------------|-----------|
| `openapi20` | OpenAPI 2.0 (Swagger) | `openapi-parser/parsers/openapi20` |
| `openapi30x` | OpenAPI 3.0.x | `openapi-parser/parsers/openapi30x` |
| `openapi31x` | OpenAPI 3.1.x / 3.2.x | `openapi-parser/parsers/openapi31x` |

All parsers share identical architecture and a common `internal/shared` package.

## Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                          parse.go                               │
│  Parse() / ParseFile() / ParseReader()                          │
└─────────────────────────────────────────────────────────────────┘
              │                              │
              ▼                              ▼
┌──────────────────────────┐   ┌──────────────────────────────┐
│  openapi.go / swagger.go │   │         resolve.go            │
│  Root document parser    │   │  $ref resolution & circular   │
│  (delegates to component │   │  detection (post-parse walk)  │
│  parsers)                │   └──────────────────────────────────┘
└──────────────────────────┘               │
  │     │     │     │                      ▼
  ▼     ▼     ▼     ▼         ┌──────────────────────────────────┐
┌────┐┌────┐┌────┐┌────┐     │    internal/shared/resolver.go   │
│info││path││comp││serv│     │    RefResolver — YAML-level       │
│ .go││item││onnt││er  │     │    ref lookup, file caching,      │
│    ││ .go││ .go││.go │     │    cycle detection                │
└────┘└────┘└────┘└────┘     └──────────────────────────────────┘
```

## Shared Internal Package (`internal/shared`)

Common utilities extracted across all three parsers:

| File | Purpose |
|------|---------|
| `resolver.go` | `RefResolver` — resolves local/external `$ref`, caches files, detects cycles |
| `node.go` | `yaml.Node` helpers — `NodeGetValue`, `NodeToMap`, `NodeMapPairs`, etc. |
| `maputil.go` | `map[string]interface{}` value extraction — `GetString`, `GetBoolPtr`, etc. |
| `errors.go` | `ParseError` with JSON path and source location |
| `unknown.go` | `DetectUnknownNodeFields` — finds non-spec fields |
| `set.go` | `ToSet` — converts `[]string` to `map[string]struct{}` |

## Per-Parser Core Files

| File | Purpose |
|------|---------|
| `parse.go` | Entry points: `Parse()`, `ParseReader()`, `ParseFile()` — all return `*ParseResult` |
| `context.go` | `ParseContext` — tracks JSON path, caches components |
| `resolve.go` | Post-parse `$ref` resolution walk with circular detection |
| `known_fields.go` | Precomputed valid field sets for unknown field detection |

## Usage

```go
// OpenAPI 3.1
import "openapi-parser/parsers/openapi31x"
result, err := openapi31x.Parse(data)
fmt.Println(result.Document.Info.Title)

// OpenAPI 3.0
import "openapi-parser/parsers/openapi30x"
result, err := openapi30x.Parse(data)

// OpenAPI 2.0 (Swagger)
import "openapi-parser/parsers/openapi20"
result, err := openapi20.Parse(data)

// Parse from file with full $ref resolution
result, err := openapi30x.ParseFile("openapi.yaml")

// Unknown fields are always available in the result
result, err := openapi30x.Parse(data)
for _, f := range result.UnknownFields {
    log.Printf("Unknown: %s at %s", f.Key, f.Path)
}

// Check for parse-time issues on nodes (errors + unknown fields)
result, err := openapi30x.Parse(data)
if err != nil {
    log.Fatal(err) // fatal: bad YAML or unsupported version
}
for _, e := range result.Document.Info.Trix.Errors {
    switch e.Kind {
    case "error":
        log.Printf("parse error: %s", e.Message)
    case "unknown_field":
        log.Printf("unknown field: %s", e.Message)
    }
}
```

## Testing

Tests follow the **Arrange-Act-Assert (AAA)** pattern and use [testify](https://github.com/stretchr/testify).

```bash
# Run all parser tests
go test -v ./parsers/...

# Run specific parser tests
go test -v ./parsers/openapi20/...
go test -v ./parsers/openapi30x/...
go test -v ./parsers/openapi31x/...
go test -v ./parsers/internal/shared/...
```

## Implementation Details

- [IMPLEMENTATION_DECISIONS.md](IMPLEMENTATION_DECISIONS.md) — Design patterns and optimizations
- [RESOLVER.md](RESOLVER.md) — Deep dive into `$ref` resolution and circular detection
