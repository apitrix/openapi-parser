package openapi20

// License provides license information for the API.
// https://swagger.io/specification/v2/#license-object
type License struct {
	Node // embedded - provides VendorExtensions and Trix

	name string
	url  string
}

func (l *License) Name() string { return l.name }
func (l *License) URL() string  { return l.url }

// NewLicense creates a new License instance.
func NewLicense(name, url string) *License {
	return &License{name: name, url: url}
}
