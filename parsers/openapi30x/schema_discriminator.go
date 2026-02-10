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

	// Collect values and create via constructor
	propertyName := nodeGetString(node, "propertyName")
	mapping := nodeGetStringMap(node, "mapping")

	disc := openapi30models.NewDiscriminator(propertyName, mapping)

	disc.VendorExtensions = parseNodeExtensions(node)
	disc.Trix.Source = pctx.nodeSource(node)

	// Detect unknown fields
	disc.Trix.Errors = append(disc.Trix.Errors, unknownFieldParseErrors(pctx.detectUnknown(node, discriminatorKnownFieldsSet))...)

	return disc, nil
}
