package openapi31

import (
	"github.com/apitrix/openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// MediaType provides schema and examples for a media type.
// https://spec.openapis.org/oas/v3.1.0#media-type-object
type MediaType struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	schema   *RefSchema
	example  interface{}
	examples map[string]*RefExample
	encoding map[string]*Encoding
}

func (m *MediaType) Schema() *RefSchema               { return m.schema }
func (m *MediaType) Example() interface{}             { return m.example }
func (m *MediaType) Examples() map[string]*RefExample { return m.examples }
func (m *MediaType) Encoding() map[string]*Encoding   { return m.encoding }

func (m *MediaType) SetSchema(schema *RefSchema) error {
	if err := m.Trix.RunHooks("schema", m.schema, schema); err != nil {
		return err
	}
	m.schema = schema
	return nil
}
func (m *MediaType) SetExample(example interface{}) error {
	if err := m.Trix.RunHooks("example", m.example, example); err != nil {
		return err
	}
	m.example = example
	return nil
}
func (m *MediaType) SetExamples(examples map[string]*RefExample) error {
	if err := m.Trix.RunHooks("examples", m.examples, examples); err != nil {
		return err
	}
	m.examples = examples
	return nil
}
func (m *MediaType) SetEncoding(encoding map[string]*Encoding) error {
	if err := m.Trix.RunHooks("encoding", m.encoding, encoding); err != nil {
		return err
	}
	m.encoding = encoding
	return nil
}

// NewMediaType creates a new MediaType instance.
func NewMediaType(schema *RefSchema, example interface{}, examples map[string]*RefExample, encoding map[string]*Encoding) *MediaType {
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

// MarshalFields implements shared.MarshalFieldsProvider for export.
func (m *MediaType) MarshalFields() []shared.Field { return m.marshalFields() }

func (m *MediaType) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(m.marshalFields())
}

func (m *MediaType) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(m.marshalFields())
}

var _ yaml.Marshaler = (*MediaType)(nil)
