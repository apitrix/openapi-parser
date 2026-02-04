package openapi30

// OpenAPI is the root document object of the OpenAPI specification.
// https://spec.openapis.org/oas/v3.0.3#openapi-object
type OpenAPI struct {
	Node // embedded - provides NodeSource and Extensions

	OpenAPI      string                `json:"openapi" yaml:"openapi"`
	Info         *Info                 `json:"info" yaml:"info"`
	Servers      []*Server             `json:"servers,omitempty" yaml:"servers,omitempty"`
	Paths        *Paths                `json:"paths,omitempty" yaml:"paths,omitempty"`
	Components   *Components           `json:"components,omitempty" yaml:"components,omitempty"`
	Security     []SecurityRequirement `json:"security,omitempty" yaml:"security,omitempty"`
	Tags         []*Tag                `json:"tags,omitempty" yaml:"tags,omitempty"`
	ExternalDocs *ExternalDocs         `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
}
