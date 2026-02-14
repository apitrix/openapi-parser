package openapi30

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// MediaType provides schema and examples for a media type.
// https://spec.openapis.org/oas/v3.0.3#media-type-object
type MediaType struct {
	Node // embedded - provides VendorExtensions and Trix

	schema   *shared.Ref[Schema]
	example  interface{}
	examples map[string]*shared.Ref[Example]
	encoding map[string]*Encoding
}

func (mt *MediaType) Schema() *shared.Ref[Schema]               { return mt.schema }
func (mt *MediaType) Example() interface{}             { return mt.example }
func (mt *MediaType) Examples() map[string]*shared.Ref[Example] { return mt.examples }
func (mt *MediaType) Encoding() map[string]*Encoding   { return mt.encoding }

// NewMediaType creates a new MediaType instance.
func NewMediaType(schema *shared.Ref[Schema], example interface{}, examples map[string]*shared.Ref[Example], encoding map[string]*Encoding) *MediaType {
	return &MediaType{schema: schema, example: example, examples: examples, encoding: encoding}
}

func (mt *MediaType) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "schema", Value: mt.schema},
		{Key: "example", Value: mt.example},
		{Key: "examples", Value: mt.examples},
		{Key: "encoding", Value: mt.encoding},
	}
	return shared.AppendExtensions(fields, mt.VendorExtensions)
}

func (mt *MediaType) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(mt.marshalFields())
}

func (mt *MediaType) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(mt.marshalFields())
}

var _ yaml.Marshaler = (*MediaType)(nil)
