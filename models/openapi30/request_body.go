package openapi30

import (
	"github.com/apitrix/openapi-parser/models/shared"

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

func (rb *RequestBody) SetDescription(description string) error {
	if err := rb.Trix.RunHooks("description", rb.description, description); err != nil {
		return err
	}
	rb.description = description
	return nil
}
func (rb *RequestBody) SetContent(content map[string]*MediaType) error {
	if err := rb.Trix.RunHooks("content", rb.content, content); err != nil {
		return err
	}
	rb.content = content
	return nil
}
func (rb *RequestBody) SetRequired(required bool) error {
	if err := rb.Trix.RunHooks("required", rb.required, required); err != nil {
		return err
	}
	rb.required = required
	return nil
}

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

// MarshalFields implements shared.MarshalFieldsProvider for export.
func (rb *RequestBody) MarshalFields() []shared.Field { return rb.marshalFields() }

func (rb *RequestBody) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(rb.marshalFields())
}

func (rb *RequestBody) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(rb.marshalFields())
}

var _ yaml.Marshaler = (*RequestBody)(nil)
