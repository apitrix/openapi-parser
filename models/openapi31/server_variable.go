package openapi31

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// ServerVariable represents a server variable for URL template substitution.
// https://spec.openapis.org/oas/v3.1.0#server-variable-object
type ServerVariable struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	enum        []string
	defaultVal  string
	description string
}

func (v *ServerVariable) Enum() []string      { return v.enum }
func (v *ServerVariable) Default() string     { return v.defaultVal }
func (v *ServerVariable) Description() string { return v.description }

func (v *ServerVariable) SetEnum(enum []string) error {
	if err := v.Trix.RunHooks("enum", v.enum, enum); err != nil {
		return err
	}
	v.enum = enum
	return nil
}
func (v *ServerVariable) SetDefault(defaultVal string) error {
	if err := v.Trix.RunHooks("default", v.defaultVal, defaultVal); err != nil {
		return err
	}
	v.defaultVal = defaultVal
	return nil
}
func (v *ServerVariable) SetDescription(description string) error {
	if err := v.Trix.RunHooks("description", v.description, description); err != nil {
		return err
	}
	v.description = description
	return nil
}

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

// MarshalFields implements shared.MarshalFieldsProvider for export.
func (v *ServerVariable) MarshalFields() []shared.Field { return v.marshalFields() }

func (v *ServerVariable) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(v.marshalFields())
}

func (v *ServerVariable) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(v.marshalFields())
}

var _ yaml.Marshaler = (*ServerVariable)(nil)
