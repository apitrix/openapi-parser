package openapi30

// Tag adds metadata to a single tag used by the Operation Object.
// https://spec.openapis.org/oas/v3.0.3#tag-object
type Tag struct {
	Node // embedded - provides VendorExtensions and Trix

	name         string
	description  string
	externalDocs *ExternalDocumentation
}

func (t *Tag) Name() string                         { return t.name }
func (t *Tag) Description() string                  { return t.description }
func (t *Tag) ExternalDocs() *ExternalDocumentation { return t.externalDocs }

// NewTag creates a new Tag instance.
func NewTag(name, description string, externalDocs *ExternalDocumentation) *Tag {
	return &Tag{name: name, description: description, externalDocs: externalDocs}
}
