package openapi30

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// Components holds reusable objects for the specification.
// https://spec.openapis.org/oas/v3.0.3#components-object
type Components struct {
	Node // embedded - provides VendorExtensions and Trix

	schemas         map[string]*SchemaRef
	responses       map[string]*ResponseRef
	parameters      map[string]*ParameterRef
	examples        map[string]*ExampleRef
	requestBodies   map[string]*RequestBodyRef
	headers         map[string]*HeaderRef
	securitySchemes map[string]*SecuritySchemeRef
	links           map[string]*LinkRef
	callbacks       map[string]*CallbackRef
}

func (c *Components) Schemas() map[string]*SchemaRef                 { return c.schemas }
func (c *Components) Responses() map[string]*ResponseRef             { return c.responses }
func (c *Components) Parameters() map[string]*ParameterRef           { return c.parameters }
func (c *Components) Examples() map[string]*ExampleRef               { return c.examples }
func (c *Components) RequestBodies() map[string]*RequestBodyRef      { return c.requestBodies }
func (c *Components) Headers() map[string]*HeaderRef                 { return c.headers }
func (c *Components) SecuritySchemes() map[string]*SecuritySchemeRef { return c.securitySchemes }
func (c *Components) Links() map[string]*LinkRef                     { return c.links }
func (c *Components) Callbacks() map[string]*CallbackRef             { return c.callbacks }

// NewComponents creates a new Components instance.
func NewComponents(
	schemas map[string]*SchemaRef,
	responses map[string]*ResponseRef,
	parameters map[string]*ParameterRef,
	examples map[string]*ExampleRef,
	requestBodies map[string]*RequestBodyRef,
	headers map[string]*HeaderRef,
	securitySchemes map[string]*SecuritySchemeRef,
	links map[string]*LinkRef,
	callbacks map[string]*CallbackRef,
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
