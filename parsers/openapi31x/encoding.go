package openapi31x

import (
	openapi31models "github.com/apitrix/openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

type encodingParser struct{}

// defaultEncodingParser is the singleton instance used by parsing functions.
var defaultEncodingParser = &encodingParser{}

// parseSharedEncoding parses an Encoding object from a yaml.Node.
func parseSharedEncoding(node *yaml.Node, ctx *ParseContext) (*openapi31models.Encoding, error) {
	return defaultEncodingParser.parse(node, ctx)
}

// Parse parses an Encoding object.
func (p *encodingParser) parse(node *yaml.Node, ctx *ParseContext) (*openapi31models.Encoding, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "encoding must be an object")
	}

	var errs []openapi31models.ParseError

	// Complex properties - delegated to dedicated files
	headers, err := p.ParseHeaders(node, ctx)
	if err != nil {
		errs = append(errs, toParseError(err))
	}

	// Create via constructor
	enc := openapi31models.NewEncoding(
		p.ParseContentType(node),
		p.ParseStyle(node),
		headers,
		p.ParseExplode(node),
		p.ParseAllowReserved(node),
	)

	enc.VendorExtensions = parseNodeExtensions(node)
	enc.Trix.Source = ctx.nodeSource(node)
	enc.Trix.Errors = append(enc.Trix.Errors, errs...)

	// Detect unknown fields
	enc.Trix.Errors = append(enc.Trix.Errors, unknownFieldParseErrors(ctx.detectUnknown(node, encodingKnownFieldsSet))...)

	return enc, nil
}

func (p *encodingParser) ParseContentType(node *yaml.Node) string {
	return nodeGetString(node, "contentType")
}

func (p *encodingParser) ParseStyle(node *yaml.Node) string {
	return nodeGetString(node, "style")
}

func (p *encodingParser) ParseExplode(node *yaml.Node) *bool {
	return nodeGetBoolPtr(node, "explode")
}

func (p *encodingParser) ParseAllowReserved(node *yaml.Node) bool {
	return nodeGetBool(node, "allowReserved")
}
