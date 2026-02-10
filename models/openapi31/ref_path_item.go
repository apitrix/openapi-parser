package openapi31

// PathItemRef represents a reference to a PathItem or an inline PathItem.
type PathItemRef struct {
	Node                  // embedded - provides VendorExtensions and Trix
	Ref         string    `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Summary     string    `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string    `json:"description,omitempty" yaml:"description,omitempty"`
	Value       *PathItem `json:"-" yaml:"-"`
	Circular    bool      `json:"-" yaml:"-"` // true if circular reference detected
}

// NewPathItemRef creates a new PathItemRef instance.
func NewPathItemRef(ref string) *PathItemRef {
	return &PathItemRef{Ref: ref}
}
