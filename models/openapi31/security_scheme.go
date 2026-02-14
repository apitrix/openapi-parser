package openapi31

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// SecurityScheme defines a security scheme for the API.
// https://spec.openapis.org/oas/v3.1.0#security-scheme-object
type SecurityScheme struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	schemeType       string
	description      string
	name             string
	in               string
	scheme           string
	bearerFormat     string
	flows            *OAuthFlows
	openIDConnectURL string
}

func (s *SecurityScheme) Type() string             { return s.schemeType }
func (s *SecurityScheme) Description() string      { return s.description }
func (s *SecurityScheme) Name() string             { return s.name }
func (s *SecurityScheme) In() string               { return s.in }
func (s *SecurityScheme) Scheme() string           { return s.scheme }
func (s *SecurityScheme) BearerFormat() string     { return s.bearerFormat }
func (s *SecurityScheme) Flows() *OAuthFlows       { return s.flows }
func (s *SecurityScheme) OpenIDConnectURL() string { return s.openIDConnectURL }

func (s *SecurityScheme) SetType(schemeType string) error {
	if err := s.Trix.RunHooks("type", s.schemeType, schemeType); err != nil {
		return err
	}
	s.schemeType = schemeType
	return nil
}
func (s *SecurityScheme) SetDescription(description string) error {
	if err := s.Trix.RunHooks("description", s.description, description); err != nil {
		return err
	}
	s.description = description
	return nil
}
func (s *SecurityScheme) SetName(name string) error {
	if err := s.Trix.RunHooks("name", s.name, name); err != nil {
		return err
	}
	s.name = name
	return nil
}
func (s *SecurityScheme) SetIn(in string) error {
	if err := s.Trix.RunHooks("in", s.in, in); err != nil {
		return err
	}
	s.in = in
	return nil
}
func (s *SecurityScheme) SetScheme(scheme string) error {
	if err := s.Trix.RunHooks("scheme", s.scheme, scheme); err != nil {
		return err
	}
	s.scheme = scheme
	return nil
}
func (s *SecurityScheme) SetBearerFormat(bearerFormat string) error {
	if err := s.Trix.RunHooks("bearerFormat", s.bearerFormat, bearerFormat); err != nil {
		return err
	}
	s.bearerFormat = bearerFormat
	return nil
}
func (s *SecurityScheme) SetFlows(flows *OAuthFlows) error {
	if err := s.Trix.RunHooks("flows", s.flows, flows); err != nil {
		return err
	}
	s.flows = flows
	return nil
}
func (s *SecurityScheme) SetOpenIDConnectURL(openIDConnectURL string) error {
	if err := s.Trix.RunHooks("openIdConnectUrl", s.openIDConnectURL, openIDConnectURL); err != nil {
		return err
	}
	s.openIDConnectURL = openIDConnectURL
	return nil
}

// NewSecurityScheme creates a new SecurityScheme instance.
func NewSecurityScheme(schemeType, description, name, in, scheme, bearerFormat, openIDConnectURL string, flows *OAuthFlows) *SecurityScheme {
	return &SecurityScheme{
		schemeType: schemeType, description: description, name: name,
		in: in, scheme: scheme, bearerFormat: bearerFormat,
		flows: flows, openIDConnectURL: openIDConnectURL,
	}
}

func (s *SecurityScheme) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "type", Value: s.schemeType},
		{Key: "description", Value: s.description},
		{Key: "name", Value: s.name},
		{Key: "in", Value: s.in},
		{Key: "scheme", Value: s.scheme},
		{Key: "bearerFormat", Value: s.bearerFormat},
		{Key: "flows", Value: s.flows},
		{Key: "openIdConnectUrl", Value: s.openIDConnectURL},
	}
	return shared.AppendExtensions(fields, s.VendorExtensions)
}

func (s *SecurityScheme) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(s.marshalFields())
}

func (s *SecurityScheme) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(s.marshalFields())
}

var _ yaml.Marshaler = (*SecurityScheme)(nil)
