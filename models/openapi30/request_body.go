package openapi30

// RequestBody describes a single request body.
// https://spec.openapis.org/oas/v3.0.3#request-body-object
type RequestBody struct {
	Node // embedded - provides VendorExtensions and Trix

	description string
	content     map[string]*MediaType
	required    bool
}

func (rb *RequestBody) Description() string            { return rb.description }
func (rb *RequestBody) Content() map[string]*MediaType { return rb.content }
func (rb *RequestBody) Required() bool                 { return rb.required }

// NewRequestBody creates a new RequestBody instance.
func NewRequestBody(description string, content map[string]*MediaType, required bool) *RequestBody {
	return &RequestBody{description: description, content: content, required: required}
}
