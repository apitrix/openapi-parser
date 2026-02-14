package openapi20

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// Contact provides contact information for the API.
// https://swagger.io/specification/v2/#contact-object
type Contact struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	name  string
	url   string
	email string
}

func (c *Contact) Name() string  { return c.name }
func (c *Contact) URL() string   { return c.url }
func (c *Contact) Email() string { return c.email }

// NewContact creates a new Contact instance.
func NewContact(name, url, email string) *Contact {
	return &Contact{name: name, url: url, email: email}
}

func (c *Contact) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "name", Value: c.name},
		{Key: "url", Value: c.url},
		{Key: "email", Value: c.email},
	}
	return shared.AppendExtensions(fields, c.VendorExtensions)
}

func (c *Contact) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(c.marshalFields())
}

func (c *Contact) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(c.marshalFields())
}

var _ yaml.Marshaler = (*Contact)(nil)
