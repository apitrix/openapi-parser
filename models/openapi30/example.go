package openapi30

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// Example represents an example of a media type.
// https://spec.openapis.org/oas/v3.0.3#example-object
type Example struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	summary       string
	description   string
	value         interface{}
	externalValue string
}

func (e *Example) Summary() string       { return e.summary }
func (e *Example) Description() string   { return e.description }
func (e *Example) Value() interface{}    { return e.value }
func (e *Example) ExternalValue() string { return e.externalValue }

func (e *Example) SetSummary(summary string) error {
	if err := e.Trix.RunHooks("summary", e.summary, summary); err != nil {
		return err
	}
	e.summary = summary
	return nil
}
func (e *Example) SetDescription(description string) error {
	if err := e.Trix.RunHooks("description", e.description, description); err != nil {
		return err
	}
	e.description = description
	return nil
}
func (e *Example) SetValue(value interface{}) error {
	if err := e.Trix.RunHooks("value", e.value, value); err != nil {
		return err
	}
	e.value = value
	return nil
}
func (e *Example) SetExternalValue(externalValue string) error {
	if err := e.Trix.RunHooks("externalValue", e.externalValue, externalValue); err != nil {
		return err
	}
	e.externalValue = externalValue
	return nil
}

// NewExample creates a new Example instance.
func NewExample(summary, description string, value interface{}, externalValue string) *Example {
	return &Example{summary: summary, description: description, value: value, externalValue: externalValue}
}

func (e *Example) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "summary", Value: e.summary},
		{Key: "description", Value: e.description},
		{Key: "value", Value: e.value},
		{Key: "externalValue", Value: e.externalValue},
	}
	return shared.AppendExtensions(fields, e.VendorExtensions)
}

func (e *Example) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(e.marshalFields())
}

func (e *Example) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(e.marshalFields())
}

var _ yaml.Marshaler = (*Example)(nil)
