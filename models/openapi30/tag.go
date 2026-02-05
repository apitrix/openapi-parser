package openapi30

// Tag adds metadata to a single tag used by the Operation Object.
// https://spec.openapis.org/oas/v3.0.3#tag-object
type Tag struct {
	Node // embedded - provides NodeSource and Extensions

	Name         string                 `json:"name" yaml:"name"`
	Description  string                 `json:"description,omitempty" yaml:"description,omitempty"`
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
}

// ExternalDocumentation allows referencing external documentation.
// https://spec.openapis.org/oas/v3.0.3#external-documentation-object
type ExternalDocumentation struct {
	Node // embedded - provides NodeSource and Extensions

	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	URL         string `json:"url" yaml:"url"`
}
