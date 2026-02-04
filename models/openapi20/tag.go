package openapi20

// Tag adds metadata to a single tag used by the Operation Object.
// https://swagger.io/specification/v2/#tag-object
type Tag struct {
	Node // embedded - provides NodeSource and Extensions

	Name         string        `json:"name" yaml:"name"`
	Description  string        `json:"description,omitempty" yaml:"description,omitempty"`
	ExternalDocs *ExternalDocs `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
}

// ExternalDocs allows referencing external documentation.
// https://swagger.io/specification/v2/#external-documentation-object
type ExternalDocs struct {
	Node // embedded - provides NodeSource and Extensions

	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	URL         string `json:"url" yaml:"url"`
}
