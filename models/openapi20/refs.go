package openapi20

// SchemaRef represents a reference to a Schema or an inline Schema.
type SchemaRef struct {
	Node  // embedded - provides NodeSource and Extensions
	Ref   string  `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Value *Schema `json:"-" yaml:"-"`
}

// ResponseRef represents a reference to a Response or an inline Response.
type ResponseRef struct {
	Node  // embedded - provides NodeSource and Extensions
	Ref   string    `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Value *Response `json:"-" yaml:"-"`
}

// ParameterRef represents a reference to a Parameter or an inline Parameter.
type ParameterRef struct {
	Node  // embedded - provides NodeSource and Extensions
	Ref   string     `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Value *Parameter `json:"-" yaml:"-"`
}
