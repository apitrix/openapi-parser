package openapi31

// Info provides metadata about the API.
// https://spec.openapis.org/oas/v3.1.0#info-object
type Info struct {
	Node // embedded - provides VendorExtensions and Trix

	Title          string   `json:"title" yaml:"title"`
	Summary        string   `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description    string   `json:"description,omitempty" yaml:"description,omitempty"`
	TermsOfService string   `json:"termsOfService,omitempty" yaml:"termsOfService,omitempty"`
	Contact        *Contact `json:"contact,omitempty" yaml:"contact,omitempty"`
	License        *License `json:"license,omitempty" yaml:"license,omitempty"`
	Version        string   `json:"version" yaml:"version"`
}

// NewInfo creates a new Info instance.
func NewInfo(title, version string) *Info {
	return &Info{Title: title, Version: version}
}
