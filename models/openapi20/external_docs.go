package openapi20

// ExternalDocs allows referencing external documentation.
// https://swagger.io/specification/v2/#external-documentation-object
type ExternalDocs struct {
	Node // embedded - provides VendorExtensions and Trix

	description string
	url         string
}

func (ed *ExternalDocs) Description() string { return ed.description }
func (ed *ExternalDocs) URL() string         { return ed.url }

// NewExternalDocs creates a new ExternalDocs instance.
func NewExternalDocs(description, url string) *ExternalDocs {
	return &ExternalDocs{description: description, url: url}
}
