package openapi30

import (
	"github.com/apitrix/openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// Encoding defines encoding for a single schema property.
// https://spec.openapis.org/oas/v3.0.3#encoding-object
type Encoding struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	contentType   string
	headers       map[string]*RefHeader
	style         string
	explode       *bool
	allowReserved bool
}

func (e *Encoding) ContentType() string            { return e.contentType }
func (e *Encoding) Headers() map[string]*RefHeader { return e.headers }
func (e *Encoding) Style() string                  { return e.style }
func (e *Encoding) Explode() *bool                 { return e.explode }
func (e *Encoding) AllowReserved() bool            { return e.allowReserved }

func (e *Encoding) SetContentType(contentType string) error {
	if err := e.Trix.RunHooks("contentType", e.contentType, contentType); err != nil {
		return err
	}
	e.contentType = contentType
	return nil
}
func (e *Encoding) SetHeaders(headers map[string]*RefHeader) error {
	if err := e.Trix.RunHooks("headers", e.headers, headers); err != nil {
		return err
	}
	e.headers = headers
	return nil
}
func (e *Encoding) SetStyle(style string) error {
	if err := e.Trix.RunHooks("style", e.style, style); err != nil {
		return err
	}
	e.style = style
	return nil
}
func (e *Encoding) SetExplode(explode *bool) error {
	if err := e.Trix.RunHooks("explode", e.explode, explode); err != nil {
		return err
	}
	e.explode = explode
	return nil
}
func (e *Encoding) SetAllowReserved(allowReserved bool) error {
	if err := e.Trix.RunHooks("allowReserved", e.allowReserved, allowReserved); err != nil {
		return err
	}
	e.allowReserved = allowReserved
	return nil
}

// NewEncoding creates a new Encoding instance.
func NewEncoding(contentType string, headers map[string]*RefHeader, style string, explode *bool, allowReserved bool) *Encoding {
	return &Encoding{contentType: contentType, headers: headers, style: style, explode: explode, allowReserved: allowReserved}
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

// MarshalFields implements shared.MarshalFieldsProvider for export.
func (e *Encoding) MarshalFields() []shared.Field { return e.marshalFields() }

func (e *Encoding) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(e.marshalFields())
}

func (e *Encoding) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(e.marshalFields())
}

var _ yaml.Marshaler = (*Encoding)(nil)
