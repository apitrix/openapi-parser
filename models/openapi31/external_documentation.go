package openapi31

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// ExternalDocumentation allows referencing external documentation.
// https://spec.openapis.org/oas/v3.1.0#external-documentation-object
type ExternalDocumentation struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	description string
	url         string
}

func (e *ExternalDocumentation) Description() string { return e.description }
func (e *ExternalDocumentation) URL() string         { return e.url }

func (e *ExternalDocumentation) SetDescription(description string) error {
	if err := e.Trix.RunHooks("description", e.description, description); err != nil {
		return err
	}
	e.description = description
	return nil
}
func (e *ExternalDocumentation) SetURL(url string) error {
	if err := e.Trix.RunHooks("url", e.url, url); err != nil {
		return err
	}
	e.url = url
	return nil
}

// NewExternalDocumentation creates a new ExternalDocumentation instance.
func NewExternalDocumentation(description, url string) *ExternalDocumentation {
	return &ExternalDocumentation{description: description, url: url}
}

func (e *ExternalDocumentation) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "description", Value: e.description},
		{Key: "url", Value: e.url},
	}
	return shared.AppendExtensions(fields, e.VendorExtensions)
}

// MarshalFields implements shared.MarshalFieldsProvider for export.
func (e *ExternalDocumentation) MarshalFields() []shared.Field { return e.marshalFields() }

func (e *ExternalDocumentation) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(e.marshalFields())
}

func (e *ExternalDocumentation) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(e.marshalFields())
}

var _ yaml.Marshaler = (*ExternalDocumentation)(nil)
