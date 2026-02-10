package openapi30

// RequestBodyRef represents a reference to a RequestBody or an inline RequestBody.
type RequestBodyRef struct {
	Node                  // embedded - provides VendorExtensions and Trix
	Ref      string       `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Value    *RequestBody `json:"-" yaml:"-"`
	Circular bool         `json:"-" yaml:"-"` // true if circular reference detected
}

// NewRequestBodyRef creates a new RequestBodyRef instance.
func NewRequestBodyRef(ref string) *RequestBodyRef {
	return &RequestBodyRef{Ref: ref}
}
