package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

type parameterParser struct{}

// defaultParameterParser is the singleton instance used by parsing functions.
var defaultParameterParser = &parameterParser{}

// parseSharedParameter parses a Parameter object from a yaml.Node.
func parseSharedParameter(node *yaml.Node, ctx *ParseContext) (*openapi31models.Parameter, error) {
	return defaultParameterParser.parse(node, ctx)
}

// Parse parses a Parameter object.
func (p *parameterParser) parse(node *yaml.Node, ctx *ParseContext) (*openapi31models.Parameter, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "parameter must be an object")
	}

	var errs []openapi31models.ParseError

	// Complex properties - delegated to dedicated files
	schema, err := p.ParseSchema(node, ctx)
	if err != nil {
		errs = append(errs, toParseError(err))
	}

	examples, err := p.ParseExamples(node, ctx)
	if err != nil {
		errs = append(errs, toParseError(err))
	}

	content, err := p.ParseContent(node, ctx)
	if err != nil {
		errs = append(errs, toParseError(err))
	}

	// Create via constructor with ParameterFields
	param := openapi31models.NewParameter(openapi31models.ParameterFields{
		Name:            p.ParseName(node),
		In:              p.ParseIn(node),
		Description:     p.ParseDescription(node),
		Required:        p.ParseRequired(node),
		Deprecated:      p.ParseDeprecated(node),
		AllowEmptyValue: p.ParseAllowEmptyValue(node),
		Style:           p.ParseStyle(node),
		Explode:         p.ParseExplode(node),
		AllowReserved:   p.ParseAllowReserved(node),
		Schema:          schema,
		Example:         p.ParseExample(node),
		Examples:        examples,
		Content:         content,
	})

	param.VendorExtensions = parseNodeExtensions(node)
	param.Trix.Source = ctx.nodeSource(node)
	param.Trix.Errors = append(param.Trix.Errors, errs...)

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
