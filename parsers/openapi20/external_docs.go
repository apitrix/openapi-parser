package openapi20

import (
	openapi20models "github.com/apitrix/openapi-parser/models/openapi20"

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

	ed := openapi20models.NewExternalDocs(
		nodeGetString(node, "description"),
		nodeGetString(node, "url"),
	)

	ed.VendorExtensions = parseNodeExtensions(node)
	ed.Trix.Source = ctx.nodeSource(node)

	// Detect unknown fields
	ed.Trix.Errors = append(ed.Trix.Errors, unknownFieldParseErrors(ctx.detectUnknown(node, externalDocsKnownFieldsSet))...)

	return ed, nil
}
