package openapi30x

import (
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

type linkParser struct{}

// defaultLinkParser is the singleton instance used by parsing functions.
var defaultLinkParser = &linkParser{}

// parseSharedLink parses a Link object from a yaml.Node.
func parseSharedLink(node *yaml.Node, ctx *ParseContext) (*openapi30models.Link, error) {
	return defaultLinkParser.parse(node, ctx)
}

// Parse parses a Link object.
func (p *linkParser) parse(node *yaml.Node, ctx *ParseContext) (*openapi30models.Link, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "link must be an object")
	}

	var errors []openapi30models.ParseError

	// Simple properties
	operationRef := p.ParseOperationRef(node)
	operationID := p.ParseOperationID(node)
	requestBody := p.ParseRequestBody(node)
	description := p.ParseDescription(node)

	// Complex properties - delegated to dedicated files
	parameters, err := p.ParseParameters(node, ctx)
	if err != nil {
		errors = append(errors, toParseError(err))
	}

	server, err := p.ParseServer(node, ctx)
	if err != nil {
		errors = append(errors, toParseError(err))
	}

	// Create via constructor
	link := openapi30models.NewLink(operationRef, operationID, parameters, requestBody, description, server)

	link.VendorExtensions = parseNodeExtensions(node)
	link.Trix.Source = ctx.nodeSource(node)
	link.Trix.Errors = append(link.Trix.Errors, errors...)

	// Detect unknown fields
	link.Trix.Errors = append(link.Trix.Errors, unknownFieldParseErrors(ctx.detectUnknown(node, linkKnownFieldsSet))...)

	return link, nil
}

func (p *linkParser) ParseOperationRef(node *yaml.Node) string {
	return nodeGetString(node, "operationRef")
}

func (p *linkParser) ParseOperationID(node *yaml.Node) string {
	return nodeGetString(node, "operationId")
}

func (p *linkParser) ParseRequestBody(node *yaml.Node) interface{} {
	return nodeGetAny(node, "requestBody")
}

func (p *linkParser) ParseDescription(node *yaml.Node) string {
	return nodeGetString(node, "description")
}
