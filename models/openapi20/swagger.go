package openapi20

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// Swagger is the root document object of the Swagger 2.0 specification.
// https://swagger.io/specification/v2/#swagger-object
type Swagger struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	swagger             string
	info                *Info
	host                string
	basePath            string
	schemes             []string
	consumes            []string
	produces            []string
	paths               *Paths
	definitions         map[string]*shared.Ref[Schema]
	parameters          map[string]*shared.Ref[Parameter]
	responses           map[string]*shared.Ref[Response]
	securityDefinitions map[string]*SecurityScheme
	security            []SecurityRequirement
	tags                []*Tag
	externalDocs        *ExternalDocs
}

func (s *Swagger) SwaggerVersion() string                          { return s.swagger }
func (s *Swagger) Info() *Info                                     { return s.info }
func (s *Swagger) Host() string                                    { return s.host }
func (s *Swagger) BasePath() string                                { return s.basePath }
func (s *Swagger) Schemes() []string                               { return s.schemes }
func (s *Swagger) Consumes() []string                              { return s.consumes }
func (s *Swagger) Produces() []string                              { return s.produces }
func (s *Swagger) Paths() *Paths                                   { return s.paths }
func (s *Swagger) Definitions() map[string]*shared.Ref[Schema]     { return s.definitions }
func (s *Swagger) Parameters() map[string]*shared.Ref[Parameter]   { return s.parameters }
func (s *Swagger) Responses() map[string]*shared.Ref[Response]     { return s.responses }
func (s *Swagger) SecurityDefinitions() map[string]*SecurityScheme { return s.securityDefinitions }
func (s *Swagger) Security() []SecurityRequirement                 { return s.security }
func (s *Swagger) Tags() []*Tag                                    { return s.tags }
func (s *Swagger) ExternalDocs() *ExternalDocs                     { return s.externalDocs }

// SetProperty sets a property on the Swagger document. Used by parsers for
// incremental building of the root document where many fields are optional.
func (s *Swagger) SetProperty(key string, value interface{}) {
	switch key {
	case "swagger":
		s.swagger = value.(string)
	case "info":
		s.info = value.(*Info)
	case "host":
		s.host = value.(string)
	case "basePath":
		s.basePath = value.(string)
	case "schemes":
		s.schemes = value.([]string)
	case "consumes":
		s.consumes = value.([]string)
	case "produces":
		s.produces = value.([]string)
	case "paths":
		s.paths = value.(*Paths)
	case "definitions":
		s.definitions = value.(map[string]*shared.Ref[Schema])
	case "parameters":
		s.parameters = value.(map[string]*shared.Ref[Parameter])
	case "responses":
		s.responses = value.(map[string]*shared.Ref[Response])
	case "securityDefinitions":
		s.securityDefinitions = value.(map[string]*SecurityScheme)
	case "security":
		s.security = value.([]SecurityRequirement)
	case "tags":
		s.tags = value.([]*Tag)
	case "externalDocs":
		s.externalDocs = value.(*ExternalDocs)
	}
}

func (s *Swagger) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "swagger", Value: s.swagger},
		{Key: "info", Value: s.info},
		{Key: "host", Value: s.host},
		{Key: "basePath", Value: s.basePath},
		{Key: "schemes", Value: s.schemes},
		{Key: "consumes", Value: s.consumes},
		{Key: "produces", Value: s.produces},
		{Key: "paths", Value: s.paths},
		{Key: "definitions", Value: s.definitions},
		{Key: "parameters", Value: s.parameters},
		{Key: "responses", Value: s.responses},
		{Key: "securityDefinitions", Value: s.securityDefinitions},
		{Key: "security", Value: s.security},
		{Key: "tags", Value: s.tags},
		{Key: "externalDocs", Value: s.externalDocs},
	}
	return shared.AppendExtensions(fields, s.VendorExtensions)
}

func (s *Swagger) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(s.marshalFields())
}

func (s *Swagger) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(s.marshalFields())
}

var _ yaml.Marshaler = (*Swagger)(nil)
