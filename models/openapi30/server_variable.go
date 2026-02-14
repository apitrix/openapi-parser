package openapi30

import (
	"openapi-parser/models/shared"

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

func (sv *ServerVariable) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(sv.marshalFields())
}

func (sv *ServerVariable) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(sv.marshalFields())
}

var _ yaml.Marshaler = (*ServerVariable)(nil)
