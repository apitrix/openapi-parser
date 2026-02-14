# ApplySpecDefaults

When `ApplySpecDefaults` is enabled in `ParseConfig`, the parser fills in OpenAPI-specified default values for optional fields that are absent in the spec. This ensures the parsed model reflects the spec's intended semantics (e.g., `servers` absent means "use `/`") rather than Go zero values.

## How to Enable

- Set `ParseConfig.ApplySpecDefaults = true` explicitly, or
- Use `ParseConfig.All()` which enables all features including `ApplySpecDefaults`

```go
// Enable defaults
result, err := openapi30x.Parse(data, openapi30x.All())

// Disable defaults (preserve original behavior)
result, err := openapi30x.Parse(data, openapi30x.None())

// Custom config
cfg := &shared.ParseConfig{ApplySpecDefaults: true}
result, err := openapi30x.Parse(data, cfg)
```

Entry points: `Parse`, `ParseReader`, `ParseFile` in `openapi20`, `openapi30x`, and `openapi31x` packages.

## Implemented Defaults (Phase 1)

| Spec | Object | Field | When Absent | Default |
|------|--------|-------|-------------|---------|
| OpenAPI 3.0/3.1 | OpenAPI | `servers` | nil or empty array | `[{ url: "/" }]` |
| Swagger 2.0 | Swagger | `basePath` | absent | `/` |

## Complete List of Spec Defaults

The following tables document all defaults defined by the OpenAPI specifications. Only Phase 1 items are currently implemented.

### OpenAPI 3.0 / 3.1

| Object | Field | When Absent | Spec Default |
|--------|-------|-------------|--------------|
| **OpenAPI** | `servers` | nil or empty array | `[{ url: "/" }]` |
| **Parameter** | `required` | absent | `false` |
| **Parameter** | `deprecated` | absent | `false` |
| **Parameter** | `allowEmptyValue` | absent | `false` |
| **Parameter** | `allowReserved` | absent | `false` |
| **Parameter** | `style` | absent | path: `simple`; query: `form`; header: `simple`; cookie: `form` |
| **Parameter** | `explode` | absent | path/header: `false`; query/cookie: `true` |
| **RequestBody** | `required` | absent | `false` |
| **Schema** | `nullable` | absent | `false` |
| **Schema** | `readOnly` | absent | `false` |
| **Schema** | `writeOnly` | absent | `false` |
| **Schema** | `deprecated` | absent | `false` |
| **Schema** | `required` | absent | `[]` (empty array) |
| **Schema** | `exclusiveMaximum` | absent | `false` |
| **Schema** | `exclusiveMinimum` | absent | `false` |
| **Schema** | `uniqueItems` | absent | `false` |
| **Encoding** | `explode` | absent | `true` |
| **Encoding** | `allowReserved` | absent | `false` |
| **Encoding** | `style` | absent | `form` |
| **Header** | `required` | absent | `false` |
| **Header** | `deprecated` | absent | `false` |
| **Header** | `allowEmptyValue` | absent | `false` |
| **Header** | `allowReserved` | absent | `false` |
| **Header** | `style` | absent | `simple` |
| **Header** | `explode` | absent | `false` |
| **Operation** | `deprecated` | absent | `false` |
| **XML** | `attribute` | absent | `false` |
| **XML** | `wrapped` | absent | `false` |

### OpenAPI 2.0 (Swagger)

| Object | Field | When Absent | Spec Default |
|--------|-------|-------------|--------------|
| **Swagger** | `host` | absent | host serving the documentation |
| **Swagger** | `basePath` | absent | `/` (API served directly under host) |
| **Swagger** | `schemes` | absent | scheme used to access the Swagger definition |
| **Parameter** | `required` | absent (non-path) | `false`; path: MUST be `true` |
| **Parameter** | `allowEmptyValue` | absent | `false` |
| **Parameter** | `collectionFormat` | absent | `csv` |
| **Operation** | `deprecated` | absent | `false` |
| **Schema** | `exclusiveMaximum` | absent | `false` |
| **Schema** | `exclusiveMinimum` | absent | `false` |
| **Schema** | `uniqueItems` | absent | `false` |
| **Schema** | `readOnly` | absent | `false` |
| **XML** | `attribute` | absent | `false` |
| **XML** | `wrapped` | absent | `false` |

### Notes

- **Boolean defaults**: Most default to `false`; Go zero value already matches.
- **Arrays**: `required`, `tags`, etc. default to `[]`; parser currently may return `nil`.
- **Parameter style**: Depends on `in`; requires lookup when absent.
- **Parameter explode**: `path`/`header` → `false`; `query`/`cookie` → `true`.
- **Encoding explode**: Defaults to `true` (different from Parameter).
- **Swagger 2.0 `host`, `schemes`**: Spec says "host serving the documentation" / "scheme used to access" — these are runtime-dependent and are not applied by the parser.
- **Server Variable**: `default` is REQUIRED in the spec — no default to apply.

## Phased Rollout

- **Phase 1** (implemented): `servers` (OpenAPI 3.x), `basePath` (Swagger 2.0)
- **Phase 2** (planned): `Parameter.style`, `Encoding.style` based on `in`/`contentType`
- **Phase 3** (planned): Remaining spec defaults (arrays, etc.) as needed
