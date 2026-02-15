package openapi30

import (
	"github.com/apitrix/openapi-parser/models/shared"

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

func (ss *SecurityScheme) SetType(secType string) error {
	if err := ss.Trix.RunHooks("type", ss.secType, secType); err != nil {
		return err
	}
	ss.secType = secType
	return nil
}
func (ss *SecurityScheme) SetDescription(description string) error {
	if err := ss.Trix.RunHooks("description", ss.description, description); err != nil {
		return err
	}
	ss.description = description
	return nil
}
func (ss *SecurityScheme) SetName(name string) error {
	if err := ss.Trix.RunHooks("name", ss.name, name); err != nil {
		return err
	}
	ss.name = name
	return nil
}
func (ss *SecurityScheme) SetIn(in string) error {
	if err := ss.Trix.RunHooks("in", ss.in, in); err != nil {
		return err
	}
	ss.in = in
	return nil
}
func (ss *SecurityScheme) SetScheme(scheme string) error {
	if err := ss.Trix.RunHooks("scheme", ss.scheme, scheme); err != nil {
		return err
	}
	ss.scheme = scheme
	return nil
}
func (ss *SecurityScheme) SetBearerFormat(bearerFormat string) error {
	if err := ss.Trix.RunHooks("bearerFormat", ss.bearerFormat, bearerFormat); err != nil {
		return err
	}
	ss.bearerFormat = bearerFormat
	return nil
}
func (ss *SecurityScheme) SetFlows(flows *OAuthFlows) error {
	if err := ss.Trix.RunHooks("flows", ss.flows, flows); err != nil {
		return err
	}
	ss.flows = flows
	return nil
}
func (ss *SecurityScheme) SetOpenIDConnectURL(openIDConnectURL string) error {
	if err := ss.Trix.RunHooks("openIdConnectUrl", ss.openIDConnectURL, openIDConnectURL); err != nil {
		return err
	}
	ss.openIDConnectURL = openIDConnectURL
	return nil
}

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

// MarshalFields implements shared.MarshalFieldsProvider for export.
func (ss *SecurityScheme) MarshalFields() []shared.Field { return ss.marshalFields() }

func (ss *SecurityScheme) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(ss.marshalFields())
}

func (ss *SecurityScheme) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(ss.marshalFields())
}

var _ yaml.Marshaler = (*SecurityScheme)(nil)
