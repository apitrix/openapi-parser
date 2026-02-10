package openapi31

// MediaType provides schema and examples for a media type.
// https://spec.openapis.org/oas/v3.1.0#media-type-object
type MediaType struct {
	Node // embedded - provides VendorExtensions and Trix

	schema   *SchemaRef
	example  interface{}
	examples map[string]*ExampleRef
	encoding map[string]*Encoding
}

func (m *MediaType) Schema() *SchemaRef               { return m.schema }
func (m *MediaType) Example() interface{}             { return m.example }
func (m *MediaType) Examples() map[string]*ExampleRef { return m.examples }
func (m *MediaType) Encoding() map[string]*Encoding   { return m.encoding }

// NewMediaType creates a new MediaType instance.
func NewMediaType(schema *SchemaRef, example interface{}, examples map[string]*ExampleRef, encoding map[string]*Encoding) *MediaType {
	return &MediaType{schema: schema, example: example, examples: examples, encoding: encoding}
}
