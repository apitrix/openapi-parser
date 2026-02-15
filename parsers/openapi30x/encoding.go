package openapi30x

import (
	openapi30models "github.com/apitrix/openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

type encodingParser struct{}

// defaultEncodingParser is the singleton instance used by parsing functions.
var defaultEncodingParser = &encodingParser{}

// parseSharedEncoding parses an Encoding object from a yaml.Node.
func parseSharedEncoding(node *yaml.Node, ctx *ParseContext) (*openapi30models.Encoding, error) {
	return defaultEncodingParser.parse(node, ctx)
}

// Parse parses an Encoding object.
func (p *encodingParser) parse(node *yaml.Node, ctx *ParseContext) (*openapi30models.Encoding, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "encoding must be an object")
	}

	// Collect values
	contentType := p.ParseContentType(node)
	style := p.ParseStyle(node)
	explode := p.ParseExplode(node)
	allowReserved := p.ParseAllowReserved(node)

	headers, err := p.ParseHeaders(node, ctx)

	// Create via constructor
	enc := openapi30models.NewEncoding(contentType, headers, style, explode, allowReserved)

	if err != nil {
		enc.Trix.Errors = append(enc.Trix.Errors, toParseError(err))
	}

	enc.VendorExtensions = parseNodeExtensions(node)
	enc.Trix.Source = ctx.nodeSource(node)

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
