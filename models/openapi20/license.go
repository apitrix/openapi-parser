package openapi20

import (
	"github.com/apitrix/openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// License provides license information for the API.
// https://swagger.io/specification/v2/#license-object
type License struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	name string
	url  string
}

func (l *License) Name() string { return l.name }
func (l *License) URL() string  { return l.url }

func (l *License) SetName(name string) error {
	if err := l.Trix.RunHooks("name", l.name, name); err != nil {
		return err
	}
	l.name = name
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
func NewLicense(name, url string) *License {
	return &License{name: name, url: url}
}

func (l *License) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "name", Value: l.name},
		{Key: "url", Value: l.url},
	}
	return shared.AppendExtensions(fields, l.VendorExtensions)
}

// MarshalFields implements shared.MarshalFieldsProvider for export.
func (l *License) MarshalFields() []shared.Field { return l.marshalFields() }

func (l *License) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(l.marshalFields())
}

func (l *License) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(l.marshalFields())
}

var _ yaml.Marshaler = (*License)(nil)
