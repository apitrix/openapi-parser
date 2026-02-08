package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// ParseXML parses the Schema.XML field.
// Complex property: nested XML object
func (p *schemaParser) ParseXML(parent *yaml.Node, c *ParseContext) (*openapi31models.XML, error) {
	node := nodeGetValue(parent, "xml")
	if node == nil {
		return nil, nil
	}

	pctx := c.Push("xml")

	if !nodeIsMapping(node) {
		return nil, pctx.errorAt(node, "xml must be an object")
	}

	xml := &openapi31models.XML{}

	// All XML properties are simple
	xml.Name = nodeGetString(node, "name")
	xml.Namespace = nodeGetString(node, "namespace")
	xml.Prefix = nodeGetString(node, "prefix")
	xml.Attribute = nodeGetBool(node, "attribute")
	xml.Wrapped = nodeGetBool(node, "wrapped")

	xml.VendorExtensions = parseNodeExtensions(node)
	xml.Trix.Source = pctx.nodeSource(node)

	// Detect unknown fields
	pctx.detectUnknown(node, xmlKnownFieldsSet)

	return xml, nil
}
