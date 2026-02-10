package openapi31

// SecuritySchemeRef represents a reference to a SecurityScheme or an inline SecurityScheme.
type SecuritySchemeRef struct {
	Node                        // embedded - provides VendorExtensions and Trix
	Ref         string          `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Summary     string          `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string          `json:"description,omitempty" yaml:"description,omitempty"`
	Value       *SecurityScheme `json:"-" yaml:"-"`
	Circular    bool            `json:"-" yaml:"-"` // true if circular reference detected
}

// NewSecuritySchemeRef creates a new SecuritySchemeRef instance.
func NewSecuritySchemeRef(ref string) *SecuritySchemeRef {
	return &SecuritySchemeRef{Ref: ref}
}
