package openapi30

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// Header represents a Header Object.
// https://spec.openapis.org/oas/v3.0.3#header-object
type Header struct {
	Node // embedded - provides VendorExtensions and Trix

	description     string
	required        bool
	deprecated      bool
	allowEmptyValue bool
	style           string
	explode         *bool
	allowReserved   bool
	schema          *SchemaRef
	example         interface{}
	examples        map[string]*ExampleRef
	content         map[string]*MediaType
}

func (h *Header) Description() string              { return h.description }
func (h *Header) Required() bool                   { return h.required }
func (h *Header) Deprecated() bool                 { return h.deprecated }
func (h *Header) AllowEmptyValue() bool            { return h.allowEmptyValue }
func (h *Header) Style() string                    { return h.style }
func (h *Header) Explode() *bool                   { return h.explode }
func (h *Header) AllowReserved() bool              { return h.allowReserved }
func (h *Header) Schema() *SchemaRef               { return h.schema }
func (h *Header) Example() interface{}             { return h.example }
func (h *Header) Examples() map[string]*ExampleRef { return h.examples }
func (h *Header) Content() map[string]*MediaType   { return h.content }

// NewHeader creates a new Header instance.
func NewHeader(
	description string, required, deprecated, allowEmptyValue bool,
	style string, explode *bool, allowReserved bool,
	schema *SchemaRef, example interface{}, examples map[string]*ExampleRef,
	content map[string]*MediaType,
) *Header {
	return &Header{
		description: description, required: required, deprecated: deprecated,
		allowEmptyValue: allowEmptyValue, style: style, explode: explode,
		allowReserved: allowReserved, schema: schema, example: example,
		examples: examples, content: content,
	}
}

func (h *Header) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "description", Value: h.description},
		{Key: "required", Value: h.required},
		{Key: "deprecated", Value: h.deprecated},
		{Key: "allowEmptyValue", Value: h.allowEmptyValue},
		{Key: "style", Value: h.style},
		{Key: "explode", Value: h.explode},
		{Key: "allowReserved", Value: h.allowReserved},
		{Key: "schema", Value: h.schema},
		{Key: "example", Value: h.example},
		{Key: "examples", Value: h.examples},
		{Key: "content", Value: h.content},
	}
	return shared.AppendExtensions(fields, h.VendorExtensions)
}

func (h *Header) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(h.marshalFields())
}

func (h *Header) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(h.marshalFields())
}

var _ yaml.Marshaler = (*Header)(nil)
