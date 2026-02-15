package openapi20

import (
	"github.com/apitrix/openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// ExternalDocs allows referencing external documentation.
// https://swagger.io/specification/v2/#external-documentation-object
type ExternalDocs struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	description string
	url         string
}

func (ed *ExternalDocs) Description() string { return ed.description }
func (ed *ExternalDocs) URL() string         { return ed.url }

func (ed *ExternalDocs) SetDescription(description string) error {
	if err := ed.Trix.RunHooks("description", ed.description, description); err != nil {
		return err
	}
	ed.description = description
	return nil
}
func (ed *ExternalDocs) SetURL(url string) error {
	if err := ed.Trix.RunHooks("url", ed.url, url); err != nil {
		return err
	}
	ed.url = url
	return nil
}

// NewExternalDocs creates a new ExternalDocs instance.
func NewExternalDocs(description, url string) *ExternalDocs {
	return &ExternalDocs{description: description, url: url}
}

func (ed *ExternalDocs) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "description", Value: ed.description},
		{Key: "url", Value: ed.url},
	}
	return shared.AppendExtensions(fields, ed.VendorExtensions)
}

// MarshalFields implements shared.MarshalFieldsProvider for export.
func (ed *ExternalDocs) MarshalFields() []shared.Field { return ed.marshalFields() }

func (ed *ExternalDocs) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(ed.marshalFields())
}

func (ed *ExternalDocs) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(ed.marshalFields())
}

var _ yaml.Marshaler = (*ExternalDocs)(nil)
