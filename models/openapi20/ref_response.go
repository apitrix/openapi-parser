package openapi20

// ResponseRef represents a reference to a Response or an inline Response.
type ResponseRef struct {
	Node               // embedded - provides VendorExtensions and Trix
	Ref      string    `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Value    *Response `json:"-" yaml:"-"`
	Circular bool      `json:"-" yaml:"-"` // true if circular reference detected
}

// NewResponseRef creates a new ResponseRef instance.
func NewResponseRef(ref string) *ResponseRef {
	return &ResponseRef{Ref: ref}
}
