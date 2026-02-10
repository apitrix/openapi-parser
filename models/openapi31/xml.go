package openapi31

// XML provides additional metadata for XML serialization.
// https://spec.openapis.org/oas/v3.1.0#xml-object
type XML struct {
	Node // embedded - provides VendorExtensions and Trix

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

// NewXML creates a new XML instance.
func NewXML(name, namespace, prefix string, attribute, wrapped bool) *XML {
	return &XML{name: name, namespace: namespace, prefix: prefix, attribute: attribute, wrapped: wrapped}
}
