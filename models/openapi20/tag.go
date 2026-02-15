package openapi20

import (
	"github.com/apitrix/openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// Tag adds metadata to a single tag used by the Operation Object.
// https://swagger.io/specification/v2/#tag-object
type Tag struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	name         string
	description  string
	externalDocs *ExternalDocs
}

func (t *Tag) Name() string                { return t.name }
func (t *Tag) Description() string         { return t.description }
func (t *Tag) ExternalDocs() *ExternalDocs { return t.externalDocs }

func (t *Tag) SetName(name string) error {
	if err := t.Trix.RunHooks("name", t.name, name); err != nil {
		return err
	}
	t.name = name
	return nil
}
func (t *Tag) SetDescription(description string) error {
	if err := t.Trix.RunHooks("description", t.description, description); err != nil {
		return err
	}
	t.description = description
	return nil
}
func (t *Tag) SetExternalDocs(externalDocs *ExternalDocs) error {
	if err := t.Trix.RunHooks("externalDocs", t.externalDocs, externalDocs); err != nil {
		return err
	}
	t.externalDocs = externalDocs
	return nil
}

// NewTag creates a new Tag instance.
func NewTag(name, description string, externalDocs *ExternalDocs) *Tag {
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

// MarshalFields implements shared.MarshalFieldsProvider for export.
func (t *Tag) MarshalFields() []shared.Field { return t.marshalFields() }

func (t *Tag) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(t.marshalFields())
}

func (t *Tag) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(t.marshalFields())
}

var _ yaml.Marshaler = (*Tag)(nil)
