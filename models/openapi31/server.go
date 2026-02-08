package openapi31

// Server represents a server.
// https://spec.openapis.org/oas/v3.1.0#server-object
type Server struct {
	Node // embedded - provides NodeSource and Extensions

	URL         string                     `json:"url" yaml:"url"`
	Description string                     `json:"description,omitempty" yaml:"description,omitempty"`
	Variables   map[string]*ServerVariable `json:"variables,omitempty" yaml:"variables,omitempty"`
}

// ServerVariable represents a server variable for URL template substitution.
// https://spec.openapis.org/oas/v3.1.0#server-variable-object
type ServerVariable struct {
	Node // embedded - provides NodeSource and Extensions

	Enum        []string `json:"enum,omitempty" yaml:"enum,omitempty"`
	Default     string   `json:"default" yaml:"default"`
	Description string   `json:"description,omitempty" yaml:"description,omitempty"`
}
