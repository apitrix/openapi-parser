package openapi30

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// Contact provides contact information for the API.
// https://spec.openapis.org/oas/v3.0.3#contact-object
type Contact struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	name  string
	url   string
	email string
}

// Name returns the identifying name of the contact person/organization.
func (c *Contact) Name() string { return c.name }

// URL returns the URL pointing to the contact information.
func (c *Contact) URL() string { return c.url }

// Email returns the email address of the contact person/organization.
func (c *Contact) Email() string { return c.email }

func (c *Contact) SetName(name string) error {
	if err := c.Trix.RunHooks("name", c.name, name); err != nil {
		return err
	}
	c.name = name
	return nil
}
func (c *Contact) SetURL(url string) error {
	if err := c.Trix.RunHooks("url", c.url, url); err != nil {
		return err
	}
	c.url = url
	return nil
}
func (c *Contact) SetEmail(email string) error {
	if err := c.Trix.RunHooks("email", c.email, email); err != nil {
		return err
	}
	c.email = email
	return nil
}

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

// MarshalFields implements shared.MarshalFieldsProvider for export.
func (c *Contact) MarshalFields() []shared.Field { return c.marshalFields() }

func (c *Contact) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(c.marshalFields())
}

func (c *Contact) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(c.marshalFields())
}

// Ensure Contact implements yaml.Marshaler.
var _ yaml.Marshaler = (*Contact)(nil)
