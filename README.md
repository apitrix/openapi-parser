# OpenAPI Parser

A Go library for parsing OpenAPI 3.0 specifications with full source location metadata and unknown field detection.

## Features

- **OpenAPI 3.0 support** - Full parsing of OpenAPI 3.0.x specifications
- **JSON and YAML** - Parses both formats seamlessly
- **Source location metadata** - Line and column numbers preserved for every parsed element
- **Unknown field detection** - Identifies typos, misplaced properties, and unsupported fields
- **Extension support** - Captures `x-*` extension fields
- **Reference types** - Handles `$ref` references and inline objects
- **Raw data access** - Original parsed structure available for tooling

## Installation

```bash
go get openapi-parser
```

## Quick Start

```go
package main

import (
    "fmt"
    "os"

    "openapi-parser/parsers/openapi30"
)

func main() {
    data, _ := os.ReadFile("openapi.yaml")
    
    doc, err := openapi30.Parse(data)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Title: %s\n", doc.Info.Title)
    fmt.Printf("Version: %s\n", doc.Info.Version)
    
    // Source location is available on every element
    fmt.Printf("Info defined at line %d\n", doc.Info.NodeSource.Start.Line)
}
```

## API

### Parse Functions

| Function | Description |
|----------|-------------|
| `Parse(data []byte)` | Parse OpenAPI 3.0 from bytes |
| `ParseReader(r io.Reader)` | Parse from an io.Reader |
| `ParseWithUnknownFields(data []byte)` | Parse and detect unknown fields |
| `ParseReaderWithUnknownFields(r io.Reader)` | Parse from reader with unknown field detection |

### Detecting Unknown Fields

The parser can identify fields that don't belong in the OpenAPI specification:

```go
result, err := openapi30.ParseWithUnknownFields(data)
if err != nil {
    log.Fatal(err)
}

for _, field := range result.UnknownFields {
    fmt.Printf("Unknown field '%s' at %s (line %d)\n",
        field.Key, field.Path, field.Line)
}
```

See [docs/unknown_fields.md](docs/unknown_fields.md) for detailed documentation.

## Project Structure

```
openapi-parser/
├── models/
│   ├── openapi20/          # Swagger 2.0 model types
│   └── openapi30/          # OpenAPI 3.0 model types
├── parsers/
│   └── openapi30/          # OpenAPI 3.0 parser
│       └── testdata/       # Test fixtures
└── docs/                   # Documentation
```

## Source Location Metadata

Every parsed object includes source location information:

```go
type Node struct {
    NodeSource NodeSource             // Location and raw data
    Extensions map[string]interface{} // Extension fields (x-*)
}

type NodeSource struct {
    Start Location    // Start line/column (1-based)
    End   Location    // End line/column (1-based)
    Raw   interface{} // Original parsed data
}
```

This enables:
- IDE integration with go-to-definition
- Precise error messages with line numbers
- Linting and validation tools

## Reference Types

Fields that can be `$ref` or inline use wrapper types:

```go
// Example: SchemaRef can be a $ref string or inline Schema
schema := operation.RequestBody.Value.Content["application/json"].Schema

if schema.Ref != "" {
    fmt.Println("Reference to:", schema.Ref)
} else {
    fmt.Println("Inline schema type:", schema.Value.Type)
}
```

## Dependencies

- [gopkg.in/yaml.v3](https://gopkg.in/yaml.v3)

## License

MIT License
