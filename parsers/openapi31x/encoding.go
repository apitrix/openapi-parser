package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"

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

	enc := &openapi31models.Encoding{}
	var err error

	// Simple properties - inline
	enc.ContentType = p.ParseContentType(node)
	enc.Style = p.ParseStyle(node)
	enc.Explode = p.ParseExplode(node)
	enc.AllowReserved = p.ParseAllowReserved(node)

	// Complex properties - delegated to dedicated files
	enc.Headers, err = p.ParseHeaders(node, ctx)
	if err != nil {
		enc.Trix.Errors = append(enc.Trix.Errors, toParseError(err))
	}

	enc.VendorExtensions = parseNodeExtensions(node)
	enc.Trix.Source = ctx.nodeSource(node)

	// Detect unknown fields
	ctx.detectUnknown(node, encodingKnownFieldsSet)

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
