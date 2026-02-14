package openapi30

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// Components holds reusable objects for the specification.
// https://spec.openapis.org/oas/v3.0.3#components-object
type Components struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	schemas         map[string]*RefSchema
	responses       map[string]*RefResponse
	parameters      map[string]*RefParameter
	examples        map[string]*RefExample
	requestBodies   map[string]*RefRequestBody
	headers         map[string]*RefHeader
	securitySchemes map[string]*RefSecurityScheme
	links           map[string]*RefLink
	callbacks       map[string]*RefCallback
}

func (c *Components) Schemas() map[string]*RefSchema            { return c.schemas }
func (c *Components) Responses() map[string]*RefResponse        { return c.responses }
func (c *Components) Parameters() map[string]*RefParameter      { return c.parameters }
func (c *Components) Examples() map[string]*RefExample          { return c.examples }
func (c *Components) RequestBodies() map[string]*RefRequestBody { return c.requestBodies }
func (c *Components) Headers() map[string]*RefHeader            { return c.headers }
func (c *Components) SecuritySchemes() map[string]*RefSecurityScheme {
	return c.securitySchemes
}
func (c *Components) Links() map[string]*RefLink         { return c.links }
func (c *Components) Callbacks() map[string]*RefCallback { return c.callbacks }

func (c *Components) SetSchemas(schemas map[string]*RefSchema) error {
	if err := c.Trix.RunHooks("schemas", c.schemas, schemas); err != nil {
		return err
	}
	c.schemas = schemas
	return nil
}
func (c *Components) SetResponses(responses map[string]*RefResponse) error {
	if err := c.Trix.RunHooks("responses", c.responses, responses); err != nil {
		return err
	}
	c.responses = responses
	return nil
}
func (c *Components) SetParameters(parameters map[string]*RefParameter) error {
	if err := c.Trix.RunHooks("parameters", c.parameters, parameters); err != nil {
		return err
	}
	c.parameters = parameters
	return nil
}
func (c *Components) SetExamples(examples map[string]*RefExample) error {
	if err := c.Trix.RunHooks("examples", c.examples, examples); err != nil {
		return err
	}
	c.examples = examples
	return nil
}
func (c *Components) SetRequestBodies(requestBodies map[string]*RefRequestBody) error {
	if err := c.Trix.RunHooks("requestBodies", c.requestBodies, requestBodies); err != nil {
		return err
	}
	c.requestBodies = requestBodies
	return nil
}
func (c *Components) SetHeaders(headers map[string]*RefHeader) error {
	if err := c.Trix.RunHooks("headers", c.headers, headers); err != nil {
		return err
	}
	c.headers = headers
	return nil
}
func (c *Components) SetSecuritySchemes(securitySchemes map[string]*RefSecurityScheme) error {
	if err := c.Trix.RunHooks("securitySchemes", c.securitySchemes, securitySchemes); err != nil {
		return err
	}
	c.securitySchemes = securitySchemes
	return nil
}
func (c *Components) SetLinks(links map[string]*RefLink) error {
	if err := c.Trix.RunHooks("links", c.links, links); err != nil {
		return err
	}
	c.links = links
	return nil
}
func (c *Components) SetCallbacks(callbacks map[string]*RefCallback) error {
	if err := c.Trix.RunHooks("callbacks", c.callbacks, callbacks); err != nil {
		return err
	}
	c.callbacks = callbacks
	return nil
}

// NewComponents creates a new Components instance.
func NewComponents(
	schemas map[string]*RefSchema,
	responses map[string]*RefResponse,
	parameters map[string]*RefParameter,
	examples map[string]*RefExample,
	requestBodies map[string]*RefRequestBody,
	headers map[string]*RefHeader,
	securitySchemes map[string]*RefSecurityScheme,
	links map[string]*RefLink,
	callbacks map[string]*RefCallback,
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
