package openapi30

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// OAuthFlow represents configuration for an OAuth flow.
// https://spec.openapis.org/oas/v3.0.3#oauth-flow-object
type OAuthFlow struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	authorizationURL string
	tokenURL         string
	refreshURL       string
	scopes           map[string]string
}

func (f *OAuthFlow) AuthorizationURL() string  { return f.authorizationURL }
func (f *OAuthFlow) TokenURL() string          { return f.tokenURL }
func (f *OAuthFlow) RefreshURL() string        { return f.refreshURL }
func (f *OAuthFlow) Scopes() map[string]string { return f.scopes }

// NewOAuthFlow creates a new OAuthFlow instance.
func NewOAuthFlow(authorizationURL, tokenURL, refreshURL string, scopes map[string]string) *OAuthFlow {
	return &OAuthFlow{authorizationURL: authorizationURL, tokenURL: tokenURL, refreshURL: refreshURL, scopes: scopes}
}

func (f *OAuthFlow) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "authorizationUrl", Value: f.authorizationURL},
		{Key: "tokenUrl", Value: f.tokenURL},
		{Key: "refreshUrl", Value: f.refreshURL},
		{Key: "scopes", Value: f.scopes},
	}
	return shared.AppendExtensions(fields, f.VendorExtensions)
}

func (f *OAuthFlow) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(f.marshalFields())
}

func (f *OAuthFlow) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(f.marshalFields())
}

var _ yaml.Marshaler = (*OAuthFlow)(nil)
