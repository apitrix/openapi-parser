package openapi31

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// MediaType provides schema and examples for a media type.
// https://spec.openapis.org/oas/v3.1.0#media-type-object
type MediaType struct {
	Node // embedded - provides VendorExtensions and Trix

	schema   *SchemaRef
	example  interface{}
	examples map[string]*ExampleRef
	encoding map[string]*Encoding
}

func (m *MediaType) Schema() *SchemaRef               { return m.schema }
func (m *MediaType) Example() interface{}             { return m.example }
func (m *MediaType) Examples() map[string]*ExampleRef { return m.examples }
func (m *MediaType) Encoding() map[string]*Encoding   { return m.encoding }

// NewMediaType creates a new MediaType instance.
func NewMediaType(schema *SchemaRef, example interface{}, examples map[string]*ExampleRef, encoding map[string]*Encoding) *MediaType {
	return &MediaType{schema: schema, example: example, examples: examples, encoding: encoding}
}

func (m *MediaType) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "schema", Value: m.schema},
		{Key: "example", Value: m.example},
		{Key: "examples", Value: m.examples},
		{Key: "encoding", Value: m.encoding},
	}
	return shared.AppendExtensions(fields, m.VendorExtensions)
}

func (m *MediaType) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(m.marshalFields())
}

func (m *MediaType) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(m.marshalFields())
}

var _ yaml.Marshaler = (*MediaType)(nil)
