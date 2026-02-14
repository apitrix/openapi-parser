package openapi31

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// Encoding defines encoding for a single schema property.
// https://spec.openapis.org/oas/v3.1.0#encoding-object
type Encoding struct {
	Node // embedded - provides VendorExtensions and Trix

	contentType   string
	headers       map[string]*shared.RefWithMeta[Header]
	style         string
	explode       *bool
	allowReserved bool
}

func (e *Encoding) ContentType() string            { return e.contentType }
func (e *Encoding) Headers() map[string]*shared.RefWithMeta[Header] { return e.headers }
func (e *Encoding) Style() string                  { return e.style }
func (e *Encoding) Explode() *bool                 { return e.explode }
func (e *Encoding) AllowReserved() bool            { return e.allowReserved }

// NewEncoding creates a new Encoding instance.
func NewEncoding(contentType, style string, headers map[string]*shared.RefWithMeta[Header], explode *bool, allowReserved bool) *Encoding {
	return &Encoding{
		contentType: contentType, headers: headers, style: style,
		explode: explode, allowReserved: allowReserved,
	}
}

func (e *Encoding) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "contentType", Value: e.contentType},
		{Key: "headers", Value: e.headers},
		{Key: "style", Value: e.style},
		{Key: "explode", Value: e.explode},
		{Key: "allowReserved", Value: e.allowReserved},
	}
	return shared.AppendExtensions(fields, e.VendorExtensions)
}

func (e *Encoding) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(e.marshalFields())
}

func (e *Encoding) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(e.marshalFields())
}

var _ yaml.Marshaler = (*Encoding)(nil)
