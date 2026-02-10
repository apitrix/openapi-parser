package openapi30

// Encoding defines encoding for a single schema property.
// https://spec.openapis.org/oas/v3.0.3#encoding-object
type Encoding struct {
	Node // embedded - provides VendorExtensions and Trix

	contentType   string
	headers       map[string]*HeaderRef
	style         string
	explode       *bool
	allowReserved bool
}

func (e *Encoding) ContentType() string            { return e.contentType }
func (e *Encoding) Headers() map[string]*HeaderRef { return e.headers }
func (e *Encoding) Style() string                  { return e.style }
func (e *Encoding) Explode() *bool                 { return e.explode }
func (e *Encoding) AllowReserved() bool            { return e.allowReserved }

// NewEncoding creates a new Encoding instance.
func NewEncoding(contentType string, headers map[string]*HeaderRef, style string, explode *bool, allowReserved bool) *Encoding {
	return &Encoding{contentType: contentType, headers: headers, style: style, explode: explode, allowReserved: allowReserved}
}
