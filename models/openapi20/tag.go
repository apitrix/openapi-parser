package openapi20

// Tag adds metadata to a single tag used by the Operation Object.
// https://swagger.io/specification/v2/#tag-object
type Tag struct {
	Node // embedded - provides VendorExtensions and Trix

	Name         string        `json:"name" yaml:"name"`
	Description  string        `json:"description,omitempty" yaml:"description,omitempty"`
	ExternalDocs *ExternalDocs `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
}

// NewTag creates a new Tag instance.
func NewTag(name string) *Tag {
	return &Tag{Name: name}
}
