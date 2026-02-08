# OpenAPI Parser

A Go library for parsing OpenAPI 2.0, 3.0, and 3.1 specifications with full source location metadata, `$ref` resolution, circular reference detection, and unknown field detection.

## Features

- **Multi-version support** — OpenAPI 2.0 (Swagger), 3.0.x, and 3.1.x / 3.2.x
- **JSON and YAML** — Parses both formats seamlessly via `yaml.v3`
- **`$ref` resolution** — Local JSON pointers and external file references with caching
- **Circular reference detection** — Marks self-referencing schemas with `Circular = true` instead of infinite recursion
- **Source location metadata** — Line and column numbers preserved for every parsed element
- **Unknown field detection** — Identifies typos, misplaced properties, and unsupported fields
- **Extension support** — Captures `x-*` extension fields
- **Reference types** — Typed `*Ref` wrappers distinguish `$ref` strings from inline objects
- **Raw data access** — Original parsed structure available for tooling

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

    "openapi-parser/parsers/openapi30x"
)

func main() {
    data, _ := os.ReadFile("openapi.yaml")

    doc, err := openapi30x.Parse(data)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Title: %s\n", doc.Info.Title)
    fmt.Printf("Version: %s\n", doc.Info.Version)

    // Source location is available on every element
    fmt.Printf("Info defined at line %d\n", doc.Info.NodeSource.Start.Line)
}
```

### Parsing from a File (with `$ref` Resolution)

`ParseFile` reads the file, parses it, and resolves all `$ref` references (local and external) relative to the file's directory:

```go
doc, err := openapi30x.ParseFile("./specs/openapi.yaml")
```

This works identically across all three parser versions.

## API

### Parse Functions

Each parser package (`openapi20`, `openapi30x`, `openapi31x`) exposes the same surface:

| Function | Description |
|----------|-------------|
| `Parse(data []byte)` | Parse from bytes |
| `ParseReader(r io.Reader)` | Parse from an io.Reader |
| `ParseWithUnknownFields(data []byte)` | Parse and detect unknown fields |
| `ParseReaderWithUnknownFields(r io.Reader)` | Parse from reader with unknown field detection |
| `ParseFile(pathOrURL string)` | Parse from file path or URL with full `$ref` resolution |

### Detecting Unknown Fields

```go
result, err := openapi30x.ParseWithUnknownFields(data)
if err != nil {
    log.Fatal(err)
}

for _, field := range result.UnknownFields {
    fmt.Printf("Unknown field '%s' at %s (line %d)\n",
        field.Key, field.Path, field.Line)
}
```

### Reference Types

Fields that can be `$ref` or inline use typed wrapper structs:

```go
schema := operation.RequestBody.Value.Content["application/json"].Schema

if schema.Ref != "" {
    fmt.Println("Reference to:", schema.Ref)
} else {
    fmt.Println("Inline schema type:", schema.Value.Type)
}

// Circular references are flagged, not followed
if schema.Circular {
    fmt.Println("This is a circular reference")
}
```

## Project Structure

```
openapi-parser/
├── models/
│   ├── openapi20/          # Swagger 2.0 model types
│   ├── openapi30/          # OpenAPI 3.0 model types
│   └── openapi31/          # OpenAPI 3.1 model types
├── parsers/
│   ├── internal/shared/    # Shared utilities, RefResolver, node helpers
│   ├── openapi20/          # OpenAPI 2.0 parser
│   ├── openapi30x/         # OpenAPI 3.0.x parser
│   └── openapi31x/         # OpenAPI 3.1.x / 3.2.x parser
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

This enables IDE integration with go-to-definition, precise error messages, and linting tools.

## Dependencies

- [gopkg.in/yaml.v3](https://gopkg.in/yaml.v3) — YAML parsing with node-level source locations
- [github.com/spf13/afero](https://github.com/spf13/afero) — Filesystem abstraction for testable file I/O
- [github.com/stretchr/testify](https://github.com/stretchr/testify) — Test assertions

## License

MIT License
