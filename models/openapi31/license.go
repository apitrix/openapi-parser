package openapi31

// License provides license information for the API.
// https://spec.openapis.org/oas/v3.1.0#license-object
type License struct {
	Node // embedded - provides VendorExtensions and Trix

	Name       string `json:"name" yaml:"name"`
	Identifier string `json:"identifier,omitempty" yaml:"identifier,omitempty"`
	URL        string `json:"url,omitempty" yaml:"url,omitempty"`
}

// NewLicense creates a new License instance.
func NewLicense(name string) *License {
	return &License{Name: name}
}
