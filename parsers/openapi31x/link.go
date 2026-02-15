package openapi31x

import (
	openapi31models "github.com/apitrix/openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

type linkParser struct{}

// defaultLinkParser is the singleton instance used by parsing functions.
var defaultLinkParser = &linkParser{}

// parseSharedLink parses a Link object from a yaml.Node.
func parseSharedLink(node *yaml.Node, ctx *ParseContext) (*openapi31models.Link, error) {
	return defaultLinkParser.parse(node, ctx)
}

// Parse parses a Link object.
func (p *linkParser) parse(node *yaml.Node, ctx *ParseContext) (*openapi31models.Link, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "link must be an object")
	}

	var errs []openapi31models.ParseError

	// Complex properties - delegated to dedicated files
	parameters, err := p.ParseParameters(node, ctx)
	if err != nil {
		errs = append(errs, toParseError(err))
	}

	server, err := p.ParseServer(node, ctx)
	if err != nil {
		errs = append(errs, toParseError(err))
	}

	// Create via constructor
	link := openapi31models.NewLink(
		p.ParseOperationRef(node),
		p.ParseOperationID(node),
		p.ParseDescription(node),
		parameters,
		p.ParseRequestBody(node),
		server,
	)

	link.VendorExtensions = parseNodeExtensions(node)
	link.Trix.Source = ctx.nodeSource(node)
	link.Trix.Errors = append(link.Trix.Errors, errs...)

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
