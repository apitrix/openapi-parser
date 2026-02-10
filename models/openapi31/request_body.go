package openapi31

// RequestBody describes a single request body.
// https://spec.openapis.org/oas/v3.1.0#request-body-object
type RequestBody struct {
	Node // embedded - provides VendorExtensions and Trix

	Description string                `json:"description,omitempty" yaml:"description,omitempty"`
	Content     map[string]*MediaType `json:"content" yaml:"content"`
	Required    bool                  `json:"required,omitempty" yaml:"required,omitempty"`
}

// NewRequestBody creates a new RequestBody instance.
func NewRequestBody() *RequestBody {
	return &RequestBody{}
}
