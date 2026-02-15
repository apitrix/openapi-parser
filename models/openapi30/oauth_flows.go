package openapi30

import (
	"github.com/apitrix/openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// OAuthFlows allows configuration of supported OAuth flows.
// https://spec.openapis.org/oas/v3.0.3#oauth-flows-object
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

func (f *OAuthFlows) SetImplicit(implicit *OAuthFlow) error {
	if err := f.Trix.RunHooks("implicit", f.implicit, implicit); err != nil {
		return err
	}
	f.implicit = implicit
	return nil
}
func (f *OAuthFlows) SetPassword(password *OAuthFlow) error {
	if err := f.Trix.RunHooks("password", f.password, password); err != nil {
		return err
	}
	f.password = password
	return nil
}
func (f *OAuthFlows) SetClientCredentials(clientCredentials *OAuthFlow) error {
	if err := f.Trix.RunHooks("clientCredentials", f.clientCredentials, clientCredentials); err != nil {
		return err
	}
	f.clientCredentials = clientCredentials
	return nil
}
func (f *OAuthFlows) SetAuthorizationCode(authorizationCode *OAuthFlow) error {
	if err := f.Trix.RunHooks("authorizationCode", f.authorizationCode, authorizationCode); err != nil {
		return err
	}
	f.authorizationCode = authorizationCode
	return nil
}

// NewOAuthFlows creates a new OAuthFlows instance.
func NewOAuthFlows(implicit, password, clientCredentials, authorizationCode *OAuthFlow) *OAuthFlows {
	return &OAuthFlows{implicit: implicit, password: password, clientCredentials: clientCredentials, authorizationCode: authorizationCode}
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

// MarshalFields implements shared.MarshalFieldsProvider for export.
func (f *OAuthFlows) MarshalFields() []shared.Field { return f.marshalFields() }

func (f *OAuthFlows) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(f.marshalFields())
}

func (f *OAuthFlows) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(f.marshalFields())
}

var _ yaml.Marshaler = (*OAuthFlows)(nil)
