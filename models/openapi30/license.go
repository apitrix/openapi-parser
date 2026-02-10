package openapi30

// License provides license information for the API.
// https://spec.openapis.org/oas/v3.0.3#license-object
type License struct {
	Node // embedded - provides VendorExtensions and Trix

	name string
	url  string
}

// Name returns the license name used for the API.
func (l *License) Name() string { return l.name }

// URL returns the URL to the license used for the API.
func (l *License) URL() string { return l.url }

// NewLicense creates a new License instance.
func NewLicense(name, url string) *License {
	return &License{name: name, url: url}
}
