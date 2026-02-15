package openapi30x

import (
	openapi30models "github.com/apitrix/openapi-parser/models/openapi30"

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

	// Collect values
	url := p.ParseURL(node)
	description := p.ParseDescription(node)

	// Create via constructor
	ed := openapi30models.NewExternalDocumentation(url, description)

	ed.VendorExtensions = parseNodeExtensions(node)
	ed.Trix.Source = ctx.nodeSource(node)

	// Detect unknown fields
	ed.Trix.Errors = append(ed.Trix.Errors, unknownFieldParseErrors(ctx.detectUnknown(node, externalDocsKnownFieldsSet))...)

	return ed, nil
}

func (p *externalDocsParser) ParseURL(node *yaml.Node) string {
	return nodeGetString(node, "url")
}

func (p *externalDocsParser) ParseDescription(node *yaml.Node) string {
	return nodeGetString(node, "description")
}
