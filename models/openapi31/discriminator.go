package openapi31

// Discriminator is used for polymorphism support.
// https://spec.openapis.org/oas/v3.1.0#discriminator-object
type Discriminator struct {
	Node // embedded - provides VendorExtensions and Trix

	propertyName string
	mapping      map[string]string
}

func (d *Discriminator) PropertyName() string       { return d.propertyName }
func (d *Discriminator) Mapping() map[string]string { return d.mapping }

// NewDiscriminator creates a new Discriminator instance.
func NewDiscriminator(propertyName string, mapping map[string]string) *Discriminator {
	return &Discriminator{propertyName: propertyName, mapping: mapping}
}
