package openapi30

// License provides license information for the API.
// https://spec.openapis.org/oas/v3.0.3#license-object
type License struct {
	Node // embedded - provides VendorExtensions and Trix

	Name string `json:"name" yaml:"name"`
	URL  string `json:"url,omitempty" yaml:"url,omitempty"`
}

// NewLicense creates a new License instance.
func NewLicense(name, url string) *License {
	return &License{Name: name, URL: url}
}
