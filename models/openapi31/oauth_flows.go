package openapi31

// OAuthFlows allows configuration of supported OAuth flows.
// https://spec.openapis.org/oas/v3.1.0#oauth-flows-object
type OAuthFlows struct {
	Node // embedded - provides VendorExtensions and Trix

	Implicit          *OAuthFlow `json:"implicit,omitempty" yaml:"implicit,omitempty"`
	Password          *OAuthFlow `json:"password,omitempty" yaml:"password,omitempty"`
	ClientCredentials *OAuthFlow `json:"clientCredentials,omitempty" yaml:"clientCredentials,omitempty"`
	AuthorizationCode *OAuthFlow `json:"authorizationCode,omitempty" yaml:"authorizationCode,omitempty"`
}

// NewOAuthFlows creates a new OAuthFlows instance.
func NewOAuthFlows() *OAuthFlows {
	return &OAuthFlows{}
}
