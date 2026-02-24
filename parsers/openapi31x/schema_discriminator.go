package openapi31x

import (
	openapi31models "github.com/apitrix/openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// ParseDiscriminator parses the Schema.Discriminator field.
// Complex property: nested Discriminator object
func (p *schemaParser) ParseDiscriminator(parent *yaml.Node, c *ParseContext) (*openapi31models.Discriminator, error) {
	node := nodeGetValue(parent, "discriminator")
	if node == nil {
		return nil, nil
	}

	pctx := c.Push("discriminator")

	if !nodeIsMapping(node) {
		return nil, pctx.errorAt(node, "discriminator must be an object")
	}

	// Create via constructor
	disc := openapi31models.NewDiscriminator(
		nodeGetString(node, "propertyName"),
		nodeGetStringMap(node, "mapping"),
	)

	disc.VendorExtensions = parseNodeExtensions(node)
	disc.Trix.Source = pctx.nodeSource(node)

	// Set OpenAPI 3.2 field via setter
	_ = disc.SetDefaultMapping(nodeGetString(node, "defaultMapping"))

	// Detect unknown fields
	disc.Trix.Errors = append(disc.Trix.Errors, unknownFieldParseErrors(pctx.detectUnknown(node, discriminatorKnownFieldsSet))...)

	return disc, nil
}
