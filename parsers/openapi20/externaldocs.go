package openapi20

import (
	openapi20models "openapi-parser/models/openapi20"

	"gopkg.in/yaml.v3"
)

// parseExternalDocs parses an ExternalDocs object from a yaml.Node.
func parseExternalDocs(node *yaml.Node, ctx *ParseContext) (*openapi20models.ExternalDocs, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "externalDocs must be an object")
	}

	ed := &openapi20models.ExternalDocs{}

	// Simple properties - inline
	ed.Description = nodeGetString(node, "description")
	ed.URL = nodeGetString(node, "url")

	ed.VendorExtensions = parseNodeExtensions(node)
	ed.Trix.Source = ctx.nodeSource(node)

	// Detect unknown fields
	ctx.detectUnknown(node, externalDocsKnownFieldsSet)

	return ed, nil
}
