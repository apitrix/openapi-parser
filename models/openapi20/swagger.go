package openapi20

// Swagger is the root document object of the Swagger 2.0 specification.
// https://swagger.io/specification/v2/#swagger-object
type Swagger struct {
	Node // embedded - provides NodeSource and Extensions

	Swagger             string                         `json:"swagger" yaml:"swagger"`
	Info                *Info                          `json:"info" yaml:"info"`
	Host                string                         `json:"host,omitempty" yaml:"host,omitempty"`
	BasePath            string                         `json:"basePath,omitempty" yaml:"basePath,omitempty"`
	Schemes             []string                       `json:"schemes,omitempty" yaml:"schemes,omitempty"`
	Consumes            []string                      `json:"consumes,omitempty" yaml:"consumes,omitempty"`
	Produces            []string                      `json:"produces,omitempty" yaml:"produces,omitempty"`
	Paths               *Paths                         `json:"paths" yaml:"paths"`
	Definitions         map[string]*SchemaRef          `json:"definitions,omitempty" yaml:"definitions,omitempty"`
	Parameters          map[string]*ParameterRef       `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Responses           map[string]*ResponseRef        `json:"responses,omitempty" yaml:"responses,omitempty"`
	SecurityDefinitions map[string]*SecurityScheme      `json:"securityDefinitions,omitempty" yaml:"securityDefinitions,omitempty"`
	Security            []SecurityRequirement         `json:"security,omitempty" yaml:"security,omitempty"`
	Tags                []*Tag                         `json:"tags,omitempty" yaml:"tags,omitempty"`
	ExternalDocs        *ExternalDocs                  `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
}
