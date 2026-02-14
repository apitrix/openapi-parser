package openapi30

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// Parameter describes a single operation parameter.
// https://spec.openapis.org/oas/v3.0.3#parameter-object
type Parameter struct {
	Node // embedded - provides VendorExtensions and Trix

	name            string
	in              string
	description     string
	required        bool
	deprecated      bool
	allowEmptyValue bool
	style           string
	explode         *bool
	allowReserved   bool
	schema          *shared.Ref[Schema]
	example         interface{}
	examples        map[string]*shared.Ref[Example]
	content         map[string]*MediaType
}

func (p *Parameter) Name() string                     { return p.name }
func (p *Parameter) In() string                       { return p.in }
func (p *Parameter) Description() string              { return p.description }
func (p *Parameter) Required() bool                   { return p.required }
func (p *Parameter) Deprecated() bool                 { return p.deprecated }
func (p *Parameter) AllowEmptyValue() bool            { return p.allowEmptyValue }
func (p *Parameter) Style() string                    { return p.style }
func (p *Parameter) Explode() *bool                   { return p.explode }
func (p *Parameter) AllowReserved() bool              { return p.allowReserved }
func (p *Parameter) Schema() *shared.Ref[Schema]               { return p.schema }
func (p *Parameter) Example() interface{}             { return p.example }
func (p *Parameter) Examples() map[string]*shared.Ref[Example] { return p.examples }
func (p *Parameter) Content() map[string]*MediaType   { return p.content }

// NewParameter creates a new Parameter instance.
func NewParameter(
	name, in, description string, required, deprecated, allowEmptyValue bool,
	style string, explode *bool, allowReserved bool,
	schema *shared.Ref[Schema], example interface{}, examples map[string]*shared.Ref[Example],
	content map[string]*MediaType,
) *Parameter {
	return &Parameter{
		name: name, in: in, description: description,
		required: required, deprecated: deprecated, allowEmptyValue: allowEmptyValue,
		style: style, explode: explode, allowReserved: allowReserved,
		schema: schema, example: example, examples: examples, content: content,
	}
}

func (p *Parameter) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "name", Value: p.name},
		{Key: "in", Value: p.in},
		{Key: "description", Value: p.description},
		{Key: "required", Value: p.required},
		{Key: "deprecated", Value: p.deprecated},
		{Key: "allowEmptyValue", Value: p.allowEmptyValue},
		{Key: "style", Value: p.style},
		{Key: "explode", Value: p.explode},
		{Key: "allowReserved", Value: p.allowReserved},
		{Key: "schema", Value: p.schema},
		{Key: "example", Value: p.example},
		{Key: "examples", Value: p.examples},
		{Key: "content", Value: p.content},
	}
	return shared.AppendExtensions(fields, p.VendorExtensions)
}

func (p *Parameter) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(p.marshalFields())
}

func (p *Parameter) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(p.marshalFields())
}

var _ yaml.Marshaler = (*Parameter)(nil)
