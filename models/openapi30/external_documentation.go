package openapi30

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// ExternalDocumentation allows referencing external documentation.
// https://spec.openapis.org/oas/v3.0.3#external-documentation-object
type ExternalDocumentation struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	description string
	url         string
}

func (e *ExternalDocumentation) Description() string { return e.description }
func (e *ExternalDocumentation) URL() string         { return e.url }

// NewExternalDocumentation creates a new ExternalDocumentation instance.
func NewExternalDocumentation(url, description string) *ExternalDocumentation {
	return &ExternalDocumentation{url: url, description: description}
}

func (e *ExternalDocumentation) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "description", Value: e.description},
		{Key: "url", Value: e.url},
	}
	return shared.AppendExtensions(fields, e.VendorExtensions)
}

func (e *ExternalDocumentation) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(e.marshalFields())
}

func (e *ExternalDocumentation) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(e.marshalFields())
}

var _ yaml.Marshaler = (*ExternalDocumentation)(nil)
