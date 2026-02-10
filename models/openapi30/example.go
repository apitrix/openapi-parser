package openapi30

// Example represents an example of a media type.
// https://spec.openapis.org/oas/v3.0.3#example-object
type Example struct {
	Node // embedded - provides VendorExtensions and Trix

	summary       string
	description   string
	value         interface{}
	externalValue string
}

func (e *Example) Summary() string       { return e.summary }
func (e *Example) Description() string   { return e.description }
func (e *Example) Value() interface{}    { return e.value }
func (e *Example) ExternalValue() string { return e.externalValue }

// NewExample creates a new Example instance.
func NewExample(summary, description string, value interface{}, externalValue string) *Example {
	return &Example{summary: summary, description: description, value: value, externalValue: externalValue}
}
