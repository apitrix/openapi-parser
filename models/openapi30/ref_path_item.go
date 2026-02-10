package openapi30

// PathItemRef represents a reference to a PathItem or an inline PathItem.
type PathItemRef struct {
	Node               // embedded - provides VendorExtensions and Trix
	Ref      string    `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Value    *PathItem `json:"-" yaml:"-"`
	Circular bool      `json:"-" yaml:"-"` // true if circular reference detected
}

// NewPathItemRef creates a new PathItemRef instance.
func NewPathItemRef(ref string) *PathItemRef {
	return &PathItemRef{Ref: ref}
}
