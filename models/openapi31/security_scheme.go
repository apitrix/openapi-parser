package openapi31

// SecurityScheme defines a security scheme for the API.
// https://spec.openapis.org/oas/v3.1.0#security-scheme-object
type SecurityScheme struct {
	Node // embedded - provides VendorExtensions and Trix

	schemeType       string
	description      string
	name             string
	in               string
	scheme           string
	bearerFormat     string
	flows            *OAuthFlows
	openIDConnectURL string
}

func (s *SecurityScheme) Type() string             { return s.schemeType }
func (s *SecurityScheme) Description() string      { return s.description }
func (s *SecurityScheme) Name() string             { return s.name }
func (s *SecurityScheme) In() string               { return s.in }
func (s *SecurityScheme) Scheme() string           { return s.scheme }
func (s *SecurityScheme) BearerFormat() string     { return s.bearerFormat }
func (s *SecurityScheme) Flows() *OAuthFlows       { return s.flows }
func (s *SecurityScheme) OpenIDConnectURL() string { return s.openIDConnectURL }

// NewSecurityScheme creates a new SecurityScheme instance.
func NewSecurityScheme(schemeType, description, name, in, scheme, bearerFormat, openIDConnectURL string, flows *OAuthFlows) *SecurityScheme {
	return &SecurityScheme{
		schemeType: schemeType, description: description, name: name,
		in: in, scheme: scheme, bearerFormat: bearerFormat,
		flows: flows, openIDConnectURL: openIDConnectURL,
	}
}
