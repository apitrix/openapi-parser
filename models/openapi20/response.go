package openapi20

// Response describes a single response from an API operation.
// https://swagger.io/specification/v2/#response-object
type Response struct {
	Node // embedded - provides VendorExtensions and Trix

	Description string                 `json:"description" yaml:"description"`
	Schema      *SchemaRef             `json:"schema,omitempty" yaml:"schema,omitempty"`
	Headers     map[string]*Header     `json:"headers,omitempty" yaml:"headers,omitempty"`
	Examples    map[string]interface{} `json:"examples,omitempty" yaml:"examples,omitempty"`
}

// NewResponse creates a new Response instance.
func NewResponse(description string) *Response {
	return &Response{Description: description}
}
