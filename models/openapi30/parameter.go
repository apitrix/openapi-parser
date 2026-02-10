package openapi30

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

// NewParameter creates a new Parameter instance.
func NewParameter(
	name, in, description string, required, deprecated, allowEmptyValue bool,
	style string, explode *bool, allowReserved bool,
	schema *SchemaRef, example interface{}, examples map[string]*ExampleRef,
	content map[string]*MediaType,
) *Parameter {
	return &Parameter{
		name: name, in: in, description: description,
		required: required, deprecated: deprecated, allowEmptyValue: allowEmptyValue,
		style: style, explode: explode, allowReserved: allowReserved,
		schema: schema, example: example, examples: examples, content: content,
	}
}
