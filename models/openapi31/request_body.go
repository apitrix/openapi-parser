package openapi31

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

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

func (r *RequestBody) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "description", Value: r.description},
		{Key: "content", Value: r.content},
		{Key: "required", Value: r.required},
	}
	return shared.AppendExtensions(fields, r.VendorExtensions)
}

func (r *RequestBody) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(r.marshalFields())
}

func (r *RequestBody) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(r.marshalFields())
}

var _ yaml.Marshaler = (*RequestBody)(nil)
