package openapi30x

import (
	openapi30models "openapi-parser/models/openapi30"

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

	mt := &openapi30models.MediaType{}
	var err error

	// Simple properties - inline
	mt.Example = p.ParseExample(node)

	// Complex properties - delegated to dedicated files
	mt.Schema, err = p.ParseSchema(node, ctx)
	if err != nil {
		return nil, err
	}

	mt.Examples, err = p.ParseExamples(node, ctx)
	if err != nil {
		return nil, err
	}

	mt.Encoding, err = p.ParseEncoding(node, ctx)
	if err != nil {
		return nil, err
	}

	mt.VendorExtensions = parseNodeExtensions(node)
	mt.Trix.Source = ctx.nodeSource(node)

	// Detect unknown fields
	ctx.detectUnknown(node, mediaTypeKnownFieldsSet)

	return mt, nil
}

func (p *mediaTypeParser) ParseExample(node *yaml.Node) interface{} {
	return nodeGetAny(node, "example")
}
