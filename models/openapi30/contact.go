package openapi30

// Contact provides contact information for the API.
// https://spec.openapis.org/oas/v3.0.3#contact-object
type Contact struct {
	Node // embedded - provides VendorExtensions and Trix

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

// NewContact creates a new Contact instance.
func NewContact(name, url, email string) *Contact {
	return &Contact{name: name, url: url, email: email}
}
