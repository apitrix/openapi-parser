package openapi31

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// Components holds reusable objects for the specification.
// https://spec.openapis.org/oas/v3.1.0#components-object
type Components struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	schemas         map[string]*shared.RefWithMeta[Schema]
	responses       map[string]*shared.RefWithMeta[Response]
	parameters      map[string]*shared.RefWithMeta[Parameter]
	examples        map[string]*shared.RefWithMeta[Example]
	requestBodies   map[string]*shared.RefWithMeta[RequestBody]
	headers         map[string]*shared.RefWithMeta[Header]
	securitySchemes map[string]*shared.RefWithMeta[SecurityScheme]
	links           map[string]*shared.RefWithMeta[Link]
	callbacks       map[string]*shared.RefWithMeta[Callback]
	pathItems       map[string]*shared.RefWithMeta[PathItem]
}

func (c *Components) Schemas() map[string]*shared.RefWithMeta[Schema]       { return c.schemas }
func (c *Components) Responses() map[string]*shared.RefWithMeta[Response]   { return c.responses }
func (c *Components) Parameters() map[string]*shared.RefWithMeta[Parameter] { return c.parameters }
func (c *Components) Examples() map[string]*shared.RefWithMeta[Example]     { return c.examples }
func (c *Components) RequestBodies() map[string]*shared.RefWithMeta[RequestBody] {
	return c.requestBodies
}
func (c *Components) Headers() map[string]*shared.RefWithMeta[Header] { return c.headers }
func (c *Components) SecuritySchemes() map[string]*shared.RefWithMeta[SecurityScheme] {
	return c.securitySchemes
}
func (c *Components) Links() map[string]*shared.RefWithMeta[Link]         { return c.links }
func (c *Components) Callbacks() map[string]*shared.RefWithMeta[Callback] { return c.callbacks }
func (c *Components) PathItems() map[string]*shared.RefWithMeta[PathItem] { return c.pathItems }

func (c *Components) SetSchemas(schemas map[string]*shared.RefWithMeta[Schema]) error {
	if err := c.Trix.RunHooks("schemas", c.schemas, schemas); err != nil {
		return err
	}
	c.schemas = schemas
	return nil
}
func (c *Components) SetResponses(responses map[string]*shared.RefWithMeta[Response]) error {
	if err := c.Trix.RunHooks("responses", c.responses, responses); err != nil {
		return err
	}
	c.responses = responses
	return nil
}
func (c *Components) SetParameters(parameters map[string]*shared.RefWithMeta[Parameter]) error {
	if err := c.Trix.RunHooks("parameters", c.parameters, parameters); err != nil {
		return err
	}
	c.parameters = parameters
	return nil
}
func (c *Components) SetExamples(examples map[string]*shared.RefWithMeta[Example]) error {
	if err := c.Trix.RunHooks("examples", c.examples, examples); err != nil {
		return err
	}
	c.examples = examples
	return nil
}
func (c *Components) SetRequestBodies(requestBodies map[string]*shared.RefWithMeta[RequestBody]) error {
	if err := c.Trix.RunHooks("requestBodies", c.requestBodies, requestBodies); err != nil {
		return err
	}
	c.requestBodies = requestBodies
	return nil
}
func (c *Components) SetHeaders(headers map[string]*shared.RefWithMeta[Header]) error {
	if err := c.Trix.RunHooks("headers", c.headers, headers); err != nil {
		return err
	}
	c.headers = headers
	return nil
}
func (c *Components) SetSecuritySchemes(securitySchemes map[string]*shared.RefWithMeta[SecurityScheme]) error {
	if err := c.Trix.RunHooks("securitySchemes", c.securitySchemes, securitySchemes); err != nil {
		return err
	}
	c.securitySchemes = securitySchemes
	return nil
}
func (c *Components) SetLinks(links map[string]*shared.RefWithMeta[Link]) error {
	if err := c.Trix.RunHooks("links", c.links, links); err != nil {
		return err
	}
	c.links = links
	return nil
}
func (c *Components) SetCallbacks(callbacks map[string]*shared.RefWithMeta[Callback]) error {
	if err := c.Trix.RunHooks("callbacks", c.callbacks, callbacks); err != nil {
		return err
	}
	c.callbacks = callbacks
	return nil
}
func (c *Components) SetPathItems(pathItems map[string]*shared.RefWithMeta[PathItem]) error {
	if err := c.Trix.RunHooks("pathItems", c.pathItems, pathItems); err != nil {
		return err
	}
	c.pathItems = pathItems
	return nil
}

// SetProperty sets a named property on the Components.
// Used by parsers for post-construction field assignment.
func (c *Components) SetProperty(name string, value interface{}) {
	switch name {
	case "schemas":
		c.schemas = value.(map[string]*shared.RefWithMeta[Schema])
	case "responses":
		c.responses = value.(map[string]*shared.RefWithMeta[Response])
	case "parameters":
		c.parameters = value.(map[string]*shared.RefWithMeta[Parameter])
	case "examples":
		c.examples = value.(map[string]*shared.RefWithMeta[Example])
	case "requestBodies":
		c.requestBodies = value.(map[string]*shared.RefWithMeta[RequestBody])
	case "headers":
		c.headers = value.(map[string]*shared.RefWithMeta[Header])
	case "securitySchemes":
		c.securitySchemes = value.(map[string]*shared.RefWithMeta[SecurityScheme])
	case "links":
		c.links = value.(map[string]*shared.RefWithMeta[Link])
	case "callbacks":
		c.callbacks = value.(map[string]*shared.RefWithMeta[Callback])
	case "pathItems":
		c.pathItems = value.(map[string]*shared.RefWithMeta[PathItem])
	}
}

// NewComponents creates a new Components instance.
func NewComponents() *Components {
	return &Components{}
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
		{Key: "pathItems", Value: c.pathItems},
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
