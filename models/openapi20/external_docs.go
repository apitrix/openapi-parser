package openapi20

import (
	"openapi-parser/models/shared"

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

func (ed *ExternalDocs) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(ed.marshalFields())
}

func (ed *ExternalDocs) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(ed.marshalFields())
}

var _ yaml.Marshaler = (*ExternalDocs)(nil)
