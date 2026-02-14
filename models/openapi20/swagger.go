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
	definitions         map[string]*RefSchema
	parameters          map[string]*RefParameter
	responses           map[string]*RefResponse
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
func (s *Swagger) Definitions() map[string]*RefSchema              { return s.definitions }
func (s *Swagger) Parameters() map[string]*RefParameter            { return s.parameters }
func (s *Swagger) Responses() map[string]*RefResponse              { return s.responses }
func (s *Swagger) SecurityDefinitions() map[string]*SecurityScheme { return s.securityDefinitions }
func (s *Swagger) Security() []SecurityRequirement                 { return s.security }
func (s *Swagger) Tags() []*Tag                                    { return s.tags }
func (s *Swagger) ExternalDocs() *ExternalDocs                     { return s.externalDocs }

func (s *Swagger) SetSwaggerVersion(swagger string) error {
	if err := s.Trix.RunHooks("swagger", s.swagger, swagger); err != nil {
		return err
	}
	s.swagger = swagger
	return nil
}
func (s *Swagger) SetInfo(info *Info) error {
	if err := s.Trix.RunHooks("info", s.info, info); err != nil {
		return err
	}
	s.info = info
	return nil
}
func (s *Swagger) SetHost(host string) error {
	if err := s.Trix.RunHooks("host", s.host, host); err != nil {
		return err
	}
	s.host = host
	return nil
}
func (s *Swagger) SetBasePath(basePath string) error {
	if err := s.Trix.RunHooks("basePath", s.basePath, basePath); err != nil {
		return err
	}
	s.basePath = basePath
	return nil
}
func (s *Swagger) SetSchemes(schemes []string) error {
	if err := s.Trix.RunHooks("schemes", s.schemes, schemes); err != nil {
		return err
	}
	s.schemes = schemes
	return nil
}
func (s *Swagger) SetConsumes(consumes []string) error {
	if err := s.Trix.RunHooks("consumes", s.consumes, consumes); err != nil {
		return err
	}
	s.consumes = consumes
	return nil
}
func (s *Swagger) SetProduces(produces []string) error {
	if err := s.Trix.RunHooks("produces", s.produces, produces); err != nil {
		return err
	}
	s.produces = produces
	return nil
}
func (s *Swagger) SetPaths(paths *Paths) error {
	if err := s.Trix.RunHooks("paths", s.paths, paths); err != nil {
		return err
	}
	s.paths = paths
	return nil
}
func (s *Swagger) SetDefinitions(definitions map[string]*RefSchema) error {
	if err := s.Trix.RunHooks("definitions", s.definitions, definitions); err != nil {
		return err
	}
	s.definitions = definitions
	return nil
}
func (s *Swagger) SetParameters(parameters map[string]*RefParameter) error {
	if err := s.Trix.RunHooks("parameters", s.parameters, parameters); err != nil {
		return err
	}
	s.parameters = parameters
	return nil
}
func (s *Swagger) SetResponses(responses map[string]*RefResponse) error {
	if err := s.Trix.RunHooks("responses", s.responses, responses); err != nil {
		return err
	}
	s.responses = responses
	return nil
}
func (s *Swagger) SetSecurityDefinitions(securityDefinitions map[string]*SecurityScheme) error {
	if err := s.Trix.RunHooks("securityDefinitions", s.securityDefinitions, securityDefinitions); err != nil {
		return err
	}
	s.securityDefinitions = securityDefinitions
	return nil
}
func (s *Swagger) SetSecurity(security []SecurityRequirement) error {
	if err := s.Trix.RunHooks("security", s.security, security); err != nil {
		return err
	}
	s.security = security
	return nil
}
func (s *Swagger) SetTags(tags []*Tag) error {
	if err := s.Trix.RunHooks("tags", s.tags, tags); err != nil {
		return err
	}
	s.tags = tags
	return nil
}
func (s *Swagger) SetExternalDocs(externalDocs *ExternalDocs) error {
	if err := s.Trix.RunHooks("externalDocs", s.externalDocs, externalDocs); err != nil {
		return err
	}
	s.externalDocs = externalDocs
	return nil
}

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
		s.definitions = value.(map[string]*RefSchema)
	case "parameters":
		s.parameters = value.(map[string]*RefParameter)
	case "responses":
		s.responses = value.(map[string]*RefResponse)
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
