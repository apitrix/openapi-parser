package openapi31

// License provides license information for the API.
// https://spec.openapis.org/oas/v3.1.0#license-object
type License struct {
	Node // embedded - provides VendorExtensions and Trix

	name       string
	identifier string
	url        string
}

func (l *License) Name() string       { return l.name }
func (l *License) Identifier() string { return l.identifier }
func (l *License) URL() string        { return l.url }

// NewLicense creates a new License instance.
func NewLicense(name, identifier, url string) *License {
	return &License{name: name, identifier: identifier, url: url}
}
