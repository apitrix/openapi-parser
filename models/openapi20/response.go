package openapi20

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// Response describes a single response from an API operation.
// https://swagger.io/specification/v2/#response-object
type Response struct {
	Node // embedded - provides VendorExtensions and Trix

	description string
	schema      *SchemaRef
	headers     map[string]*Header
	examples    map[string]interface{}
}

func (r *Response) Description() string              { return r.description }
func (r *Response) Schema() *SchemaRef               { return r.schema }
func (r *Response) Headers() map[string]*Header      { return r.headers }
func (r *Response) Examples() map[string]interface{} { return r.examples }

// NewResponse creates a new Response instance.
func NewResponse(description string, schema *SchemaRef, headers map[string]*Header, examples map[string]interface{}) *Response {
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

func (r *Response) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(r.marshalFields())
}

func (r *Response) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(r.marshalFields())
}

var _ yaml.Marshaler = (*Response)(nil)
