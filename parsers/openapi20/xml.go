package openapi20

import (
	openapi20models "github.com/apitrix/openapi-parser/models/openapi20"

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

	xml := openapi20models.NewXML(
		nodeGetString(node, "name"),
		nodeGetString(node, "namespace"),
		nodeGetString(node, "prefix"),
		nodeGetBool(node, "attribute"),
		nodeGetBool(node, "wrapped"),
	)

	xml.VendorExtensions = parseNodeExtensions(node)
	xml.Trix.Source = ctx.nodeSource(node)

	// Detect unknown fields
	xml.Trix.Errors = append(xml.Trix.Errors, unknownFieldParseErrors(ctx.detectUnknown(node, xmlKnownFieldsSet))...)

	return xml, nil
}
