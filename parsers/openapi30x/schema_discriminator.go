package openapi30x

import (
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// ParseDiscriminator parses the Schema.Discriminator field.
// Complex property: nested Discriminator object
func (p *schemaParser) ParseDiscriminator(parent *yaml.Node, c *ParseContext) (*openapi30models.Discriminator, error) {
	node := nodeGetValue(parent, "discriminator")
	if node == nil {
		return nil, nil
	}

	pctx := c.Push("discriminator")

	if !nodeIsMapping(node) {
		return nil, pctx.errorAt(node, "discriminator must be an object")
	}

	disc := &openapi30models.Discriminator{}

	// All discriminator properties are simple
	disc.PropertyName = nodeGetString(node, "propertyName")
	disc.Mapping = nodeGetStringMap(node, "mapping")

	disc.VendorExtensions = parseNodeExtensions(node)
	disc.Trix.Source = pctx.nodeSource(node)

	// Detect unknown fields
	pctx.detectUnknown(node, discriminatorKnownFieldsSet)

	return disc, nil
}
