package openapi20

import (
	"github.com/apitrix/openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// SecurityScheme defines a security scheme for the API.
// https://swagger.io/specification/v2/#security-scheme-object
type SecurityScheme struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	securityType     string
	description      string
	name             string
	in               string
	flow             string
	authorizationURL string
	tokenURL         string
	scopes           map[string]string
}

func (s *SecurityScheme) Type() string              { return s.securityType }
func (s *SecurityScheme) Description() string       { return s.description }
func (s *SecurityScheme) Name() string              { return s.name }
func (s *SecurityScheme) In() string                { return s.in }
func (s *SecurityScheme) Flow() string              { return s.flow }
func (s *SecurityScheme) AuthorizationURL() string  { return s.authorizationURL }
func (s *SecurityScheme) TokenURL() string          { return s.tokenURL }
func (s *SecurityScheme) Scopes() map[string]string { return s.scopes }

func (s *SecurityScheme) SetType(securityType string) error {
	if err := s.Trix.RunHooks("type", s.securityType, securityType); err != nil {
		return err
	}
	s.securityType = securityType
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
func (s *SecurityScheme) SetFlow(flow string) error {
	if err := s.Trix.RunHooks("flow", s.flow, flow); err != nil {
		return err
	}
	s.flow = flow
	return nil
}
func (s *SecurityScheme) SetAuthorizationURL(authorizationURL string) error {
	if err := s.Trix.RunHooks("authorizationUrl", s.authorizationURL, authorizationURL); err != nil {
		return err
	}
	s.authorizationURL = authorizationURL
	return nil
}
func (s *SecurityScheme) SetTokenURL(tokenURL string) error {
	if err := s.Trix.RunHooks("tokenUrl", s.tokenURL, tokenURL); err != nil {
		return err
	}
	s.tokenURL = tokenURL
	return nil
}
func (s *SecurityScheme) SetScopes(scopes map[string]string) error {
	if err := s.Trix.RunHooks("scopes", s.scopes, scopes); err != nil {
		return err
	}
	s.scopes = scopes
	return nil
}

// NewSecurityScheme creates a new SecurityScheme instance.
func NewSecurityScheme(
	securityType, description, name, in, flow, authorizationURL, tokenURL string,
	scopes map[string]string,
) *SecurityScheme {
	return &SecurityScheme{
		securityType: securityType, description: description,
		name: name, in: in, flow: flow,
		authorizationURL: authorizationURL, tokenURL: tokenURL,
		scopes: scopes,
	}
}

func (s *SecurityScheme) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "type", Value: s.securityType},
		{Key: "description", Value: s.description},
		{Key: "name", Value: s.name},
		{Key: "in", Value: s.in},
		{Key: "flow", Value: s.flow},
		{Key: "authorizationUrl", Value: s.authorizationURL},
		{Key: "tokenUrl", Value: s.tokenURL},
		{Key: "scopes", Value: s.scopes},
	}
	return shared.AppendExtensions(fields, s.VendorExtensions)
}

// MarshalFields implements shared.MarshalFieldsProvider for export.
func (s *SecurityScheme) MarshalFields() []shared.Field { return s.marshalFields() }

func (s *SecurityScheme) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(s.marshalFields())
}

func (s *SecurityScheme) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(s.marshalFields())
}

var _ yaml.Marshaler = (*SecurityScheme)(nil)
