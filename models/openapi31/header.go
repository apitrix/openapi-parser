package openapi31

// Header represents a Header Object.
// https://spec.openapis.org/oas/v3.1.0#header-object
type Header struct {
	Node // embedded - provides VendorExtensions and Trix

	Description     string                 `json:"description,omitempty" yaml:"description,omitempty"`
	Required        bool                   `json:"required,omitempty" yaml:"required,omitempty"`
	Deprecated      bool                   `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	AllowEmptyValue bool                   `json:"allowEmptyValue,omitempty" yaml:"allowEmptyValue,omitempty"`
	Style           string                 `json:"style,omitempty" yaml:"style,omitempty"`
	Explode         *bool                  `json:"explode,omitempty" yaml:"explode,omitempty"`
	AllowReserved   bool                   `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"`
	Schema          *SchemaRef             `json:"schema,omitempty" yaml:"schema,omitempty"`
	Example         interface{}            `json:"example,omitempty" yaml:"example,omitempty"`
	Examples        map[string]*ExampleRef `json:"examples,omitempty" yaml:"examples,omitempty"`
	Content         map[string]*MediaType  `json:"content,omitempty" yaml:"content,omitempty"`
}

// NewHeader creates a new Header instance.
func NewHeader() *Header {
	return &Header{}
}
