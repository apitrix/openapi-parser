package openapi31

// OAuthFlow represents configuration for an OAuth flow.
// https://spec.openapis.org/oas/v3.1.0#oauth-flow-object
type OAuthFlow struct {
	Node // embedded - provides VendorExtensions and Trix

	AuthorizationURL string            `json:"authorizationUrl,omitempty" yaml:"authorizationUrl,omitempty"`
	TokenURL         string            `json:"tokenUrl,omitempty" yaml:"tokenUrl,omitempty"`
	RefreshURL       string            `json:"refreshUrl,omitempty" yaml:"refreshUrl,omitempty"`
	Scopes           map[string]string `json:"scopes" yaml:"scopes"`
}

// NewOAuthFlow creates a new OAuthFlow instance.
func NewOAuthFlow() *OAuthFlow {
	return &OAuthFlow{}
}
