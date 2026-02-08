package openapi30x

import (
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

type requestBodyParser struct{}

// defaultRequestBodyParser is the singleton instance used by parsing functions.
var defaultRequestBodyParser = &requestBodyParser{}

// parseSharedRequestBody parses a RequestBody object from a yaml.Node.
func parseSharedRequestBody(node *yaml.Node, ctx *ParseContext) (*openapi30models.RequestBody, error) {
	return defaultRequestBodyParser.parse(node, ctx)
}

// Parse parses a RequestBody object.
func (p *requestBodyParser) parse(node *yaml.Node, ctx *ParseContext) (*openapi30models.RequestBody, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "requestBody must be an object")
	}

	rb := &openapi30models.RequestBody{}
	var err error

	// Simple properties - inline
	rb.Description = p.ParseDescription(node)
	rb.Required = p.ParseRequired(node)

	// Complex properties - delegated to dedicated files
	rb.Content, err = p.ParseContent(node, ctx)
	if err != nil {
		rb.Trix.Errors = append(rb.Trix.Errors, toParseError(err))
	}

	rb.VendorExtensions = parseNodeExtensions(node)
	rb.Trix.Source = ctx.nodeSource(node)

	// Detect unknown fields
	ctx.detectUnknown(node, requestBodyKnownFieldsSet)

	return rb, nil
}

func (p *requestBodyParser) ParseDescription(node *yaml.Node) string {
	return nodeGetString(node, "description")
}

func (p *requestBodyParser) ParseRequired(node *yaml.Node) bool {
	return nodeGetBool(node, "required")
}
