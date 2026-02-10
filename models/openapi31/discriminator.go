package openapi31

// Discriminator is used for polymorphism support.
// https://spec.openapis.org/oas/v3.1.0#discriminator-object
type Discriminator struct {
	Node // embedded - provides VendorExtensions and Trix

	PropertyName string            `json:"propertyName" yaml:"propertyName"`
	Mapping      map[string]string `json:"mapping,omitempty" yaml:"mapping,omitempty"`
}

// NewDiscriminator creates a new Discriminator instance.
func NewDiscriminator(propertyName string, mapping map[string]string) *Discriminator {
	return &Discriminator{PropertyName: propertyName, Mapping: mapping}
}
