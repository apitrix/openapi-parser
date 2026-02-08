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

	param := &openapi30models.Parameter{}
	var err error

	// Simple properties - inline
	param.Name = p.ParseName(node)
	param.In = p.ParseIn(node)
	param.Description = p.ParseDescription(node)
	param.Required = p.ParseRequired(node)
	param.Deprecated = p.ParseDeprecated(node)
	param.AllowEmptyValue = p.ParseAllowEmptyValue(node)
	param.Style = p.ParseStyle(node)
	param.Explode = p.ParseExplode(node)
	param.AllowReserved = p.ParseAllowReserved(node)

	// Complex properties - delegated to dedicated files
	param.Schema, err = p.ParseSchema(node, ctx)
	if err != nil {
		return nil, err
	}

	param.Example = p.ParseExample(node)

	param.Examples, err = p.ParseExamples(node, ctx)
	if err != nil {
		return nil, err
	}

	param.Content, err = p.ParseContent(node, ctx)
	if err != nil {
		return nil, err
	}

	param.VendorExtensions = parseNodeExtensions(node)
	param.NodeSource = ctx.nodeSource(node)

	// Detect unknown fields
	ctx.detectUnknown(node, parameterKnownFieldsSet)

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
