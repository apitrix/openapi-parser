package openapi30

// Components holds reusable objects for the specification.
// https://spec.openapis.org/oas/v3.0.3#components-object
type Components struct {
	Node // embedded - provides VendorExtensions and Trix

	Schemas         map[string]*SchemaRef         `json:"schemas,omitempty" yaml:"schemas,omitempty"`
	Responses       map[string]*ResponseRef       `json:"responses,omitempty" yaml:"responses,omitempty"`
	Parameters      map[string]*ParameterRef      `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Examples        map[string]*ExampleRef        `json:"examples,omitempty" yaml:"examples,omitempty"`
	RequestBodies   map[string]*RequestBodyRef    `json:"requestBodies,omitempty" yaml:"requestBodies,omitempty"`
	Headers         map[string]*HeaderRef         `json:"headers,omitempty" yaml:"headers,omitempty"`
	SecuritySchemes map[string]*SecuritySchemeRef `json:"securitySchemes,omitempty" yaml:"securitySchemes,omitempty"`
	Links           map[string]*LinkRef           `json:"links,omitempty" yaml:"links,omitempty"`
	Callbacks       map[string]*CallbackRef       `json:"callbacks,omitempty" yaml:"callbacks,omitempty"`
}
