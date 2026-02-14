package openapi30

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// Components holds reusable objects for the specification.
// https://spec.openapis.org/oas/v3.0.3#components-object
type Components struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	schemas         map[string]*shared.Ref[Schema]
	responses       map[string]*shared.Ref[Response]
	parameters      map[string]*shared.Ref[Parameter]
	examples        map[string]*shared.Ref[Example]
	requestBodies   map[string]*shared.Ref[RequestBody]
	headers         map[string]*shared.Ref[Header]
	securitySchemes map[string]*shared.Ref[SecurityScheme]
	links           map[string]*shared.Ref[Link]
	callbacks       map[string]*shared.Ref[Callback]
}

func (c *Components) Schemas() map[string]*shared.Ref[Schema]            { return c.schemas }
func (c *Components) Responses() map[string]*shared.Ref[Response]        { return c.responses }
func (c *Components) Parameters() map[string]*shared.Ref[Parameter]      { return c.parameters }
func (c *Components) Examples() map[string]*shared.Ref[Example]          { return c.examples }
func (c *Components) RequestBodies() map[string]*shared.Ref[RequestBody] { return c.requestBodies }
func (c *Components) Headers() map[string]*shared.Ref[Header]            { return c.headers }
func (c *Components) SecuritySchemes() map[string]*shared.Ref[SecurityScheme] {
	return c.securitySchemes
}
func (c *Components) Links() map[string]*shared.Ref[Link]         { return c.links }
func (c *Components) Callbacks() map[string]*shared.Ref[Callback] { return c.callbacks }

// NewComponents creates a new Components instance.
func NewComponents(
	schemas map[string]*shared.Ref[Schema],
	responses map[string]*shared.Ref[Response],
	parameters map[string]*shared.Ref[Parameter],
	examples map[string]*shared.Ref[Example],
	requestBodies map[string]*shared.Ref[RequestBody],
	headers map[string]*shared.Ref[Header],
	securitySchemes map[string]*shared.Ref[SecurityScheme],
	links map[string]*shared.Ref[Link],
	callbacks map[string]*shared.Ref[Callback],
) *Components {
	return &Components{
		schemas: schemas, responses: responses, parameters: parameters,
		examples: examples, requestBodies: requestBodies, headers: headers,
		securitySchemes: securitySchemes, links: links, callbacks: callbacks,
	}
}

func (c *Components) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "schemas", Value: c.schemas},
		{Key: "responses", Value: c.responses},
		{Key: "parameters", Value: c.parameters},
		{Key: "examples", Value: c.examples},
		{Key: "requestBodies", Value: c.requestBodies},
		{Key: "headers", Value: c.headers},
		{Key: "securitySchemes", Value: c.securitySchemes},
		{Key: "links", Value: c.links},
		{Key: "callbacks", Value: c.callbacks},
	}
	return shared.AppendExtensions(fields, c.VendorExtensions)
}

func (c *Components) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(c.marshalFields())
}

func (c *Components) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(c.marshalFields())
}

var _ yaml.Marshaler = (*Components)(nil)
