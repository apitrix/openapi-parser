package openapi31

// Responses is a container for expected responses of an operation.
// https://spec.openapis.org/oas/v3.1.0#responses-object
type Responses struct {
	Node // embedded - provides NodeSource and Extensions

	Default *ResponseRef            `json:"default,omitempty" yaml:"default,omitempty"`
	Codes   map[string]*ResponseRef `json:"-" yaml:"-"` // HTTP status codes (e.g., "200", "404", "5XX")
}

// Response describes a single response from an API operation.
// https://spec.openapis.org/oas/v3.1.0#response-object
type Response struct {
	Node // embedded - provides NodeSource and Extensions

	Description string                `json:"description" yaml:"description"`
	Headers     map[string]*HeaderRef `json:"headers,omitempty" yaml:"headers,omitempty"`
	Content     map[string]*MediaType `json:"content,omitempty" yaml:"content,omitempty"`
	Links       map[string]*LinkRef   `json:"links,omitempty" yaml:"links,omitempty"`
}
