package openapi31

// Example represents an example of a media type.
// https://spec.openapis.org/oas/v3.1.0#example-object
type Example struct {
	Node // embedded - provides VendorExtensions and Trix

	Summary       string      `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description   string      `json:"description,omitempty" yaml:"description,omitempty"`
	Value         interface{} `json:"value,omitempty" yaml:"value,omitempty"`
	ExternalValue string      `json:"externalValue,omitempty" yaml:"externalValue,omitempty"`
}

// NewExample creates a new Example instance.
func NewExample() *Example {
	return &Example{}
}
