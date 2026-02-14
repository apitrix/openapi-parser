package openapi31

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// Response describes a single response from an API operation.
// https://spec.openapis.org/oas/v3.1.0#response-object
type Response struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	description string
	headers     map[string]*RefHeader
	content     map[string]*MediaType
	links       map[string]*RefLink
}

func (r *Response) Description() string            { return r.description }
func (r *Response) Headers() map[string]*RefHeader { return r.headers }
func (r *Response) Content() map[string]*MediaType { return r.content }
func (r *Response) Links() map[string]*RefLink     { return r.links }

func (r *Response) SetDescription(description string) error {
	if err := r.Trix.RunHooks("description", r.description, description); err != nil {
		return err
	}
	r.description = description
	return nil
}
func (r *Response) SetHeaders(headers map[string]*RefHeader) error {
	if err := r.Trix.RunHooks("headers", r.headers, headers); err != nil {
		return err
	}
	r.headers = headers
	return nil
}
func (r *Response) SetContent(content map[string]*MediaType) error {
	if err := r.Trix.RunHooks("content", r.content, content); err != nil {
		return err
	}
	r.content = content
	return nil
}
func (r *Response) SetLinks(links map[string]*RefLink) error {
	if err := r.Trix.RunHooks("links", r.links, links); err != nil {
		return err
	}
	r.links = links
	return nil
}

// NewResponse creates a new Response instance.
func NewResponse(description string, headers map[string]*RefHeader, content map[string]*MediaType, links map[string]*RefLink) *Response {
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

// MarshalFields implements shared.MarshalFieldsProvider for export.
func (r *Response) MarshalFields() []shared.Field { return r.marshalFields() }

func (r *Response) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(r.marshalFields())
}

func (r *Response) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(r.marshalFields())
}

var _ yaml.Marshaler = (*Response)(nil)
