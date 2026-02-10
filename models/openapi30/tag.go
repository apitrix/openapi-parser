package openapi30

// Tag adds metadata to a single tag used by the Operation Object.
// https://spec.openapis.org/oas/v3.0.3#tag-object
type Tag struct {
	Node // embedded - provides VendorExtensions and Trix

	Name         string                 `json:"name" yaml:"name"`
	Description  string                 `json:"description,omitempty" yaml:"description,omitempty"`
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
}

// NewTag creates a new Tag instance.
func NewTag(name string) *Tag {
	return &Tag{Name: name}
}
