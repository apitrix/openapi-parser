package openapi30x

import (
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

type parameterParser struct{}

// defaultParameterParser is the singleton instance used by parsing functions.
var defaultParameterParser = &parameterParser{}

// parseSharedParameter parses a Parameter object from a yaml.Node.
func parseSharedParameter(node *yaml.Node, ctx *ParseContext) (*openapi30models.Parameter, error) {
	return defaultParameterParser.parse(node, ctx)
}

// Parse parses a Parameter object.
func (p *parameterParser) parse(node *yaml.Node, ctx *ParseContext) (*openapi30models.Parameter, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "parameter must be an object")
	}

	var errors []openapi30models.ParseError

	// Simple properties
	name := p.ParseName(node)
	in := p.ParseIn(node)
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
	param := openapi30models.NewParameter(name, in, description, required, deprecated, allowEmptyValue, style, explode, allowReserved, schema, example, examples, content)

	param.VendorExtensions = parseNodeExtensions(node)
	param.Trix.Source = ctx.nodeSource(node)
	param.Trix.Errors = append(param.Trix.Errors, errors...)

	// Detect unknown fields
	param.Trix.Errors = append(param.Trix.Errors, unknownFieldParseErrors(ctx.detectUnknown(node, parameterKnownFieldsSet))...)

	return param, nil
}

func (p *parameterParser) ParseName(node *yaml.Node) string {
	return nodeGetString(node, "name")
}

func (p *parameterParser) ParseIn(node *yaml.Node) string {
	return nodeGetString(node, "in")
}

func (p *parameterParser) ParseDescription(node *yaml.Node) string {
	return nodeGetString(node, "description")
}

func (p *parameterParser) ParseRequired(node *yaml.Node) bool {
	return nodeGetBool(node, "required")
}

func (p *parameterParser) ParseDeprecated(node *yaml.Node) bool {
	return nodeGetBool(node, "deprecated")
}

func (p *parameterParser) ParseAllowEmptyValue(node *yaml.Node) bool {
	return nodeGetBool(node, "allowEmptyValue")
}

func (p *parameterParser) ParseStyle(node *yaml.Node) string {
	return nodeGetString(node, "style")
}

func (p *parameterParser) ParseExplode(node *yaml.Node) *bool {
	return nodeGetBoolPtr(node, "explode")
}

func (p *parameterParser) ParseAllowReserved(node *yaml.Node) bool {
	return nodeGetBool(node, "allowReserved")
}

func (p *parameterParser) ParseExample(node *yaml.Node) interface{} {
	return nodeGetAny(node, "example")
}
