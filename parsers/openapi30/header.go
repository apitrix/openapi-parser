package openapi30

import (
	openapi30models "openapi-parser/models/openapi30"
	"gopkg.in/yaml.v3"
)

type headerParser struct{}

// defaultHeaderParser is the singleton instance used by parsing functions.
var defaultHeaderParser = &headerParser{}

// parseSharedHeader parses a Header object from a yaml.Node.
func parseSharedHeader(node *yaml.Node, ctx *ParseContext) (*openapi30models.Header, error) {
	return defaultHeaderParser.Parse(node, ctx)
}

// Parse parses a Header object.
func (p *headerParser) Parse(node *yaml.Node, ctx *ParseContext) (*openapi30models.Header, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "header must be an object")
	}

	header := &openapi30models.Header{}
	var err error

	// Simple properties - inline
	header.Description = p.ParseDescription(node)
	header.Required = p.ParseRequired(node)
	header.Deprecated = p.ParseDeprecated(node)
	header.AllowEmptyValue = p.ParseAllowEmptyValue(node)
	header.Style = p.ParseStyle(node)
	header.Explode = p.ParseExplode(node)
	header.AllowReserved = p.ParseAllowReserved(node)
	header.Example = p.ParseExample(node)

	// Complex properties - delegated to dedicated files
	header.Schema, err = p.ParseSchema(node, ctx)
	if err != nil {
		return nil, err
	}

	header.Examples, err = p.ParseExamples(node, ctx)
	if err != nil {
		return nil, err
	}

	header.Content, err = p.ParseContent(node, ctx)
	if err != nil {
		return nil, err
	}

	header.Extensions = parseNodeExtensions(node)
	header.NodeSource = ctx.nodeSource(node)

	// Detect unknown fields
	ctx.detectUnknown(node, headerKnownFields)

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
