package openapi31

// SchemaRef represents a reference to a Schema or an inline Schema.
type SchemaRef struct {
	Node                // embedded - provides VendorExtensions and Trix
	Ref         string  `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Summary     string  `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string  `json:"description,omitempty" yaml:"description,omitempty"`
	Value       *Schema `json:"-" yaml:"-"`
	Circular    bool    `json:"-" yaml:"-"` // true if circular reference detected
}

// NewSchemaRef creates a new SchemaRef instance.
func NewSchemaRef(ref string) *SchemaRef {
	return &SchemaRef{Ref: ref}
}
