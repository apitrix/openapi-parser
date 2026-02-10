package openapi31

// RequestBodyRef represents a reference to a RequestBody or an inline RequestBody.
type RequestBodyRef struct {
	Node                     // embedded - provides VendorExtensions and Trix
	Ref         string       `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Summary     string       `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string       `json:"description,omitempty" yaml:"description,omitempty"`
	Value       *RequestBody `json:"-" yaml:"-"`
	Circular    bool         `json:"-" yaml:"-"` // true if circular reference detected
}

// NewRequestBodyRef creates a new RequestBodyRef instance.
func NewRequestBodyRef(ref string) *RequestBodyRef {
	return &RequestBodyRef{Ref: ref}
}
