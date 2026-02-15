package openapi30

import (
	"github.com/apitrix/openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// MediaType provides schema and examples for a media type.
// https://spec.openapis.org/oas/v3.0.3#media-type-object
type MediaType struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	schema   *RefSchema
	example  interface{}
	examples map[string]*RefExample
	encoding map[string]*Encoding
}

func (mt *MediaType) Schema() *RefSchema               { return mt.schema }
func (mt *MediaType) Example() interface{}             { return mt.example }
func (mt *MediaType) Examples() map[string]*RefExample { return mt.examples }
func (mt *MediaType) Encoding() map[string]*Encoding   { return mt.encoding }

func (mt *MediaType) SetSchema(schema *RefSchema) error {
	if err := mt.Trix.RunHooks("schema", mt.schema, schema); err != nil {
		return err
	}
	mt.schema = schema
	return nil
}
func (mt *MediaType) SetExample(example interface{}) error {
	if err := mt.Trix.RunHooks("example", mt.example, example); err != nil {
		return err
	}
	mt.example = example
	return nil
}
func (mt *MediaType) SetExamples(examples map[string]*RefExample) error {
	if err := mt.Trix.RunHooks("examples", mt.examples, examples); err != nil {
		return err
	}
	mt.examples = examples
	return nil
}
func (mt *MediaType) SetEncoding(encoding map[string]*Encoding) error {
	if err := mt.Trix.RunHooks("encoding", mt.encoding, encoding); err != nil {
		return err
	}
	mt.encoding = encoding
	return nil
}

// NewMediaType creates a new MediaType instance.
func NewMediaType(schema *RefSchema, example interface{}, examples map[string]*RefExample, encoding map[string]*Encoding) *MediaType {
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

// MarshalFields implements shared.MarshalFieldsProvider for export.
func (mt *MediaType) MarshalFields() []shared.Field { return mt.marshalFields() }

func (mt *MediaType) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(mt.marshalFields())
}

func (mt *MediaType) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(mt.marshalFields())
}

var _ yaml.Marshaler = (*MediaType)(nil)
