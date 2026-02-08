package openapi31

// RequestBody describes a single request body.
// https://spec.openapis.org/oas/v3.1.0#request-body-object
type RequestBody struct {
	Node // embedded - provides VendorExtensions and Trix

	Description string                `json:"description,omitempty" yaml:"description,omitempty"`
	Content     map[string]*MediaType `json:"content" yaml:"content"`
	Required    bool                  `json:"required,omitempty" yaml:"required,omitempty"`
}

// MediaType provides schema and examples for a media type.
// https://spec.openapis.org/oas/v3.1.0#media-type-object
type MediaType struct {
	Node // embedded - provides VendorExtensions and Trix

	Schema   *SchemaRef             `json:"schema,omitempty" yaml:"schema,omitempty"`
	Example  interface{}            `json:"example,omitempty" yaml:"example,omitempty"`
	Examples map[string]*ExampleRef `json:"examples,omitempty" yaml:"examples,omitempty"`
	Encoding map[string]*Encoding   `json:"encoding,omitempty" yaml:"encoding,omitempty"`
}

// Encoding defines encoding for a single schema property.
// https://spec.openapis.org/oas/v3.1.0#encoding-object
type Encoding struct {
	Node // embedded - provides VendorExtensions and Trix

	ContentType   string                `json:"contentType,omitempty" yaml:"contentType,omitempty"`
	Headers       map[string]*HeaderRef `json:"headers,omitempty" yaml:"headers,omitempty"`
	Style         string                `json:"style,omitempty" yaml:"style,omitempty"`
	Explode       *bool                 `json:"explode,omitempty" yaml:"explode,omitempty"`
	AllowReserved bool                  `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"`
}
