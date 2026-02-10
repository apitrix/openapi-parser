package openapi30

// MediaType provides schema and examples for a media type.
// https://spec.openapis.org/oas/v3.0.3#media-type-object
type MediaType struct {
	Node // embedded - provides VendorExtensions and Trix

	schema   *SchemaRef
	example  interface{}
	examples map[string]*ExampleRef
	encoding map[string]*Encoding
}

func (mt *MediaType) Schema() *SchemaRef               { return mt.schema }
func (mt *MediaType) Example() interface{}             { return mt.example }
func (mt *MediaType) Examples() map[string]*ExampleRef { return mt.examples }
func (mt *MediaType) Encoding() map[string]*Encoding   { return mt.encoding }

// NewMediaType creates a new MediaType instance.
func NewMediaType(schema *SchemaRef, example interface{}, examples map[string]*ExampleRef, encoding map[string]*Encoding) *MediaType {
	return &MediaType{schema: schema, example: example, examples: examples, encoding: encoding}
}
