package openapi31x

import (
	openapi31models "github.com/apitrix/openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

type requestBodyParser struct{}

// defaultRequestBodyParser is the singleton instance used by parsing functions.
var defaultRequestBodyParser = &requestBodyParser{}

// parseSharedRequestBody parses a RequestBody object from a yaml.Node.
func parseSharedRequestBody(node *yaml.Node, ctx *ParseContext) (*openapi31models.RequestBody, error) {
	return defaultRequestBodyParser.parse(node, ctx)
}

// Parse parses a RequestBody object.
func (p *requestBodyParser) parse(node *yaml.Node, ctx *ParseContext) (*openapi31models.RequestBody, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "requestBody must be an object")
	}

	var errs []openapi31models.ParseError

	// Complex properties - delegated to dedicated files
	content, err := p.ParseContent(node, ctx)
	if err != nil {
		errs = append(errs, toParseError(err))
	}

	// Create via constructor
	rb := openapi31models.NewRequestBody(
		p.ParseDescription(node),
		content,
		p.ParseRequired(node),
	)

	rb.VendorExtensions = parseNodeExtensions(node)
	rb.Trix.Source = ctx.nodeSource(node)
	rb.Trix.Errors = append(rb.Trix.Errors, errs...)

	// Detect unknown fields
	rb.Trix.Errors = append(rb.Trix.Errors, unknownFieldParseErrors(ctx.detectUnknown(node, requestBodyKnownFieldsSet))...)

	return rb, nil
}

func (p *requestBodyParser) ParseDescription(node *yaml.Node) string {
	return nodeGetString(node, "description")
}

func (p *requestBodyParser) ParseRequired(node *yaml.Node) bool {
	return nodeGetBool(node, "required")
}
