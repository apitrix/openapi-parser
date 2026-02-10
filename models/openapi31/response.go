package openapi31

// Response describes a single response from an API operation.
// https://spec.openapis.org/oas/v3.1.0#response-object
type Response struct {
	Node // embedded - provides VendorExtensions and Trix

	Description string                `json:"description" yaml:"description"`
	Headers     map[string]*HeaderRef `json:"headers,omitempty" yaml:"headers,omitempty"`
	Content     map[string]*MediaType `json:"content,omitempty" yaml:"content,omitempty"`
	Links       map[string]*LinkRef   `json:"links,omitempty" yaml:"links,omitempty"`
}

// NewResponse creates a new Response instance.
func NewResponse(description string) *Response {
	return &Response{Description: description}
}
