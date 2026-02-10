package openapi30

// OAuthFlow represents configuration for an OAuth flow.
// https://spec.openapis.org/oas/v3.0.3#oauth-flow-object
type OAuthFlow struct {
	Node // embedded - provides VendorExtensions and Trix

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
