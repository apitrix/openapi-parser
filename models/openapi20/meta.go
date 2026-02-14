package openapi20

import "openapi-parser/models/shared"

// Type aliases for shared meta types.
type Location = shared.Location
type NodeSource = shared.NodeSource
type ParseError = shared.ParseError
type Trix = shared.Trix

// ElementBase is embedded in all v20 types to provide vendor extensions and library metadata.
type ElementBase = shared.ElementBase

// Ref type aliases — use these instead of shared.Ref[T] throughout the package.
type RefSchema = shared.Ref[Schema]
type RefParameter = shared.Ref[Parameter]
type RefResponse = shared.Ref[Response]

// NewRefSchema creates a Ref to a Schema.
func NewRefSchema(ref string) *RefSchema { return shared.NewRef[Schema](ref) }

// NewRefParameter creates a Ref to a Parameter.
func NewRefParameter(ref string) *RefParameter { return shared.NewRef[Parameter](ref) }

// NewRefResponse creates a Ref to a Response.
func NewRefResponse(ref string) *RefResponse { return shared.NewRef[Response](ref) }
