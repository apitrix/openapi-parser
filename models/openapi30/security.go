package openapi30

// SecurityScheme defines a security scheme for the API.
// https://spec.openapis.org/oas/v3.0.3#security-scheme-object
type SecurityScheme struct {
	Node // embedded - provides NodeSource and Extensions

	Type             string      `json:"type" yaml:"type"`
	Description      string      `json:"description,omitempty" yaml:"description,omitempty"`
	Name             string      `json:"name,omitempty" yaml:"name,omitempty"`
	In               string      `json:"in,omitempty" yaml:"in,omitempty"`
	Scheme           string      `json:"scheme,omitempty" yaml:"scheme,omitempty"`
	BearerFormat     string      `json:"bearerFormat,omitempty" yaml:"bearerFormat,omitempty"`
	Flows            *OAuthFlows `json:"flows,omitempty" yaml:"flows,omitempty"`
	OpenIDConnectURL string      `json:"openIdConnectUrl,omitempty" yaml:"openIdConnectUrl,omitempty"`
}

// OAuthFlows allows configuration of supported OAuth flows.
// https://spec.openapis.org/oas/v3.0.3#oauth-flows-object
type OAuthFlows struct {
	Node // embedded - provides NodeSource and Extensions

	Implicit          *OAuthFlow `json:"implicit,omitempty" yaml:"implicit,omitempty"`
	Password          *OAuthFlow `json:"password,omitempty" yaml:"password,omitempty"`
	ClientCredentials *OAuthFlow `json:"clientCredentials,omitempty" yaml:"clientCredentials,omitempty"`
	AuthorizationCode *OAuthFlow `json:"authorizationCode,omitempty" yaml:"authorizationCode,omitempty"`
}

// OAuthFlow represents configuration for an OAuth flow.
// https://spec.openapis.org/oas/v3.0.3#oauth-flow-object
type OAuthFlow struct {
	Node // embedded - provides NodeSource and Extensions

	AuthorizationURL string            `json:"authorizationUrl,omitempty" yaml:"authorizationUrl,omitempty"`
	TokenURL         string            `json:"tokenUrl,omitempty" yaml:"tokenUrl,omitempty"`
	RefreshURL       string            `json:"refreshUrl,omitempty" yaml:"refreshUrl,omitempty"`
	Scopes           map[string]string `json:"scopes" yaml:"scopes"`
}

// SecurityRequirement lists required security schemes to execute an operation.
// https://spec.openapis.org/oas/v3.0.3#security-requirement-object
type SecurityRequirement map[string][]string
