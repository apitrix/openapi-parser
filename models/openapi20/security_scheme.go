package openapi20

// SecurityScheme defines a security scheme for the API.
// https://swagger.io/specification/v2/#security-scheme-object
type SecurityScheme struct {
	Node // embedded - provides VendorExtensions and Trix

	securityType     string
	description      string
	name             string
	in               string
	flow             string
	authorizationURL string
	tokenURL         string
	scopes           map[string]string
}

func (s *SecurityScheme) Type() string              { return s.securityType }
func (s *SecurityScheme) Description() string       { return s.description }
func (s *SecurityScheme) Name() string              { return s.name }
func (s *SecurityScheme) In() string                { return s.in }
func (s *SecurityScheme) Flow() string              { return s.flow }
func (s *SecurityScheme) AuthorizationURL() string  { return s.authorizationURL }
func (s *SecurityScheme) TokenURL() string          { return s.tokenURL }
func (s *SecurityScheme) Scopes() map[string]string { return s.scopes }

// NewSecurityScheme creates a new SecurityScheme instance.
func NewSecurityScheme(
	securityType, description, name, in, flow, authorizationURL, tokenURL string,
	scopes map[string]string,
) *SecurityScheme {
	return &SecurityScheme{
		securityType: securityType, description: description,
		name: name, in: in, flow: flow,
		authorizationURL: authorizationURL, tokenURL: tokenURL,
		scopes: scopes,
	}
}
