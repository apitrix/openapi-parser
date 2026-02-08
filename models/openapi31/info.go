package openapi31

// Info provides metadata about the API.
// https://spec.openapis.org/oas/v3.1.0#info-object
type Info struct {
	Node // embedded - provides NodeSource and Extensions

	Title          string   `json:"title" yaml:"title"`
	Summary        string   `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description    string   `json:"description,omitempty" yaml:"description,omitempty"`
	TermsOfService string   `json:"termsOfService,omitempty" yaml:"termsOfService,omitempty"`
	Contact        *Contact `json:"contact,omitempty" yaml:"contact,omitempty"`
	License        *License `json:"license,omitempty" yaml:"license,omitempty"`
	Version        string   `json:"version" yaml:"version"`
}

// Contact provides contact information for the API.
// https://spec.openapis.org/oas/v3.1.0#contact-object
type Contact struct {
	Node // embedded - provides NodeSource and Extensions

	Name  string `json:"name,omitempty" yaml:"name,omitempty"`
	URL   string `json:"url,omitempty" yaml:"url,omitempty"`
	Email string `json:"email,omitempty" yaml:"email,omitempty"`
}

// License provides license information for the API.
// https://spec.openapis.org/oas/v3.1.0#license-object
type License struct {
	Node // embedded - provides NodeSource and Extensions

	Name       string `json:"name" yaml:"name"`
	Identifier string `json:"identifier,omitempty" yaml:"identifier,omitempty"`
	URL        string `json:"url,omitempty" yaml:"url,omitempty"`
}
