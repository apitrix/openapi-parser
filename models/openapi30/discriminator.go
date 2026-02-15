package openapi30

import (
	"github.com/apitrix/openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// Discriminator is used for polymorphism support.
// https://spec.openapis.org/oas/v3.0.3#discriminator-object
type Discriminator struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	propertyName string
	mapping      map[string]string
}

func (d *Discriminator) PropertyName() string       { return d.propertyName }
func (d *Discriminator) Mapping() map[string]string { return d.mapping }

func (d *Discriminator) SetPropertyName(propertyName string) error {
	if err := d.Trix.RunHooks("propertyName", d.propertyName, propertyName); err != nil {
		return err
	}
	d.propertyName = propertyName
	return nil
}
func (d *Discriminator) SetMapping(mapping map[string]string) error {
	if err := d.Trix.RunHooks("mapping", d.mapping, mapping); err != nil {
		return err
	}
	d.mapping = mapping
	return nil
}

// NewDiscriminator creates a new Discriminator instance.
func NewDiscriminator(propertyName string, mapping map[string]string) *Discriminator {
	return &Discriminator{propertyName: propertyName, mapping: mapping}
}

func (d *Discriminator) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "propertyName", Value: d.propertyName},
		{Key: "mapping", Value: d.mapping},
	}
	return shared.AppendExtensions(fields, d.VendorExtensions)
}

// MarshalFields implements shared.MarshalFieldsProvider for export.
func (d *Discriminator) MarshalFields() []shared.Field { return d.marshalFields() }

func (d *Discriminator) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(d.marshalFields())
}

func (d *Discriminator) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(d.marshalFields())
}

var _ yaml.Marshaler = (*Discriminator)(nil)
