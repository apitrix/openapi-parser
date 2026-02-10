package openapi31

// Components holds reusable objects for the specification.
// https://spec.openapis.org/oas/v3.1.0#components-object
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
	pathItems       map[string]*PathItemRef
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
func (c *Components) PathItems() map[string]*PathItemRef             { return c.pathItems }

// SetProperty sets a named property on the Components.
// Used by parsers for post-construction field assignment.
func (c *Components) SetProperty(name string, value interface{}) {
	switch name {
	case "schemas":
		c.schemas = value.(map[string]*SchemaRef)
	case "responses":
		c.responses = value.(map[string]*ResponseRef)
	case "parameters":
		c.parameters = value.(map[string]*ParameterRef)
	case "examples":
		c.examples = value.(map[string]*ExampleRef)
	case "requestBodies":
		c.requestBodies = value.(map[string]*RequestBodyRef)
	case "headers":
		c.headers = value.(map[string]*HeaderRef)
	case "securitySchemes":
		c.securitySchemes = value.(map[string]*SecuritySchemeRef)
	case "links":
		c.links = value.(map[string]*LinkRef)
	case "callbacks":
		c.callbacks = value.(map[string]*CallbackRef)
	case "pathItems":
		c.pathItems = value.(map[string]*PathItemRef)
	}
}

// NewComponents creates a new Components instance.
func NewComponents() *Components {
	return &Components{}
}
