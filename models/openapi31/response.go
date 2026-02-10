package openapi31

// Response describes a single response from an API operation.
// https://spec.openapis.org/oas/v3.1.0#response-object
type Response struct {
	Node // embedded - provides VendorExtensions and Trix

	description string
	headers     map[string]*HeaderRef
	content     map[string]*MediaType
	links       map[string]*LinkRef
}

func (r *Response) Description() string            { return r.description }
func (r *Response) Headers() map[string]*HeaderRef { return r.headers }
func (r *Response) Content() map[string]*MediaType { return r.content }
func (r *Response) Links() map[string]*LinkRef     { return r.links }

// NewResponse creates a new Response instance.
func NewResponse(description string, headers map[string]*HeaderRef, content map[string]*MediaType, links map[string]*LinkRef) *Response {
	return &Response{description: description, headers: headers, content: content, links: links}
}
