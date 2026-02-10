package openapi31

// Contact provides contact information for the API.
// https://spec.openapis.org/oas/v3.1.0#contact-object
type Contact struct {
	Node // embedded - provides VendorExtensions and Trix

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
