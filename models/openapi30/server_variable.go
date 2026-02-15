package openapi30

import (
	"github.com/apitrix/openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// ServerVariable represents a server variable for URL template substitution.
// https://spec.openapis.org/oas/v3.0.3#server-variable-object
type ServerVariable struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	enum        []string
	defaultVal  string
	description string
}

func (sv *ServerVariable) Enum() []string      { return sv.enum }
func (sv *ServerVariable) Default() string     { return sv.defaultVal }
func (sv *ServerVariable) Description() string { return sv.description }

func (sv *ServerVariable) SetEnum(enum []string) error {
	if err := sv.Trix.RunHooks("enum", sv.enum, enum); err != nil {
		return err
	}
	sv.enum = enum
	return nil
}
func (sv *ServerVariable) SetDefault(defaultVal string) error {
	if err := sv.Trix.RunHooks("default", sv.defaultVal, defaultVal); err != nil {
		return err
	}
	sv.defaultVal = defaultVal
	return nil
}
func (sv *ServerVariable) SetDescription(description string) error {
	if err := sv.Trix.RunHooks("description", sv.description, description); err != nil {
		return err
	}
	sv.description = description
	return nil
}

// NewServerVariable creates a new ServerVariable instance.
func NewServerVariable(defaultValue, description string, enum []string) *ServerVariable {
	return &ServerVariable{defaultVal: defaultValue, description: description, enum: enum}
}

func (sv *ServerVariable) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "enum", Value: sv.enum},
		{Key: "default", Value: sv.defaultVal},
		{Key: "description", Value: sv.description},
	}
	return shared.AppendExtensions(fields, sv.VendorExtensions)
}

// MarshalFields implements shared.MarshalFieldsProvider for export.
func (sv *ServerVariable) MarshalFields() []shared.Field { return sv.marshalFields() }

func (sv *ServerVariable) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(sv.marshalFields())
}

func (sv *ServerVariable) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(sv.marshalFields())
}

var _ yaml.Marshaler = (*ServerVariable)(nil)
