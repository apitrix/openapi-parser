package openapi20

// XML provides additional metadata for XML serialization.
// https://swagger.io/specification/v2/#xml-object
type XML struct {
	Node // embedded - provides VendorExtensions and Trix

	Name      string `json:"name,omitempty" yaml:"name,omitempty"`
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	Prefix    string `json:"prefix,omitempty" yaml:"prefix,omitempty"`
	Attribute bool   `json:"attribute,omitempty" yaml:"attribute,omitempty"`
	Wrapped   bool   `json:"wrapped,omitempty" yaml:"wrapped,omitempty"`
}

// NewXML creates a new XML instance.
func NewXML() *XML {
	return &XML{}
}
