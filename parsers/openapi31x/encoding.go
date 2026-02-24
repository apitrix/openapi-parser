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

	// Set OpenAPI 3.2 fields via setters
	_ = enc.SetEncoding(p.ParseEncoding(node, ctx))
	_ = enc.SetPrefixEncoding(p.ParsePrefixEncoding(node, ctx))
	_ = enc.SetItemEncoding(p.ParseItemEncoding(node, ctx))

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

func (p *encodingParser) ParseEncoding(node *yaml.Node, ctx *ParseContext) map[string]*openapi31models.Encoding {
	n := nodeGetValue(node, "encoding")
	if n == nil || !nodeIsMapping(n) {
		return nil
	}
	out := make(map[string]*openapi31models.Encoding)
	for i := 0; i < len(n.Content)-1; i += 2 {
		key := n.Content[i].Value
		enc, err := parseSharedEncoding(n.Content[i+1], ctx.Push(key))
		if err == nil && enc != nil {
			out[key] = enc
		}
	}
	return out
}

func (p *encodingParser) ParsePrefixEncoding(node *yaml.Node, ctx *ParseContext) []*openapi31models.Encoding {
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

func (p *encodingParser) ParseItemEncoding(node *yaml.Node, ctx *ParseContext) *openapi31models.Encoding {
	n := nodeGetValue(node, "itemEncoding")
	if n == nil {
		return nil
	}
	enc, _ := parseSharedEncoding(n, ctx.Push("itemEncoding"))
	return enc
}
