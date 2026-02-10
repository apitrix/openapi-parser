package openapi30

// ExternalDocumentation allows referencing external documentation.
// https://spec.openapis.org/oas/v3.0.3#external-documentation-object
type ExternalDocumentation struct {
	Node // embedded - provides VendorExtensions and Trix

	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	URL         string `json:"url" yaml:"url"`
}

// NewExternalDocumentation creates a new ExternalDocumentation instance.
func NewExternalDocumentation(url string) *ExternalDocumentation {
	return &ExternalDocumentation{URL: url}
}
