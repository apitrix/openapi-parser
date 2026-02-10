package openapi31

// ServerVariable represents a server variable for URL template substitution.
// https://spec.openapis.org/oas/v3.1.0#server-variable-object
type ServerVariable struct {
	Node // embedded - provides VendorExtensions and Trix

	Enum        []string `json:"enum,omitempty" yaml:"enum,omitempty"`
	Default     string   `json:"default" yaml:"default"`
	Description string   `json:"description,omitempty" yaml:"description,omitempty"`
}

// NewServerVariable creates a new ServerVariable instance.
func NewServerVariable(defaultValue string) *ServerVariable {
	return &ServerVariable{Default: defaultValue}
}
