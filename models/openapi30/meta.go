package openapi30

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

	// ResolvedMapping holds discriminator mapping values resolved to schema refs (Discriminator only).
	ResolvedMapping map[string]*RefSchema `json:"-" yaml:"-"`

	// ResolvedOperation holds the operation resolved from operationRef (Link only).
	ResolvedOperation *Operation `json:"-" yaml:"-"`
}

// Ref type aliases — use these instead of shared.Ref[T] throughout the package.
type RefSchema = shared.Ref[Schema]
type RefParameter = shared.Ref[Parameter]
type RefResponse = shared.Ref[Response]
type RefHeader = shared.Ref[Header]
type RefExample = shared.Ref[Example]
type RefRequestBody = shared.Ref[RequestBody]
type RefSecurityScheme = shared.Ref[SecurityScheme]
type RefLink = shared.Ref[Link]
type RefCallback = shared.Ref[Callback]
type RefPathItem = shared.Ref[PathItem]

func NewRefSchema(ref string) *RefSchema                 { return shared.NewRef[Schema](ref) }
func NewRefParameter(ref string) *RefParameter           { return shared.NewRef[Parameter](ref) }
func NewRefResponse(ref string) *RefResponse             { return shared.NewRef[Response](ref) }
func NewRefHeader(ref string) *RefHeader                 { return shared.NewRef[Header](ref) }
func NewRefExample(ref string) *RefExample               { return shared.NewRef[Example](ref) }
func NewRefRequestBody(ref string) *RefRequestBody       { return shared.NewRef[RequestBody](ref) }
func NewRefSecurityScheme(ref string) *RefSecurityScheme { return shared.NewRef[SecurityScheme](ref) }
func NewRefLink(ref string) *RefLink                     { return shared.NewRef[Link](ref) }
func NewRefCallback(ref string) *RefCallback             { return shared.NewRef[Callback](ref) }
func NewRefPathItem(ref string) *RefPathItem             { return shared.NewRef[PathItem](ref) }

// ElementBase is embedded in all v30 types to provide vendor extensions and library metadata.
type ElementBase struct {
	VendorExtensions map[string]interface{} `json:"-" yaml:"-"`
	Trix             Trix                   `json:"-" yaml:"-"`
}
