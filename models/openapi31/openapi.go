package openapi31

// OpenAPI is the root document object of the OpenAPI specification.
// https://spec.openapis.org/oas/v3.1.0#openapi-object
type OpenAPI struct {
	Node // embedded - provides VendorExtensions and Trix

	openAPI           string
	info              *Info
	jsonSchemaDialect string
	servers           []*Server
	paths             *Paths
	webhooks          map[string]*PathItemRef
	components        *Components
	security          []SecurityRequirement
	tags              []*Tag
	externalDocs      *ExternalDocumentation
}

func (o *OpenAPI) OpenAPIVersion() string               { return o.openAPI }
func (o *OpenAPI) Info() *Info                          { return o.info }
func (o *OpenAPI) JsonSchemaDialect() string            { return o.jsonSchemaDialect }
func (o *OpenAPI) Servers() []*Server                   { return o.servers }
func (o *OpenAPI) Paths() *Paths                        { return o.paths }
func (o *OpenAPI) Webhooks() map[string]*PathItemRef    { return o.webhooks }
func (o *OpenAPI) Components() *Components              { return o.components }
func (o *OpenAPI) Security() []SecurityRequirement      { return o.security }
func (o *OpenAPI) Tags() []*Tag                         { return o.tags }
func (o *OpenAPI) ExternalDocs() *ExternalDocumentation { return o.externalDocs }

// SetProperty sets a named property on the OpenAPI document.
// Used by parsers for post-construction field assignment.
func (o *OpenAPI) SetProperty(name string, value interface{}) {
	switch name {
	case "openapi":
		o.openAPI = value.(string)
	case "info":
		o.info = value.(*Info)
	case "jsonSchemaDialect":
		o.jsonSchemaDialect = value.(string)
	case "servers":
		o.servers = value.([]*Server)
	case "paths":
		o.paths = value.(*Paths)
	case "webhooks":
		o.webhooks = value.(map[string]*PathItemRef)
	case "components":
		o.components = value.(*Components)
	case "security":
		o.security = value.([]SecurityRequirement)
	case "tags":
		o.tags = value.([]*Tag)
	case "externalDocs":
		o.externalDocs = value.(*ExternalDocumentation)
	}
}

// NewOpenAPI creates a new OpenAPI root document instance.
func NewOpenAPI(version string, info *Info) *OpenAPI {
	return &OpenAPI{openAPI: version, info: info}
}
