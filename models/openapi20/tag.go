package openapi20

// Tag adds metadata to a single tag used by the Operation Object.
// https://swagger.io/specification/v2/#tag-object
type Tag struct {
	Node // embedded - provides VendorExtensions and Trix

	name         string
	description  string
	externalDocs *ExternalDocs
}

func (t *Tag) Name() string                { return t.name }
func (t *Tag) Description() string         { return t.description }
func (t *Tag) ExternalDocs() *ExternalDocs { return t.externalDocs }

// NewTag creates a new Tag instance.
func NewTag(name, description string, externalDocs *ExternalDocs) *Tag {
	return &Tag{name: name, description: description, externalDocs: externalDocs}
}
