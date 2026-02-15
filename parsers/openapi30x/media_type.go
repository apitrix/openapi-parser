package openapi30x

import (
	openapi30models "github.com/apitrix/openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

type mediaTypeParser struct{}

// defaultMediaTypeParser is the singleton instance used by parsing functions.
var defaultMediaTypeParser = &mediaTypeParser{}

// parseSharedMediaType parses a MediaType object from a yaml.Node.
func parseSharedMediaType(node *yaml.Node, ctx *ParseContext) (*openapi30models.MediaType, error) {
	return defaultMediaTypeParser.parse(node, ctx)
}

// Parse parses a MediaType object.
func (p *mediaTypeParser) parse(node *yaml.Node, ctx *ParseContext) (*openapi30models.MediaType, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "mediaType must be an object")
	}

	var errors []openapi30models.ParseError

	// Simple properties
	example := p.ParseExample(node)

	// Complex properties - delegated to dedicated files
	schema, err := p.ParseSchema(node, ctx)
	if err != nil {
		errors = append(errors, toParseError(err))
	}

	examples, err := p.ParseExamples(node, ctx)
	if err != nil {
		errors = append(errors, toParseError(err))
	}

	encoding, err := p.ParseEncoding(node, ctx)
	if err != nil {
		errors = append(errors, toParseError(err))
	}

	// Create via constructor
	mt := openapi30models.NewMediaType(schema, example, examples, encoding)

	mt.VendorExtensions = parseNodeExtensions(node)
	mt.Trix.Source = ctx.nodeSource(node)
	mt.Trix.Errors = append(mt.Trix.Errors, errors...)

	// Detect unknown fields
	mt.Trix.Errors = append(mt.Trix.Errors, unknownFieldParseErrors(ctx.detectUnknown(node, mediaTypeKnownFieldsSet))...)

	return mt, nil
}

func (p *mediaTypeParser) ParseExample(node *yaml.Node) interface{} {
	return nodeGetAny(node, "example")
}
