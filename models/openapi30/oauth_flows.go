package openapi30

// OAuthFlows allows configuration of supported OAuth flows.
// https://spec.openapis.org/oas/v3.0.3#oauth-flows-object
type OAuthFlows struct {
	Node // embedded - provides VendorExtensions and Trix

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
	return &OAuthFlows{implicit: implicit, password: password, clientCredentials: clientCredentials, authorizationCode: authorizationCode}
}
