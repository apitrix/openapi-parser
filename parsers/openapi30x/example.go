package openapi30x

import (
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

type exampleParser struct{}

// defaultExampleParser is the singleton instance used by parsing functions.
var defaultExampleParser = &exampleParser{}

// parseSharedExample parses an Example object from a yaml.Node.
func parseSharedExample(node *yaml.Node, ctx *ParseContext) (*openapi30models.Example, error) {
	return defaultExampleParser.parse(node, ctx)
}

// Parse parses an Example object.
func (p *exampleParser) parse(node *yaml.Node, ctx *ParseContext) (*openapi30models.Example, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "example must be an object")
	}

	example := &openapi30models.Example{}

	// All properties are simple - inline
	example.Summary = p.ParseSummary(node)
	example.Description = p.ParseDescription(node)
	example.Value = p.ParseValue(node)
	example.ExternalValue = p.ParseExternalValue(node)

	example.VendorExtensions = parseNodeExtensions(node)
	example.Trix.Source = ctx.nodeSource(node)

	// Detect unknown fields
	example.Trix.Errors = append(example.Trix.Errors, unknownFieldParseErrors(ctx.detectUnknown(node, exampleKnownFieldsSet))...)

	return example, nil
}

func (p *exampleParser) ParseSummary(node *yaml.Node) string {
	return nodeGetString(node, "summary")
}

func (p *exampleParser) ParseDescription(node *yaml.Node) string {
	return nodeGetString(node, "description")
}

func (p *exampleParser) ParseValue(node *yaml.Node) interface{} {
	return nodeGetAny(node, "value")
}

func (p *exampleParser) ParseExternalValue(node *yaml.Node) string {
	return nodeGetString(node, "externalValue")
}
