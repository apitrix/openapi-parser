package openapi31x

import (
	openapi31models "github.com/apitrix/openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

type mediaTypeParser struct{}

// defaultMediaTypeParser is the singleton instance used by parsing functions.
var defaultMediaTypeParser = &mediaTypeParser{}

// parseSharedMediaType parses a MediaType object from a yaml.Node.
func parseSharedMediaType(node *yaml.Node, ctx *ParseContext) (*openapi31models.MediaType, error) {
	return defaultMediaTypeParser.parse(node, ctx)
}

// Parse parses a MediaType object.
func (p *mediaTypeParser) parse(node *yaml.Node, ctx *ParseContext) (*openapi31models.MediaType, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "mediaType must be an object")
	}

	var errs []openapi31models.ParseError

	// Complex properties - delegated to dedicated files
	schema, err := p.ParseSchema(node, ctx)
	if err != nil {
		errs = append(errs, toParseError(err))
	}

	examples, err := p.ParseExamples(node, ctx)
	if err != nil {
		errs = append(errs, toParseError(err))
	}

	encoding, err := p.ParseEncoding(node, ctx)
	if err != nil {
		errs = append(errs, toParseError(err))
	}

	// Create via constructor
	mt := openapi31models.NewMediaType(
		schema,
		p.ParseExample(node),
		examples,
		encoding,
	)

	mt.VendorExtensions = parseNodeExtensions(node)
	mt.Trix.Source = ctx.nodeSource(node)
	mt.Trix.Errors = append(mt.Trix.Errors, errs...)

	// Set OpenAPI 3.2 fields via setters
	_ = mt.SetDescription(p.ParseDescription(node))
	_ = mt.SetItemSchema(p.ParseItemSchema(node, ctx))
	_ = mt.SetPrefixEncoding(p.ParsePrefixEncoding(node, ctx))
	_ = mt.SetItemEncoding(p.ParseItemEncoding(node, ctx))

	// Detect unknown fields
	mt.Trix.Errors = append(mt.Trix.Errors, unknownFieldParseErrors(ctx.detectUnknown(node, mediaTypeKnownFieldsSet))...)

	return mt, nil
}

func (p *mediaTypeParser) ParseExample(node *yaml.Node) interface{} {
	return nodeGetAny(node, "example")
}

func (p *mediaTypeParser) ParseDescription(node *yaml.Node) string {
	return nodeGetString(node, "description")
}

func (p *mediaTypeParser) ParseItemSchema(node *yaml.Node, ctx *ParseContext) *openapi31models.RefSchema {
	schemaNode := nodeGetValue(node, "itemSchema")
	if schemaNode == nil {
		return nil
	}
	ref, _ := parseSchemaRef(schemaNode, ctx.Push("itemSchema"))
	return ref
}

func (p *mediaTypeParser) ParsePrefixEncoding(node *yaml.Node, ctx *ParseContext) []*openapi31models.Encoding {
	n := nodeGetValue(node, "prefixEncoding")
	if n == nil || !nodeIsSequence(n) {
		return nil
	}
	out := make([]*openapi31models.Encoding, 0, len(n.Content))
	for i, item := range n.Content {
		enc, err := parseSharedEncoding(item, ctx.Push(itoa(i)))
		if err == nil && enc != nil {
			out = append(out, enc)
		}
	}
	return out
}

func (p *mediaTypeParser) ParseItemEncoding(node *yaml.Node, ctx *ParseContext) *openapi31models.Encoding {
	n := nodeGetValue(node, "itemEncoding")
	if n == nil {
		return nil
	}
	enc, _ := parseSharedEncoding(n, ctx.Push("itemEncoding"))
	return enc
}
