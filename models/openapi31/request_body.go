package openapi31

// RequestBody describes a single request body.
// https://spec.openapis.org/oas/v3.1.0#request-body-object
type RequestBody struct {
	Node // embedded - provides VendorExtensions and Trix

	description string
	content     map[string]*MediaType
	required    bool
}

func (r *RequestBody) Description() string            { return r.description }
func (r *RequestBody) Content() map[string]*MediaType { return r.content }
func (r *RequestBody) Required() bool                 { return r.required }

// NewRequestBody creates a new RequestBody instance.
func NewRequestBody(description string, content map[string]*MediaType, required bool) *RequestBody {
	return &RequestBody{description: description, content: content, required: required}
}
