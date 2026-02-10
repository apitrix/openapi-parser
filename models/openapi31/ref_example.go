package openapi31

// ExampleRef represents a reference to an Example or an inline Example.
type ExampleRef struct {
	Node                 // embedded - provides VendorExtensions and Trix
	Ref         string   `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Summary     string   `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string   `json:"description,omitempty" yaml:"description,omitempty"`
	Value       *Example `json:"-" yaml:"-"`
	Circular    bool     `json:"-" yaml:"-"` // true if circular reference detected
}

// NewExampleRef creates a new ExampleRef instance.
func NewExampleRef(ref string) *ExampleRef {
	return &ExampleRef{Ref: ref}
}
