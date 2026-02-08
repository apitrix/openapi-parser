package openapi30x

import (
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// ParseXML parses the Schema.XML field.
// Complex property: nested XML object
func (p *schemaParser) ParseXML(parent *yaml.Node, c *ParseContext) (*openapi30models.XML, error) {
	node := nodeGetValue(parent, "xml")
	if node == nil {
		return nil, nil
	}

	pctx := c.Push("xml")

	if !nodeIsMapping(node) {
		return nil, pctx.errorAt(node, "xml must be an object")
	}

	xml := &openapi30models.XML{}

	// All XML properties are simple
	xml.Name = nodeGetString(node, "name")
	xml.Namespace = nodeGetString(node, "namespace")
	xml.Prefix = nodeGetString(node, "prefix")
	xml.Attribute = nodeGetBool(node, "attribute")
	xml.Wrapped = nodeGetBool(node, "wrapped")

	xml.VendorExtensions = parseNodeExtensions(node)
	xml.Trix.Source = pctx.nodeSource(node)

	// Detect unknown fields
	xml.Trix.Errors = append(xml.Trix.Errors, unknownFieldParseErrors(pctx.detectUnknown(node, xmlKnownFieldsSet))...)

	return xml, nil
}
