package openapi30x

import (
	openapi30models "openapi-parser/models/openapi30"

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

	resp := &openapi30models.Response{}
	var err error

	// Simple properties - inline
	resp.Description = p.ParseDescription(node)

	// Complex properties - delegated to dedicated files
	resp.Headers, err = p.ParseHeaders(node, ctx)
	if err != nil {
		resp.Trix.Errors = append(resp.Trix.Errors, toParseError(err))
	}

	resp.Content, err = p.ParseContent(node, ctx)
	if err != nil {
		resp.Trix.Errors = append(resp.Trix.Errors, toParseError(err))
	}

	resp.Links, err = p.ParseLinks(node, ctx)
	if err != nil {
		resp.Trix.Errors = append(resp.Trix.Errors, toParseError(err))
	}

	resp.VendorExtensions = parseNodeExtensions(node)
	resp.Trix.Source = ctx.nodeSource(node)

	// Detect unknown fields
	resp.Trix.Errors = append(resp.Trix.Errors, unknownFieldParseErrors(ctx.detectUnknown(node, responseKnownFieldsSet))...)

	return resp, nil
}

func (p *responseParser) ParseDescription(node *yaml.Node) string {
	return nodeGetString(node, "description")
}
