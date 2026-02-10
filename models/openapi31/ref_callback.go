package openapi31

// CallbackRef represents a reference to a Callback or an inline Callback.
type CallbackRef struct {
	Node                  // embedded - provides VendorExtensions and Trix
	Ref         string    `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Summary     string    `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string    `json:"description,omitempty" yaml:"description,omitempty"`
	Value       *Callback `json:"-" yaml:"-"`
	Circular    bool      `json:"-" yaml:"-"` // true if circular reference detected
}

// NewCallbackRef creates a new CallbackRef instance.
func NewCallbackRef(ref string) *CallbackRef {
	return &CallbackRef{Ref: ref}
}
