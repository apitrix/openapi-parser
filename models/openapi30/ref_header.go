package openapi30

// HeaderRef represents a reference to a Header or an inline Header.
type HeaderRef struct {
	Node             // embedded - provides VendorExtensions and Trix
	Ref      string  `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Value    *Header `json:"-" yaml:"-"`
	Circular bool    `json:"-" yaml:"-"` // true if circular reference detected
}

// NewHeaderRef creates a new HeaderRef instance.
func NewHeaderRef(ref string) *HeaderRef {
	return &HeaderRef{Ref: ref}
}
