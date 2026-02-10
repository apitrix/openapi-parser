package openapi31

// ExternalDocumentation allows referencing external documentation.
// https://spec.openapis.org/oas/v3.1.0#external-documentation-object
type ExternalDocumentation struct {
	Node // embedded - provides VendorExtensions and Trix

	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	URL         string `json:"url" yaml:"url"`
}

// NewExternalDocumentation creates a new ExternalDocumentation instance.
func NewExternalDocumentation(url string) *ExternalDocumentation {
	return &ExternalDocumentation{URL: url}
}
