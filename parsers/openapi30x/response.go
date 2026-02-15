package openapi30x

import (
	openapi30models "github.com/apitrix/openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

type responseParser struct{}

// defaultResponseParser is the singleton instance used by parsing functions.
var defaultResponseParser = &responseParser{}

// parseSharedResponse parses a Response object from a yaml.Node.
func parseSharedResponse(node *yaml.Node, ctx *ParseContext) (*openapi30models.Response, error) {
	return defaultResponseParser.parse(node, ctx)
}

// Parse parses a Response object.
func (p *responseParser) parse(node *yaml.Node, ctx *ParseContext) (*openapi30models.Response, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "response must be an object")
	}

	var errors []openapi30models.ParseError

	// Simple properties
	description := p.ParseDescription(node)

	// Complex properties - delegated to dedicated files
	headers, err := p.ParseHeaders(node, ctx)
	if err != nil {
		errors = append(errors, toParseError(err))
	}

	content, err := p.ParseContent(node, ctx)
	if err != nil {
		errors = append(errors, toParseError(err))
	}

	links, err := p.ParseLinks(node, ctx)
	if err != nil {
		errors = append(errors, toParseError(err))
	}

	// Create via constructor
	resp := openapi30models.NewResponse(description, headers, content, links)

	resp.VendorExtensions = parseNodeExtensions(node)
	resp.Trix.Source = ctx.nodeSource(node)
	resp.Trix.Errors = append(resp.Trix.Errors, errors...)

	// Detect unknown fields
	resp.Trix.Errors = append(resp.Trix.Errors, unknownFieldParseErrors(ctx.detectUnknown(node, responseKnownFieldsSet))...)

	return resp, nil
}

func (p *responseParser) ParseDescription(node *yaml.Node) string {
	return nodeGetString(node, "description")
}
