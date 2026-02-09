package shared

// Ref31 represents a JSON Reference wrapper for OpenAPI 3.1 component types.
// In 3.1, $ref objects can carry summary and description alongside the pointer.
type Ref31[T any] struct {
	VendorExtensions map[string]interface{} `json:"-" yaml:"-"`
	Trix             Trix                   `json:"-" yaml:"-"`
	Ref              string                 `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Summary          string                 `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description      string                 `json:"description,omitempty" yaml:"description,omitempty"`
	Value            *T                     `json:"-" yaml:"-"`
	Circular         bool                   `json:"-" yaml:"-"` // true if circular reference detected
}
