package openapi31

// HeaderRef represents a reference to a Header or an inline Header.
type HeaderRef struct {
	Node                // embedded - provides VendorExtensions and Trix
	Ref         string  `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Summary     string  `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string  `json:"description,omitempty" yaml:"description,omitempty"`
	Value       *Header `json:"-" yaml:"-"`
	Circular    bool    `json:"-" yaml:"-"` // true if circular reference detected
}

// NewHeaderRef creates a new HeaderRef instance.
func NewHeaderRef(ref string) *HeaderRef {
	return &HeaderRef{Ref: ref}
}
