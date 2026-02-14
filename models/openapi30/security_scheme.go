package openapi30

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// SecurityScheme defines a security scheme for the API.
// https://spec.openapis.org/oas/v3.0.3#security-scheme-object
type SecurityScheme struct {
	ElementBase // embedded - provides VendorExtensions and Trix

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

func (ss *SecurityScheme) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "type", Value: ss.secType},
		{Key: "description", Value: ss.description},
		{Key: "name", Value: ss.name},
		{Key: "in", Value: ss.in},
		{Key: "scheme", Value: ss.scheme},
		{Key: "bearerFormat", Value: ss.bearerFormat},
		{Key: "flows", Value: ss.flows},
		{Key: "openIdConnectUrl", Value: ss.openIDConnectURL},
	}
	return shared.AppendExtensions(fields, ss.VendorExtensions)
}

func (ss *SecurityScheme) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(ss.marshalFields())
}

func (ss *SecurityScheme) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(ss.marshalFields())
}

var _ yaml.Marshaler = (*SecurityScheme)(nil)
