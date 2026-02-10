package openapi31

// ServerVariable represents a server variable for URL template substitution.
// https://spec.openapis.org/oas/v3.1.0#server-variable-object
type ServerVariable struct {
	Node // embedded - provides VendorExtensions and Trix

	enum        []string
	defaultVal  string
	description string
}

func (v *ServerVariable) Enum() []string      { return v.enum }
func (v *ServerVariable) Default() string     { return v.defaultVal }
func (v *ServerVariable) Description() string { return v.description }

// NewServerVariable creates a new ServerVariable instance.
func NewServerVariable(enum []string, defaultValue, description string) *ServerVariable {
	return &ServerVariable{enum: enum, defaultVal: defaultValue, description: description}
}
