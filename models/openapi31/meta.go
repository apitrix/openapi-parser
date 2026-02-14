package openapi31

import "openapi-parser/models/shared"

// Location represents a position in the source file.
type Location = shared.Location

// NodeSource contains source location and raw parsed data for a node.
type NodeSource = shared.NodeSource

// ParseError represents a parsing error associated with a specific node.
type ParseError = shared.ParseError

// Trix contains all library-level metadata and functionality.
// Everything under Trix is provided by the apitrix library,
// not part of the OpenAPI specification itself.
type Trix struct {
	shared.Trix // promotes Source, Errors, OnSet, RunHooks

	// Resolved reference fields — populated by the resolver, not part of the spec.
	// These are nil/empty unless the resolver has run.

	// ResolvedDynamicRef holds the schema resolved from $dynamicRef (Schema only).
	ResolvedDynamicRef *RefSchema `json:"-" yaml:"-"`

	// ResolvedMapping holds discriminator mapping values resolved to schema refs (Discriminator only).
	ResolvedMapping map[string]*RefSchema `json:"-" yaml:"-"`

	// ResolvedOperation holds the operation resolved from operationRef (Link only).
	ResolvedOperation *Operation `json:"-" yaml:"-"`
}

// RefWithMeta type aliases — use these instead of shared.RefWithMeta[T] throughout the package.
type RefSchema = shared.RefWithMeta[Schema]
type RefParameter = shared.RefWithMeta[Parameter]
type RefResponse = shared.RefWithMeta[Response]
type RefHeader = shared.RefWithMeta[Header]
type RefExample = shared.RefWithMeta[Example]
type RefRequestBody = shared.RefWithMeta[RequestBody]
type RefSecurityScheme = shared.RefWithMeta[SecurityScheme]
type RefLink = shared.RefWithMeta[Link]
type RefCallback = shared.RefWithMeta[Callback]
type RefPathItem = shared.RefWithMeta[PathItem]

func NewRefSchema(ref string) *RefSchema           { return shared.NewRefWithMeta[Schema](ref) }
func NewRefParameter(ref string) *RefParameter     { return shared.NewRefWithMeta[Parameter](ref) }
func NewRefResponse(ref string) *RefResponse       { return shared.NewRefWithMeta[Response](ref) }
func NewRefHeader(ref string) *RefHeader           { return shared.NewRefWithMeta[Header](ref) }
func NewRefExample(ref string) *RefExample         { return shared.NewRefWithMeta[Example](ref) }
func NewRefRequestBody(ref string) *RefRequestBody { return shared.NewRefWithMeta[RequestBody](ref) }
func NewRefSecurityScheme(ref string) *RefSecurityScheme {
	return shared.NewRefWithMeta[SecurityScheme](ref)
}
func NewRefLink(ref string) *RefLink         { return shared.NewRefWithMeta[Link](ref) }
func NewRefCallback(ref string) *RefCallback { return shared.NewRefWithMeta[Callback](ref) }
func NewRefPathItem(ref string) *RefPathItem { return shared.NewRefWithMeta[PathItem](ref) }

// ElementBase is embedded in all v31 types to provide vendor extensions and library metadata.
type ElementBase struct {
	VendorExtensions map[string]interface{} `json:"-" yaml:"-"`
	Trix             Trix                   `json:"-" yaml:"-"`
}
