package openapi20

// Contact provides contact information for the API.
// https://swagger.io/specification/v2/#contact-object
type Contact struct {
	Node // embedded - provides VendorExtensions and Trix

	Name  string `json:"name,omitempty" yaml:"name,omitempty"`
	URL   string `json:"url,omitempty" yaml:"url,omitempty"`
	Email string `json:"email,omitempty" yaml:"email,omitempty"`
}

// NewContact creates a new Contact instance.
func NewContact(name, url, email string) *Contact {
	return &Contact{Name: name, URL: url, Email: email}
}
