package openapi31

// SchemaRef represents a reference to a Schema or an inline Schema.
type SchemaRef struct {
	Node                // embedded - provides NodeSource and Extensions
	Ref         string  `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Summary     string  `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string  `json:"description,omitempty" yaml:"description,omitempty"`
	Value       *Schema `json:"-" yaml:"-"`
}

// ResponseRef represents a reference to a Response or an inline Response.
type ResponseRef struct {
	Node                  // embedded - provides NodeSource and Extensions
	Ref         string    `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Summary     string    `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string    `json:"description,omitempty" yaml:"description,omitempty"`
	Value       *Response `json:"-" yaml:"-"`
}

// ParameterRef represents a reference to a Parameter or an inline Parameter.
type ParameterRef struct {
	Node                   // embedded - provides NodeSource and Extensions
	Ref         string     `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Summary     string     `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string     `json:"description,omitempty" yaml:"description,omitempty"`
	Value       *Parameter `json:"-" yaml:"-"`
}

// ExampleRef represents a reference to an Example or an inline Example.
type ExampleRef struct {
	Node                 // embedded - provides NodeSource and Extensions
	Ref         string   `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Summary     string   `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string   `json:"description,omitempty" yaml:"description,omitempty"`
	Value       *Example `json:"-" yaml:"-"`
}

// RequestBodyRef represents a reference to a RequestBody or an inline RequestBody.
type RequestBodyRef struct {
	Node                     // embedded - provides NodeSource and Extensions
	Ref         string       `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Summary     string       `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string       `json:"description,omitempty" yaml:"description,omitempty"`
	Value       *RequestBody `json:"-" yaml:"-"`
}

// HeaderRef represents a reference to a Header or an inline Header.
type HeaderRef struct {
	Node                // embedded - provides NodeSource and Extensions
	Ref         string  `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Summary     string  `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string  `json:"description,omitempty" yaml:"description,omitempty"`
	Value       *Header `json:"-" yaml:"-"`
}

// SecuritySchemeRef represents a reference to a SecurityScheme or an inline SecurityScheme.
type SecuritySchemeRef struct {
	Node                        // embedded - provides NodeSource and Extensions
	Ref         string          `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Summary     string          `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string          `json:"description,omitempty" yaml:"description,omitempty"`
	Value       *SecurityScheme `json:"-" yaml:"-"`
}

// LinkRef represents a reference to a Link or an inline Link.
type LinkRef struct {
	Node               // embedded - provides NodeSource and Extensions
	Ref         string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Summary     string `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Value       *Link  `json:"-" yaml:"-"`
}

// CallbackRef represents a reference to a Callback or an inline Callback.
type CallbackRef struct {
	Node                  // embedded - provides NodeSource and Extensions
	Ref         string    `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Summary     string    `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string    `json:"description,omitempty" yaml:"description,omitempty"`
	Value       *Callback `json:"-" yaml:"-"`
}

// PathItemRef represents a reference to a PathItem or an inline PathItem.
type PathItemRef struct {
	Node                  // embedded - provides NodeSource and Extensions
	Ref         string    `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Summary     string    `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string    `json:"description,omitempty" yaml:"description,omitempty"`
	Value       *PathItem `json:"-" yaml:"-"`
}
