package openapi30x

import (
	openapi30models "github.com/apitrix/openapi-parser/models/openapi30"

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

	// Collect values and create via constructor
	name := nodeGetString(node, "name")
	namespace := nodeGetString(node, "namespace")
	prefix := nodeGetString(node, "prefix")
	attribute := nodeGetBool(node, "attribute")
	wrapped := nodeGetBool(node, "wrapped")

	xml := openapi30models.NewXML(name, namespace, prefix, attribute, wrapped)

	xml.VendorExtensions = parseNodeExtensions(node)
	xml.Trix.Source = pctx.nodeSource(node)

	// Detect unknown fields
	xml.Trix.Errors = append(xml.Trix.Errors, unknownFieldParseErrors(pctx.detectUnknown(node, xmlKnownFieldsSet))...)

	return xml, nil
}
