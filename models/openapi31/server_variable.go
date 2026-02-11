package openapi31

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

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

func (v *ServerVariable) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "enum", Value: v.enum},
		{Key: "default", Value: v.defaultVal},
		{Key: "description", Value: v.description},
	}
	return shared.AppendExtensions(fields, v.VendorExtensions)
}

func (v *ServerVariable) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(v.marshalFields())
}

func (v *ServerVariable) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(v.marshalFields())
}

var _ yaml.Marshaler = (*ServerVariable)(nil)
