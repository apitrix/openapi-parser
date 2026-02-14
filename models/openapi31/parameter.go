package openapi31

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// Parameter describes a single operation parameter.
// https://spec.openapis.org/oas/v3.1.0#parameter-object
type Parameter struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	name            string
	in              string
	description     string
	required        bool
	deprecated      bool
	allowEmptyValue bool
	style           string
	explode         *bool
	allowReserved   bool
	schema          *shared.RefWithMeta[Schema]
	example         interface{}
	examples        map[string]*shared.RefWithMeta[Example]
	content         map[string]*MediaType
}

func (p *Parameter) Name() string                                      { return p.name }
func (p *Parameter) In() string                                        { return p.in }
func (p *Parameter) Description() string                               { return p.description }
func (p *Parameter) Required() bool                                    { return p.required }
func (p *Parameter) Deprecated() bool                                  { return p.deprecated }
func (p *Parameter) AllowEmptyValue() bool                             { return p.allowEmptyValue }
func (p *Parameter) Style() string                                     { return p.style }
func (p *Parameter) Explode() *bool                                    { return p.explode }
func (p *Parameter) AllowReserved() bool                               { return p.allowReserved }
func (p *Parameter) Schema() *shared.RefWithMeta[Schema]               { return p.schema }
func (p *Parameter) Example() interface{}                              { return p.example }
func (p *Parameter) Examples() map[string]*shared.RefWithMeta[Example] { return p.examples }
func (p *Parameter) Content() map[string]*MediaType                    { return p.content }

// ParameterFields holds all fields for constructing a Parameter.
type ParameterFields struct {
	Name            string
	In              string
	Description     string
	Required        bool
	Deprecated      bool
	AllowEmptyValue bool
	Style           string
	Explode         *bool
	AllowReserved   bool
	Schema          *shared.RefWithMeta[Schema]
	Example         interface{}
	Examples        map[string]*shared.RefWithMeta[Example]
	Content         map[string]*MediaType
}

// NewParameter creates a new Parameter instance.
func NewParameter(f ParameterFields) *Parameter {
	return &Parameter{
		name: f.Name, in: f.In, description: f.Description,
		required: f.Required, deprecated: f.Deprecated,
		allowEmptyValue: f.AllowEmptyValue, style: f.Style,
		explode: f.Explode, allowReserved: f.AllowReserved,
		schema: f.Schema, example: f.Example,
		examples: f.Examples, content: f.Content,
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
