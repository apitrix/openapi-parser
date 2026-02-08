package openapi30

// Parameter describes a single operation parameter.
// https://spec.openapis.org/oas/v3.0.3#parameter-object
type Parameter struct {
	Node // embedded - provides VendorExtensions and Trix

	Name            string                 `json:"name" yaml:"name"`
	In              string                 `json:"in" yaml:"in"`
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

// Header represents a Header Object.
// https://spec.openapis.org/oas/v3.0.3#header-object
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
