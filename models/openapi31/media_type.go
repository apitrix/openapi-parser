package openapi31

// MediaType provides schema and examples for a media type.
// https://spec.openapis.org/oas/v3.1.0#media-type-object
type MediaType struct {
	Node // embedded - provides VendorExtensions and Trix

	Schema   *SchemaRef             `json:"schema,omitempty" yaml:"schema,omitempty"`
	Example  interface{}            `json:"example,omitempty" yaml:"example,omitempty"`
	Examples map[string]*ExampleRef `json:"examples,omitempty" yaml:"examples,omitempty"`
	Encoding map[string]*Encoding   `json:"encoding,omitempty" yaml:"encoding,omitempty"`
}

// NewMediaType creates a new MediaType instance.
func NewMediaType() *MediaType {
	return &MediaType{}
}
