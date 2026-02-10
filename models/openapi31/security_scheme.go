package openapi31

// SecurityScheme defines a security scheme for the API.
// https://spec.openapis.org/oas/v3.1.0#security-scheme-object
type SecurityScheme struct {
	Node // embedded - provides VendorExtensions and Trix

	Type             string      `json:"type" yaml:"type"`
	Description      string      `json:"description,omitempty" yaml:"description,omitempty"`
	Name             string      `json:"name,omitempty" yaml:"name,omitempty"`
	In               string      `json:"in,omitempty" yaml:"in,omitempty"`
	Scheme           string      `json:"scheme,omitempty" yaml:"scheme,omitempty"`
	BearerFormat     string      `json:"bearerFormat,omitempty" yaml:"bearerFormat,omitempty"`
	Flows            *OAuthFlows `json:"flows,omitempty" yaml:"flows,omitempty"`
	OpenIDConnectURL string      `json:"openIdConnectUrl,omitempty" yaml:"openIdConnectUrl,omitempty"`
}

// NewSecurityScheme creates a new SecurityScheme instance.
func NewSecurityScheme(securityType string) *SecurityScheme {
	return &SecurityScheme{Type: securityType}
}
