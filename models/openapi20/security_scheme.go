package openapi20

// SecurityScheme defines a security scheme for the API.
// https://swagger.io/specification/v2/#security-scheme-object
type SecurityScheme struct {
	Node // embedded - provides VendorExtensions and Trix

	Type             string            `json:"type" yaml:"type"`
	Description      string            `json:"description,omitempty" yaml:"description,omitempty"`
	Name             string            `json:"name,omitempty" yaml:"name,omitempty"`
	In               string            `json:"in,omitempty" yaml:"in,omitempty"`
	Flow             string            `json:"flow,omitempty" yaml:"flow,omitempty"`
	AuthorizationURL string            `json:"authorizationUrl,omitempty" yaml:"authorizationUrl,omitempty"`
	TokenURL         string            `json:"tokenUrl,omitempty" yaml:"tokenUrl,omitempty"`
	Scopes           map[string]string `json:"scopes,omitempty" yaml:"scopes,omitempty"`
}

// NewSecurityScheme creates a new SecurityScheme instance.
func NewSecurityScheme(securityType string) *SecurityScheme {
	return &SecurityScheme{Type: securityType}
}
