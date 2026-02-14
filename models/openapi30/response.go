package openapi30

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// Response describes a single response from an API operation.
// https://spec.openapis.org/oas/v3.0.3#response-object
type Response struct {
	Node // embedded - provides VendorExtensions and Trix

	description string
	headers     map[string]*shared.Ref[Header]
	content     map[string]*MediaType
	links       map[string]*shared.Ref[Link]
}

func (r *Response) Description() string            { return r.description }
func (r *Response) Headers() map[string]*shared.Ref[Header] { return r.headers }
func (r *Response) Content() map[string]*MediaType { return r.content }
func (r *Response) Links() map[string]*shared.Ref[Link]     { return r.links }

// NewResponse creates a new Response instance.
func NewResponse(description string, headers map[string]*shared.Ref[Header], content map[string]*MediaType, links map[string]*shared.Ref[Link]) *Response {
	return &Response{description: description, headers: headers, content: content, links: links}
}

func (r *Response) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "description", Value: r.description},
		{Key: "headers", Value: r.headers},
		{Key: "content", Value: r.content},
		{Key: "links", Value: r.links},
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
