package openapi31

import (
	"github.com/apitrix/openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// RequestBody describes a single request body.
// https://spec.openapis.org/oas/v3.1.0#request-body-object
type RequestBody struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	description string
	content     map[string]*MediaType
	required    bool
}

func (r *RequestBody) Description() string            { return r.description }
func (r *RequestBody) Content() map[string]*MediaType { return r.content }
func (r *RequestBody) Required() bool                 { return r.required }

func (r *RequestBody) SetDescription(description string) error {
	if err := r.Trix.RunHooks("description", r.description, description); err != nil {
		return err
	}
	r.description = description
	return nil
}
func (r *RequestBody) SetContent(content map[string]*MediaType) error {
	if err := r.Trix.RunHooks("content", r.content, content); err != nil {
		return err
	}
	r.content = content
	return nil
}
func (r *RequestBody) SetRequired(required bool) error {
	if err := r.Trix.RunHooks("required", r.required, required); err != nil {
		return err
	}
	r.required = required
	return nil
}

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

// MarshalFields implements shared.MarshalFieldsProvider for export.
func (r *RequestBody) MarshalFields() []shared.Field { return r.marshalFields() }

func (r *RequestBody) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(r.marshalFields())
}

func (r *RequestBody) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(r.marshalFields())
}

var _ yaml.Marshaler = (*RequestBody)(nil)
