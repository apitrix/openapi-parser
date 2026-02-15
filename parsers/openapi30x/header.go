package openapi30x

import (
	openapi30models "github.com/apitrix/openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

type headerParser struct{}

// defaultHeaderParser is the singleton instance used by parsing functions.
var defaultHeaderParser = &headerParser{}

// parseSharedHeader parses a Header object from a yaml.Node.
func parseSharedHeader(node *yaml.Node, ctx *ParseContext) (*openapi30models.Header, error) {
	return defaultHeaderParser.parse(node, ctx)
}

// Parse parses a Header object.
func (p *headerParser) parse(node *yaml.Node, ctx *ParseContext) (*openapi30models.Header, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "header must be an object")
	}

	var errors []openapi30models.ParseError

	// Simple properties
	description := p.ParseDescription(node)
	required := p.ParseRequired(node)
	deprecated := p.ParseDeprecated(node)
	allowEmptyValue := p.ParseAllowEmptyValue(node)
	style := p.ParseStyle(node)
	explode := p.ParseExplode(node)
	allowReserved := p.ParseAllowReserved(node)
	example := p.ParseExample(node)

	// Complex properties - delegated to dedicated files
	schema, err := p.ParseSchema(node, ctx)
	if err != nil {
		errors = append(errors, toParseError(err))
	}

	examples, err := p.ParseExamples(node, ctx)
	if err != nil {
		errors = append(errors, toParseError(err))
	}

	content, err := p.ParseContent(node, ctx)
	if err != nil {
		errors = append(errors, toParseError(err))
	}

	// Create via constructor
	header := openapi30models.NewHeader(description, required, deprecated, allowEmptyValue, style, explode, allowReserved, schema, example, examples, content)

	header.VendorExtensions = parseNodeExtensions(node)
	header.Trix.Source = ctx.nodeSource(node)
	header.Trix.Errors = append(header.Trix.Errors, errors...)

	// Detect unknown fields
	header.Trix.Errors = append(header.Trix.Errors, unknownFieldParseErrors(ctx.detectUnknown(node, headerKnownFieldsSet))...)

	return header, nil
}

func (p *headerParser) ParseDescription(node *yaml.Node) string {
	return nodeGetString(node, "description")
}

func (p *headerParser) ParseRequired(node *yaml.Node) bool {
	return nodeGetBool(node, "required")
}

func (p *headerParser) ParseDeprecated(node *yaml.Node) bool {
	return nodeGetBool(node, "deprecated")
}

func (p *headerParser) ParseAllowEmptyValue(node *yaml.Node) bool {
	return nodeGetBool(node, "allowEmptyValue")
}

func (p *headerParser) ParseStyle(node *yaml.Node) string {
	return nodeGetString(node, "style")
}

func (p *headerParser) ParseExplode(node *yaml.Node) *bool {
	return nodeGetBoolPtr(node, "explode")
}

func (p *headerParser) ParseAllowReserved(node *yaml.Node) bool {
	return nodeGetBool(node, "allowReserved")
}

func (p *headerParser) ParseExample(node *yaml.Node) interface{} {
	return nodeGetAny(node, "example")
}
