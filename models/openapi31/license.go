package openapi31

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// License provides license information for the API.
// https://spec.openapis.org/oas/v3.1.0#license-object
type License struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	name       string
	identifier string
	url        string
}

func (l *License) Name() string       { return l.name }
func (l *License) Identifier() string { return l.identifier }
func (l *License) URL() string        { return l.url }

func (l *License) SetName(name string) error {
	if err := l.Trix.RunHooks("name", l.name, name); err != nil {
		return err
	}
	l.name = name
	return nil
}
func (l *License) SetIdentifier(identifier string) error {
	if err := l.Trix.RunHooks("identifier", l.identifier, identifier); err != nil {
		return err
	}
	l.identifier = identifier
	return nil
}
func (l *License) SetURL(url string) error {
	if err := l.Trix.RunHooks("url", l.url, url); err != nil {
		return err
	}
	l.url = url
	return nil
}

// NewLicense creates a new License instance.
func NewLicense(name, identifier, url string) *License {
	return &License{name: name, identifier: identifier, url: url}
}

func (l *License) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "name", Value: l.name},
		{Key: "identifier", Value: l.identifier},
		{Key: "url", Value: l.url},
	}
	return shared.AppendExtensions(fields, l.VendorExtensions)
}

func (l *License) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(l.marshalFields())
}

func (l *License) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(l.marshalFields())
}

var _ yaml.Marshaler = (*License)(nil)
