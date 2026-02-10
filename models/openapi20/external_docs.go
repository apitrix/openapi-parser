package openapi20

// ExternalDocs allows referencing external documentation.
// https://swagger.io/specification/v2/#external-documentation-object
type ExternalDocs struct {
	Node // embedded - provides VendorExtensions and Trix

	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	URL         string `json:"url" yaml:"url"`
}

// NewExternalDocs creates a new ExternalDocs instance.
func NewExternalDocs(url string) *ExternalDocs {
	return &ExternalDocs{URL: url}
}
