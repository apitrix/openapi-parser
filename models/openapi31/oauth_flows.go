package openapi31

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// OAuthFlows allows configuration of supported OAuth flows.
// https://spec.openapis.org/oas/v3.1.0#oauth-flows-object
type OAuthFlows struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	implicit          *OAuthFlow
	password          *OAuthFlow
	clientCredentials *OAuthFlow
	authorizationCode *OAuthFlow
}

func (f *OAuthFlows) Implicit() *OAuthFlow          { return f.implicit }
func (f *OAuthFlows) Password() *OAuthFlow          { return f.password }
func (f *OAuthFlows) ClientCredentials() *OAuthFlow { return f.clientCredentials }
func (f *OAuthFlows) AuthorizationCode() *OAuthFlow { return f.authorizationCode }

// NewOAuthFlows creates a new OAuthFlows instance.
func NewOAuthFlows(implicit, password, clientCredentials, authorizationCode *OAuthFlow) *OAuthFlows {
	return &OAuthFlows{
		implicit: implicit, password: password,
		clientCredentials: clientCredentials, authorizationCode: authorizationCode,
	}
}

func (f *OAuthFlows) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "implicit", Value: f.implicit},
		{Key: "password", Value: f.password},
		{Key: "clientCredentials", Value: f.clientCredentials},
		{Key: "authorizationCode", Value: f.authorizationCode},
	}
	return shared.AppendExtensions(fields, f.VendorExtensions)
}

func (f *OAuthFlows) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(f.marshalFields())
}

func (f *OAuthFlows) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(f.marshalFields())
}

var _ yaml.Marshaler = (*OAuthFlows)(nil)
