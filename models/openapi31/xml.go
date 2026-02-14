package openapi31

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// XML provides additional metadata for XML serialization.
// https://spec.openapis.org/oas/v3.1.0#xml-object
type XML struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	name      string
	namespace string
	prefix    string
	attribute bool
	wrapped   bool
}

func (x *XML) Name() string      { return x.name }
func (x *XML) Namespace() string { return x.namespace }
func (x *XML) Prefix() string    { return x.prefix }
func (x *XML) Attribute() bool   { return x.attribute }
func (x *XML) Wrapped() bool     { return x.wrapped }

func (x *XML) SetName(name string) error {
	if err := x.Trix.RunHooks("name", x.name, name); err != nil {
		return err
	}
	x.name = name
	return nil
}
func (x *XML) SetNamespace(namespace string) error {
	if err := x.Trix.RunHooks("namespace", x.namespace, namespace); err != nil {
		return err
	}
	x.namespace = namespace
	return nil
}
func (x *XML) SetPrefix(prefix string) error {
	if err := x.Trix.RunHooks("prefix", x.prefix, prefix); err != nil {
		return err
	}
	x.prefix = prefix
	return nil
}
func (x *XML) SetAttribute(attribute bool) error {
	if err := x.Trix.RunHooks("attribute", x.attribute, attribute); err != nil {
		return err
	}
	x.attribute = attribute
	return nil
}
func (x *XML) SetWrapped(wrapped bool) error {
	if err := x.Trix.RunHooks("wrapped", x.wrapped, wrapped); err != nil {
		return err
	}
	x.wrapped = wrapped
	return nil
}

// NewXML creates a new XML instance.
func NewXML(name, namespace, prefix string, attribute, wrapped bool) *XML {
	return &XML{name: name, namespace: namespace, prefix: prefix, attribute: attribute, wrapped: wrapped}
}

func (x *XML) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "name", Value: x.name},
		{Key: "namespace", Value: x.namespace},
		{Key: "prefix", Value: x.prefix},
		{Key: "attribute", Value: x.attribute},
		{Key: "wrapped", Value: x.wrapped},
	}
	return shared.AppendExtensions(fields, x.VendorExtensions)
}

func (x *XML) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(x.marshalFields())
}

func (x *XML) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(x.marshalFields())
}

var _ yaml.Marshaler = (*XML)(nil)
