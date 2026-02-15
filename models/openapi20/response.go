package openapi20

import (
	"github.com/apitrix/openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// Response describes a single response from an API operation.
// https://swagger.io/specification/v2/#response-object
type Response struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	description string
	schema      *RefSchema
	headers     map[string]*Header
	examples    map[string]interface{}
}

func (r *Response) Description() string              { return r.description }
func (r *Response) Schema() *RefSchema               { return r.schema }
func (r *Response) Headers() map[string]*Header      { return r.headers }
func (r *Response) Examples() map[string]interface{} { return r.examples }

func (r *Response) SetDescription(description string) error {
	if err := r.Trix.RunHooks("description", r.description, description); err != nil {
		return err
	}
	r.description = description
	return nil
}
func (r *Response) SetSchema(schema *RefSchema) error {
	if err := r.Trix.RunHooks("schema", r.schema, schema); err != nil {
		return err
	}
	r.schema = schema
	return nil
}
func (r *Response) SetHeaders(headers map[string]*Header) error {
	if err := r.Trix.RunHooks("headers", r.headers, headers); err != nil {
		return err
	}
	r.headers = headers
	return nil
}
func (r *Response) SetExamples(examples map[string]interface{}) error {
	if err := r.Trix.RunHooks("examples", r.examples, examples); err != nil {
		return err
	}
	r.examples = examples
	return nil
}

// NewResponse creates a new Response instance.
func NewResponse(description string, schema *RefSchema, headers map[string]*Header, examples map[string]interface{}) *Response {
	return &Response{
		description: description, schema: schema,
		headers: headers, examples: examples,
	}
}

func (r *Response) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "description", Value: r.description},
		{Key: "schema", Value: r.schema},
		{Key: "headers", Value: r.headers},
		{Key: "examples", Value: r.examples},
	}
	return shared.AppendExtensions(fields, r.VendorExtensions)
}

// MarshalFields implements shared.MarshalFieldsProvider for export.
func (r *Response) MarshalFields() []shared.Field { return r.marshalFields() }

func (r *Response) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(r.marshalFields())
}

func (r *Response) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(r.marshalFields())
}

var _ yaml.Marshaler = (*Response)(nil)
