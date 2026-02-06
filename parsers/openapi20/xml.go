package openapi20

import (
	openapi20models "openapi-parser/models/openapi20"

	"gopkg.in/yaml.v3"
)

// parseXML parses an XML object from a yaml.Node.
func parseXML(node *yaml.Node, ctx *ParseContext) (*openapi20models.XML, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "xml must be an object")
	}

	xml := &openapi20models.XML{}

	// Simple properties - inline
	xml.Name = nodeGetString(node, "name")
	xml.Namespace = nodeGetString(node, "namespace")
	xml.Prefix = nodeGetString(node, "prefix")
	xml.Attribute = nodeGetBool(node, "attribute")
	xml.Wrapped = nodeGetBool(node, "wrapped")

	xml.Extensions = parseNodeExtensions(node)
	xml.NodeSource = ctx.nodeSource(node)

	// Detect unknown fields
	ctx.detectUnknown(node, xmlKnownFields)

	return xml, nil
}
