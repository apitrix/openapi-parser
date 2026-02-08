package openapi30x

import (
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// parseParameterRef parses a ParameterRef from a yaml.Node.
func parseParameterRef(node *yaml.Node, ctx *ParseContext) (*openapi30models.ParameterRef, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "parameter must be an object")
	}

	ref := &openapi30models.ParameterRef{}
	ref.NodeSource = ctx.nodeSource(node)
	ref.Extensions = parseNodeExtensions(node)

	// Check for $ref
	if nodeHasRef(node) {
		ref.Ref = nodeGetRef(node)
		return ref, nil
	}

	// Parse inline parameter
	parameter, err := parseSharedParameter(node, ctx)
	if err != nil {
		return nil, err
	}
	ref.Value = parameter

	return ref, nil
}
