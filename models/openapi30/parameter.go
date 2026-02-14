package openapi30

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// Parameter describes a single operation parameter.
// https://spec.openapis.org/oas/v3.0.3#parameter-object
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
	schema          *RefSchema
	example         interface{}
	examples        map[string]*RefExample
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
func (p *Parameter) Schema() *RefSchema               { return p.schema }
func (p *Parameter) Example() interface{}             { return p.example }
func (p *Parameter) Examples() map[string]*RefExample { return p.examples }
func (p *Parameter) Content() map[string]*MediaType   { return p.content }

func (p *Parameter) SetName(name string) error {
	if err := p.Trix.RunHooks("name", p.name, name); err != nil {
		return err
	}
	p.name = name
	return nil
}
func (p *Parameter) SetIn(in string) error {
	if err := p.Trix.RunHooks("in", p.in, in); err != nil {
		return err
	}
	p.in = in
	return nil
}
func (p *Parameter) SetDescription(description string) error {
	if err := p.Trix.RunHooks("description", p.description, description); err != nil {
		return err
	}
	p.description = description
	return nil
}
func (p *Parameter) SetRequired(required bool) error {
	if err := p.Trix.RunHooks("required", p.required, required); err != nil {
		return err
	}
	p.required = required
	return nil
}
func (p *Parameter) SetDeprecated(deprecated bool) error {
	if err := p.Trix.RunHooks("deprecated", p.deprecated, deprecated); err != nil {
		return err
	}
	p.deprecated = deprecated
	return nil
}
func (p *Parameter) SetAllowEmptyValue(allowEmptyValue bool) error {
	if err := p.Trix.RunHooks("allowEmptyValue", p.allowEmptyValue, allowEmptyValue); err != nil {
		return err
	}
	p.allowEmptyValue = allowEmptyValue
	return nil
}
func (p *Parameter) SetStyle(style string) error {
	if err := p.Trix.RunHooks("style", p.style, style); err != nil {
		return err
	}
	p.style = style
	return nil
}
func (p *Parameter) SetExplode(explode *bool) error {
	if err := p.Trix.RunHooks("explode", p.explode, explode); err != nil {
		return err
	}
	p.explode = explode
	return nil
}
func (p *Parameter) SetAllowReserved(allowReserved bool) error {
	if err := p.Trix.RunHooks("allowReserved", p.allowReserved, allowReserved); err != nil {
		return err
	}
	p.allowReserved = allowReserved
	return nil
}
func (p *Parameter) SetSchema(schema *RefSchema) error {
	if err := p.Trix.RunHooks("schema", p.schema, schema); err != nil {
		return err
	}
	p.schema = schema
	return nil
}
func (p *Parameter) SetExample(example interface{}) error {
	if err := p.Trix.RunHooks("example", p.example, example); err != nil {
		return err
	}
	p.example = example
	return nil
}
func (p *Parameter) SetExamples(examples map[string]*RefExample) error {
	if err := p.Trix.RunHooks("examples", p.examples, examples); err != nil {
		return err
	}
	p.examples = examples
	return nil
}
func (p *Parameter) SetContent(content map[string]*MediaType) error {
	if err := p.Trix.RunHooks("content", p.content, content); err != nil {
		return err
	}
	p.content = content
	return nil
}

// NewParameter creates a new Parameter instance.
func NewParameter(
	name, in, description string, required, deprecated, allowEmptyValue bool,
	style string, explode *bool, allowReserved bool,
	schema *RefSchema, example interface{}, examples map[string]*RefExample,
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

// MarshalFields implements shared.MarshalFieldsProvider for export.
func (p *Parameter) MarshalFields() []shared.Field { return p.marshalFields() }

func (p *Parameter) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(p.marshalFields())
}

func (p *Parameter) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(p.marshalFields())
}

var _ yaml.Marshaler = (*Parameter)(nil)
