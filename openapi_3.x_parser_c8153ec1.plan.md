---
name: OpenAPI 3.x Parser
overview: Implement OpenAPI 3.1 support as a separate `openapi31` package (models + parser), following the existing per-version package pattern. Extract common parsing logic into a shared internal package to reduce duplication with the existing 3.0 parser.
todos:
  - id: models-31
    content: Create models/openapi31/ package -- copy models/openapi30/, then modify schema.go (type arrays, new JSON Schema keywords, remove nullable), openapi.go (webhooks, jsonSchemaDialect), info.go (summary), refs.go (summary/description on Ref), components.go (pathItems)
    status: pending
  - id: shared-parsers
    content: Optionally extract shared 3.0/3.1 parsing logic into parsers/internal/openapi3shared/ for identical objects (server, operation, mediatype, encoding, oauth, tag, externaldocs, security)
    status: pending
  - id: parser-core
    content: "Create parsers/openapi31/ -- core files: parse.go, context.go, node_helpers.go, errors.go, unknown_fields.go (copy from openapi30 with import changes)"
    status: pending
  - id: parser-schema
    content: Implement schema.go + schema_*.go for JSON Schema Draft 2020-12 -- handle type arrays, exclusiveMin/Max as numbers, new keywords (const, if/then/else, prefixItems, etc.)
    status: pending
  - id: parser-root
    content: Implement openapi.go (version check 3.1.x, webhooks, jsonSchemaDialect), components.go (pathItems), info.go (summary), info_license.go (identifier)
    status: pending
  - id: parser-refs
    content: Update all ref_*.go files (10) to handle summary/description alongside $ref
    status: pending
  - id: parser-known-fields
    content: Update known_fields.go with all 3.1 field lists (add new schema keywords, webhooks, jsonSchemaDialect, license identifier, info summary, ref summary/description, components pathItems)
    status: pending
  - id: tests
    content: Create tests for openapi31 parser -- version validation, schema type arrays, nullable removal, webhooks, new JSON Schema keywords, ref summary/description
    status: pending
  - id: update-docs
    content: Update parsers/README.md and IMPLEMENTATION_DECISIONS.md to document the openapi31 parser
    status: pending
isProject: false
---

# OpenAPI 3.x Parser Implementation Plan

## Version Landscape

- **OpenAPI 3.0.x** -- already supported
- **OpenAPI 3.1.x** -- first target; has **breaking schema changes** vs 3.0
- **OpenAPI 3.2.0** -- backwards-compatible with 3.1, additive features only; can be handled later as an extension of the 3.1 parser or a thin separate package
- **OpenAPI 3.5** -- does not exist as a released spec version

## Recommendation: Separate `openapi31` package

The 3.1 parser should be a **separate package**, not an extension of 3.0. Key reasons:

1. **Schema Object has incompatible Go types** -- `type` changes from `string` to `string | []string`, `exclusiveMinimum`/`exclusiveMaximum` change from `bool` to `*float64`, `nullable` is removed entirely. These require different struct definitions.
2. **Reference Object gains fields** -- `summary` and `description` alongside `$ref`, affecting all `*Ref` types.
3. **New top-level objects** -- `webhooks`, `jsonSchemaDialect`.
4. **Consistent with existing pattern** -- `openapi20` and `openapi30` are separate packages.
5. **Type safety** -- separate model types prevent accidental cross-version mixing.

However, **~60% of objects are structurally identical** between 3.0 and 3.1 (Server, Operation, Parameter, Response, MediaType, Encoding, etc.). To avoid duplication, shared parsing logic should be extracted.

## Key 3.0 vs 3.1 Differences Affecting Code

### Schema Object (biggest impact -- ~6 parser files)


| Field              | 3.0           | 3.1                                                                                                                                                                                                             |
| ------------------ | ------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `type`             | `string`      | `string` or `[]string` (e.g. `["string", "null"]`)                                                                                                                                                              |
| `nullable`         | `bool`        | removed (use type array)                                                                                                                                                                                        |
| `exclusiveMinimum` | `bool`        | `*float64`                                                                                                                                                                                                      |
| `exclusiveMaximum` | `bool`        | `*float64`                                                                                                                                                                                                      |
| `example`          | `interface{}` | deprecated, replaced by `examples` (array)                                                                                                                                                                      |
| New keywords       | --            | `const`, `if`/`then`/`else`, `dependentSchemas`, `prefixItems`, `$anchor`, `$dynamicRef`, `$dynamicAnchor`, `contentEncoding`, `contentMediaType`, `contentSchema`, `unevaluatedItems`, `unevaluatedProperties` |


### Root/Structural Changes

- Root: new `webhooks` (map of PathItemRef), new `jsonSchemaDialect` (string)
- License: new `identifier` field (SPDX)
- Info: new `summary` field
- Reference Object: new `summary`, `description` fields alongside `$ref` -- affects all `ref_*.go` parsers
- Components: new `pathItems` section

### Unchanged Objects (can share parsing logic)

Server, ServerVariable, Operation, Parameter (structure), MediaType, Encoding, Response (structure), Header, ExternalDocs, Tag, Example, OAuthFlows, OAuthFlow, SecurityScheme, Discriminator, XML, SecurityRequirement

## Implementation Structure

### Phase 1: Models (`models/openapi31/`)

Create new model package mirroring `models/openapi30/` (18 files). Most files are near-identical copies with targeted changes:

- `schema.go` -- rewrite with new JSON Schema Draft 2020-12 fields
- `openapi.go` -- add `Webhooks`, `JsonSchemaDialect`
- `info.go` -- add `Summary` field
- `refs.go` -- add `Summary`, `Description` to all Ref types
- `components.go` -- add `PathItems` map
- Remaining files: copy from 3.0 with package rename (identical structs)

### Phase 2: Shared Parser Internals (`parsers/internal/openapi3shared/`)

Extract parsing functions for objects identical across 3.0 and 3.1 to reduce duplication. These are generic over the model types since they use the same field names:

- Server, ServerVariable parsing
- Operation parsing (minus schema-related delegation)
- MediaType, Encoding parsing
- OAuthFlows, OAuthFlow parsing
- Tag, ExternalDocs parsing
- SecurityRequirement, SecurityScheme parsing

This is optional but recommended for long-term maintainability. The alternative is to copy all parsers from 3.0 and modify the differences (faster to implement, more duplication).

### Phase 3: Parser (`parsers/openapi31/`)

Create new parser package mirroring `parsers/openapi30/` (~59 files). Key changes:

- `openapi.go` -- version check `"3.1."`, parse `webhooks`, `jsonSchemaDialect`
- `schema.go` + `schema_*.go` -- rewrite for JSON Schema 2020-12
- `ref_*.go` (10 files) -- handle `summary`/`description` alongside `$ref`
- `known_fields.go` -- update all field lists for 3.1 spec
- `components.go` -- add `pathItems` parsing
- `info.go` -- parse `summary`
- `info_license.go` -- parse `identifier`
- `context.go`, `parse.go`, `node_helpers.go`, `errors.go`, `unknown_fields.go` -- copy with package/import rename

### Phase 4: Tests (`parsers/openapi31/`)

- Copy test structure from `parsers/openapi30/`
- Add 3.1-specific test cases: type arrays, removed nullable, webhooks, new schema keywords, ref summary/description
- Test version validation (`"3.1."` prefix)

## Regarding OpenAPI 3.2

OpenAPI 3.2 is **fully backwards compatible** with 3.1. Two approaches:

- **Recommended**: Handle 3.2 within the `openapi31` parser (rename to `openapi3` or keep as `openapi31` accepting `3.1.` and `3.2.`), adding 3.2 fields as optional. Similar to how `openapi30` handles all `3.0.x`.
- **Alternative**: Separate `openapi32` package, very thin, mostly delegating to shared code.

This can be decided later; the 3.1 parser should be built first.