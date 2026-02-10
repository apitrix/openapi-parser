package openapi30

// SecuritySchemeRef represents a reference to a SecurityScheme or an inline SecurityScheme.
type SecuritySchemeRef struct {
	Node                     // embedded - provides VendorExtensions and Trix
	Ref      string          `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Value    *SecurityScheme `json:"-" yaml:"-"`
	Circular bool            `json:"-" yaml:"-"` // true if circular reference detected
}

// NewSecuritySchemeRef creates a new SecuritySchemeRef instance.
func NewSecuritySchemeRef(ref string) *SecuritySchemeRef {
	return &SecuritySchemeRef{Ref: ref}
}
