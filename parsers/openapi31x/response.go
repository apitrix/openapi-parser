package openapi31x

import (
	openapi31models "github.com/apitrix/openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

type responseParser struct{}

// defaultResponseParser is the singleton instance used by parsing functions.
var defaultResponseParser = &responseParser{}

// parseSharedResponse parses a Response object from a yaml.Node.
func parseSharedResponse(node *yaml.Node, ctx *ParseContext) (*openapi31models.Response, error) {
	return defaultResponseParser.parse(node, ctx)
}

// Parse parses a Response object.
func (p *responseParser) parse(node *yaml.Node, ctx *ParseContext) (*openapi31models.Response, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "response must be an object")
	}

	var errs []openapi31models.ParseError

	// Complex properties - delegated to dedicated files
	headers, err := p.ParseHeaders(node, ctx)
	if err != nil {
		errs = append(errs, toParseError(err))
	}

	content, err := p.ParseContent(node, ctx)
	if err != nil {
		errs = append(errs, toParseError(err))
	}

	links, err := p.ParseLinks(node, ctx)
	if err != nil {
		errs = append(errs, toParseError(err))
	}

	// Create via constructor
	resp := openapi31models.NewResponse(
		p.ParseDescription(node),
		headers,
		content,
		links,
	)

	resp.VendorExtensions = parseNodeExtensions(node)
	resp.Trix.Source = ctx.nodeSource(node)
	resp.Trix.Errors = append(resp.Trix.Errors, errs...)

	// Detect unknown fields
	resp.Trix.Errors = append(resp.Trix.Errors, unknownFieldParseErrors(ctx.detectUnknown(node, responseKnownFieldsSet))...)

	return resp, nil
}

func (p *responseParser) ParseDescription(node *yaml.Node) string {
	return nodeGetString(node, "description")
}
