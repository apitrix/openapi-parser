package openapi31

// Parameter describes a single operation parameter.
// https://spec.openapis.org/oas/v3.1.0#parameter-object
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
	schema          *SchemaRef
	example         interface{}
	examples        map[string]*ExampleRef
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
func (p *Parameter) Schema() *SchemaRef               { return p.schema }
func (p *Parameter) Example() interface{}             { return p.example }
func (p *Parameter) Examples() map[string]*ExampleRef { return p.examples }
func (p *Parameter) Content() map[string]*MediaType   { return p.content }

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
	Schema          *SchemaRef
	Example         interface{}
	Examples        map[string]*ExampleRef
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
