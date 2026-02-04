package openapi20

// Info provides metadata about the API.
// https://swagger.io/specification/v2/#info-object
type Info struct {
	Node // embedded - provides NodeSource and Extensions

	Title          string   `json:"title" yaml:"title"`
	Description    string   `json:"description,omitempty" yaml:"description,omitempty"`
	TermsOfService string   `json:"termsOfService,omitempty" yaml:"termsOfService,omitempty"`
	Contact        *Contact `json:"contact,omitempty" yaml:"contact,omitempty"`
	License        *License `json:"license,omitempty" yaml:"license,omitempty"`
	Version        string   `json:"version" yaml:"version"`
}

// Contact provides contact information for the API.
// https://swagger.io/specification/v2/#contact-object
type Contact struct {
	Node // embedded - provides NodeSource and Extensions

	Name  string `json:"name,omitempty" yaml:"name,omitempty"`
	URL   string `json:"url,omitempty" yaml:"url,omitempty"`
	Email string `json:"email,omitempty" yaml:"email,omitempty"`
}

// License provides license information for the API.
// https://swagger.io/specification/v2/#license-object
type License struct {
	Node // embedded - provides NodeSource and Extensions

	Name string `json:"name" yaml:"name"`
	URL  string `json:"url,omitempty" yaml:"url,omitempty"`
}
