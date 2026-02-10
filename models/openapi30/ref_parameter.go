package openapi30

// ParameterRef represents a reference to a Parameter or an inline Parameter.
type ParameterRef struct {
	Node                // embedded - provides VendorExtensions and Trix
	Ref      string     `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Value    *Parameter `json:"-" yaml:"-"`
	Circular bool       `json:"-" yaml:"-"` // true if circular reference detected
}

// NewParameterRef creates a new ParameterRef instance.
func NewParameterRef(ref string) *ParameterRef {
	return &ParameterRef{Ref: ref}
}
