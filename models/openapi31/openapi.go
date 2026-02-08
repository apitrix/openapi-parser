package openapi31

// OpenAPI is the root document object of the OpenAPI specification.
// https://spec.openapis.org/oas/v3.1.0#openapi-object
type OpenAPI struct {
	Node // embedded - provides VendorExtensions and Trix

	OpenAPI           string                  `json:"openapi" yaml:"openapi"`
	Info              *Info                   `json:"info" yaml:"info"`
	JsonSchemaDialect string                  `json:"jsonSchemaDialect,omitempty" yaml:"jsonSchemaDialect,omitempty"`
	Servers           []*Server               `json:"servers,omitempty" yaml:"servers,omitempty"`
	Paths             *Paths                  `json:"paths,omitempty" yaml:"paths,omitempty"`
	Webhooks          map[string]*PathItemRef `json:"webhooks,omitempty" yaml:"webhooks,omitempty"`
	Components        *Components             `json:"components,omitempty" yaml:"components,omitempty"`
	Security          []SecurityRequirement   `json:"security,omitempty" yaml:"security,omitempty"`
	Tags              []*Tag                  `json:"tags,omitempty" yaml:"tags,omitempty"`
	ExternalDocs      *ExternalDocumentation  `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
}
