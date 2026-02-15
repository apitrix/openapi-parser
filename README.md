# OpenAPI Parser

Go library for parsing OpenAPI 2.0, 3.0, and 3.1 specifications.

```bash
go get github.com/apitrix/openapi-parser
```

## Parse

```go
result, _ := openapi30x.Parse(data)          // from bytes
result, _ := openapi30x.ParseFile("api.yaml") // from file (resolves $ref)
result, _ := openapi30x.ParseReader(r)        // from io.Reader
```

All parsers share the same API: `openapi20`, `openapi30x`, `openapi31x`.

## Config

`nil` config enables everything. Use flags to control individual features:

```go
result, _ := openapi30x.Parse(data, &openapi30x.ParseConfig{
    ResolveInternalRefs: true,
    ResolveExternalRefs: true,
    DetectUnknownFields: true,
    ApplySpecDefaults:   true,
})
```

## Source locations

Every element carries line/column via `Trix.Source`:

```go
info := result.Document.Info()
fmt.Println(info.Trix.Source.Start.Line, info.Trix.Source.Start.Column)
```

## $ref resolution

Refs resolve in the background. `Value()` blocks until resolved:

```go
result, _ := openapi30x.ParseFile("api.yaml")
result.Wait() // wait for all refs

schemaRef := mediaType.Schema()
schemaRef.Ref         // "$ref" string, e.g. "#/components/schemas/Pet"
schemaRef.Value()     // resolved *Schema (blocks until ready)
schemaRef.Circular()  // true if circular reference detected
schemaRef.ResolveErr() // resolution error, if any
```

See [docs/references.md](docs/references.md) and [docs/background_resolve.md](docs/background_resolve.md).

## Unknown field detection

Fields not in the OpenAPI spec appear in `result.Errors` with `Kind == "unknown_field"`:

```go
result, _ := openapi30x.Parse(data) // DetectUnknownFields on by default
for _, e := range result.Errors {
    if e.Kind == "unknown_field" {
        fmt.Println(e.Path, e.Message, e.Line)
    }
}
```

See [docs/unknown_fields.md](docs/unknown_fields.md).

## Spec defaults

When enabled, fills in spec-defined defaults (e.g. missing `servers` becomes `[{url: "/"}]`):

```go
result, _ := openapi30x.Parse(data) // ApplySpecDefaults on by default
servers := result.Document.Servers()
// servers[0].URL() == "/" when servers was absent
```

See [docs/defaults.md](docs/defaults.md).

## Vendor extensions

`x-*` fields are captured on every element:

```go
info := result.Document.Info()
fmt.Println(info.VendorExtensions["x-custom"])
```

## Setters and hooks

All fields have setters. Hooks run before mutation and can reject changes:

```go
info := result.Document.Info()
info.Trix.OnSet("title", func(field string, oldVal, newVal any) error {
    if newVal.(string) == "" {
        return fmt.Errorf("title cannot be empty")
    }
    return nil
})
err := info.SetTitle("New Title")
```

See [docs/setters-and-hooks.md](docs/setters-and-hooks.md).

## Errors

All errors (parse, unknown fields, resolve) are in `result.Errors`:

```go
for _, e := range result.Errors {
    fmt.Println(e.Kind, e.Path, e.Message) // "error", "unknown_field", or "resolve_error"
}
```

## License

MIT
