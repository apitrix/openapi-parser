package openapi30

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// RequestBody describes a single request body.
// https://spec.openapis.org/oas/v3.0.3#request-body-object
type RequestBody struct {
	ElementBase // embedded - provides VendorExtensions and Trix

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

func (rb *RequestBody) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "description", Value: rb.description},
		{Key: "content", Value: rb.content},
		{Key: "required", Value: rb.required},
	}
	return shared.AppendExtensions(fields, rb.VendorExtensions)
}

func (rb *RequestBody) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(rb.marshalFields())
}

func (rb *RequestBody) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(rb.marshalFields())
}

var _ yaml.Marshaler = (*RequestBody)(nil)
