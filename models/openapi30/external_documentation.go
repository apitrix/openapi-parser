package openapi30

// ExternalDocumentation allows referencing external documentation.
// https://spec.openapis.org/oas/v3.0.3#external-documentation-object
type ExternalDocumentation struct {
	Node // embedded - provides VendorExtensions and Trix

	description string
	url         string
}

func (e *ExternalDocumentation) Description() string { return e.description }
func (e *ExternalDocumentation) URL() string         { return e.url }

// NewExternalDocumentation creates a new ExternalDocumentation instance.
func NewExternalDocumentation(url, description string) *ExternalDocumentation {
	return &ExternalDocumentation{url: url, description: description}
}
