package openapi30

// SecurityScheme defines a security scheme for the API.
// https://spec.openapis.org/oas/v3.0.3#security-scheme-object
type SecurityScheme struct {
	Node // embedded - provides VendorExtensions and Trix

	secType          string
	description      string
	name             string
	in               string
	scheme           string
	bearerFormat     string
	flows            *OAuthFlows
	openIDConnectURL string
}

func (ss *SecurityScheme) Type() string             { return ss.secType }
func (ss *SecurityScheme) Description() string      { return ss.description }
func (ss *SecurityScheme) Name() string             { return ss.name }
func (ss *SecurityScheme) In() string               { return ss.in }
func (ss *SecurityScheme) Scheme() string           { return ss.scheme }
func (ss *SecurityScheme) BearerFormat() string     { return ss.bearerFormat }
func (ss *SecurityScheme) Flows() *OAuthFlows       { return ss.flows }
func (ss *SecurityScheme) OpenIDConnectURL() string { return ss.openIDConnectURL }

// NewSecurityScheme creates a new SecurityScheme instance.
func NewSecurityScheme(secType, description, name, in, scheme, bearerFormat string, flows *OAuthFlows, openIDConnectURL string) *SecurityScheme {
	return &SecurityScheme{secType: secType, description: description, name: name, in: in, scheme: scheme, bearerFormat: bearerFormat, flows: flows, openIDConnectURL: openIDConnectURL}
}
