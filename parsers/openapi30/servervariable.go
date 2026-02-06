package openapi30

import (
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

type serverVariableParser struct{}

// defaultServerVariableParser is the singleton instance used by parsing functions.
var defaultServerVariableParser = &serverVariableParser{}

// parseSharedServerVariable parses a ServerVariable object from a yaml.Node.
func parseSharedServerVariable(node *yaml.Node, ctx *ParseContext) (*openapi30models.ServerVariable, error) {
	return defaultServerVariableParser.parse(node, ctx)
}

// Parse parses a ServerVariable object.
func (p *serverVariableParser) parse(node *yaml.Node, ctx *ParseContext) (*openapi30models.ServerVariable, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "serverVariable must be an object")
	}

	sv := &openapi30models.ServerVariable{}

	// All properties are simple - inline
	sv.Enum = p.ParseEnum(node)
	sv.Default = p.ParseDefault(node)
	sv.Description = p.ParseDescription(node)

	sv.Extensions = parseNodeExtensions(node)
	sv.NodeSource = ctx.nodeSource(node)

	// Detect unknown fields
	ctx.detectUnknown(node, serverVariableKnownFieldsSet)

	return sv, nil
}

func (p *serverVariableParser) ParseEnum(node *yaml.Node) []string {
	return nodeGetStringSlice(node, "enum")
}

func (p *serverVariableParser) ParseDefault(node *yaml.Node) string {
	return nodeGetString(node, "default")
}

func (p *serverVariableParser) ParseDescription(node *yaml.Node) string {
	return nodeGetString(node, "description")
}
