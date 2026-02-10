package openapi30

// Encoding defines encoding for a single schema property.
// https://spec.openapis.org/oas/v3.0.3#encoding-object
type Encoding struct {
	Node // embedded - provides VendorExtensions and Trix

	ContentType   string                `json:"contentType,omitempty" yaml:"contentType,omitempty"`
	Headers       map[string]*HeaderRef `json:"headers,omitempty" yaml:"headers,omitempty"`
	Style         string                `json:"style,omitempty" yaml:"style,omitempty"`
	Explode       *bool                 `json:"explode,omitempty" yaml:"explode,omitempty"`
	AllowReserved bool                  `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"`
}

// NewEncoding creates a new Encoding instance.
func NewEncoding() *Encoding {
	return &Encoding{}
}
