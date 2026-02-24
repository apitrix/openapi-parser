package openapi31x

import (
	openapi31models "github.com/apitrix/openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

type exampleParser struct{}

// defaultExampleParser is the singleton instance used by parsing functions.
var defaultExampleParser = &exampleParser{}

// parseSharedExample parses an Example object from a yaml.Node.
func parseSharedExample(node *yaml.Node, ctx *ParseContext) (*openapi31models.Example, error) {
	return defaultExampleParser.parse(node, ctx)
}

// Parse parses an Example object.
func (p *exampleParser) parse(node *yaml.Node, ctx *ParseContext) (*openapi31models.Example, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "example must be an object")
	}

	// Create via constructor
	example := openapi31models.NewExample(
		p.ParseSummary(node),
		p.ParseDescription(node),
		p.ParseValue(node),
		p.ParseExternalValue(node),
	)

	example.VendorExtensions = parseNodeExtensions(node)
	example.Trix.Source = ctx.nodeSource(node)

	// Set OpenAPI 3.2 fields via setters
	_ = example.SetDataValue(p.ParseDataValue(node))
	_ = example.SetSerializedValue(p.ParseSerializedValue(node))

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

func (p *exampleParser) ParseDataValue(node *yaml.Node) interface{} {
	return nodeGetAny(node, "dataValue")
}

func (p *exampleParser) ParseSerializedValue(node *yaml.Node) string {
	return nodeGetString(node, "serializedValue")
}
