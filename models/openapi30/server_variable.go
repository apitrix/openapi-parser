package openapi30

// ServerVariable represents a server variable for URL template substitution.
// https://spec.openapis.org/oas/v3.0.3#server-variable-object
type ServerVariable struct {
	Node // embedded - provides VendorExtensions and Trix

	enum        []string
	defaultVal  string
	description string
}

func (sv *ServerVariable) Enum() []string      { return sv.enum }
func (sv *ServerVariable) Default() string     { return sv.defaultVal }
func (sv *ServerVariable) Description() string { return sv.description }

// NewServerVariable creates a new ServerVariable instance.
func NewServerVariable(defaultValue, description string, enum []string) *ServerVariable {
	return &ServerVariable{defaultVal: defaultValue, description: description, enum: enum}
}
