package openapi30x

import (
	openapi30models "github.com/apitrix/openapi-parser/models/openapi30"

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

	// Collect values
	enum := p.ParseEnum(node)
	defaultVal := p.ParseDefault(node)
	description := p.ParseDescription(node)

	// Create via constructor
	sv := openapi30models.NewServerVariable(defaultVal, description, enum)

	sv.VendorExtensions = parseNodeExtensions(node)
	sv.Trix.Source = ctx.nodeSource(node)

	// Detect unknown fields
	sv.Trix.Errors = append(sv.Trix.Errors, unknownFieldParseErrors(ctx.detectUnknown(node, serverVariableKnownFieldsSet))...)

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
