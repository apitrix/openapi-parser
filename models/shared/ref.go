package shared

// Ref represents a JSON Reference wrapper for any OpenAPI component type.
// Used by OpenAPI 2.0 and 3.0 where $ref is a simple pointer.
type Ref[T any] struct {
	VendorExtensions map[string]interface{} `json:"-" yaml:"-"`
	Trix             Trix                   `json:"-" yaml:"-"`
	Ref              string                 `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Value            *T                     `json:"-" yaml:"-"`
	Circular         bool                   `json:"-" yaml:"-"` // true if circular reference detected
}
