package openapi30x

import (
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

type externalDocsParser struct{}

// defaultExternalDocsParser is the singleton instance used by parsing functions.
var defaultExternalDocsParser = &externalDocsParser{}

// parseSharedExternalDocs parses an ExternalDocs object from a yaml.Node.
func parseSharedExternalDocs(node *yaml.Node, ctx *ParseContext) (*openapi30models.ExternalDocumentation, error) {
	return defaultExternalDocsParser.parse(node, ctx)
}

// Parse parses an ExternalDocs object.
func (p *externalDocsParser) parse(node *yaml.Node, ctx *ParseContext) (*openapi30models.ExternalDocumentation, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "externalDocs must be an object")
	}

	ed := &openapi30models.ExternalDocumentation{}

	// All properties are simple - inline
	ed.URL = p.ParseURL(node)
	ed.Description = p.ParseDescription(node)

	ed.VendorExtensions = parseNodeExtensions(node)
	ed.NodeSource = ctx.nodeSource(node)

	// Detect unknown fields
	ctx.detectUnknown(node, externalDocsKnownFieldsSet)

	return ed, nil
}

func (p *externalDocsParser) ParseURL(node *yaml.Node) string {
	return nodeGetString(node, "url")
}

func (p *externalDocsParser) ParseDescription(node *yaml.Node) string {
	return nodeGetString(node, "description")
}
