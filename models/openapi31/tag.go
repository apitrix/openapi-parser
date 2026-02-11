package openapi31

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// Tag adds metadata to a single tag used by the Operation Object.
// https://spec.openapis.org/oas/v3.1.0#tag-object
type Tag struct {
	Node // embedded - provides VendorExtensions and Trix

	name         string
	description  string
	externalDocs *ExternalDocumentation
}

func (t *Tag) Name() string                         { return t.name }
func (t *Tag) Description() string                  { return t.description }
func (t *Tag) ExternalDocs() *ExternalDocumentation { return t.externalDocs }

// NewTag creates a new Tag instance.
func NewTag(name, description string, externalDocs *ExternalDocumentation) *Tag {
	return &Tag{name: name, description: description, externalDocs: externalDocs}
}

func (t *Tag) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "name", Value: t.name},
		{Key: "description", Value: t.description},
		{Key: "externalDocs", Value: t.externalDocs},
	}
	return shared.AppendExtensions(fields, t.VendorExtensions)
}

func (t *Tag) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(t.marshalFields())
}

func (t *Tag) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(t.marshalFields())
}

var _ yaml.Marshaler = (*Tag)(nil)
