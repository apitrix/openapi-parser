package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// parseParameterRef parses a ParameterRef from a yaml.Node.
func parseParameterRef(node *yaml.Node, ctx *ParseContext) (*openapi31models.ParameterRef, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "parameter must be an object")
	}

	ref := &openapi31models.ParameterRef{}
	ref.NodeSource = ctx.nodeSource(node)
	ref.Extensions = parseNodeExtensions(node)

	if nodeHasRef(node) {
		ref.Ref = nodeGetRef(node)
		ref.Summary = nodeGetString(node, "summary")
		ref.Description = nodeGetString(node, "description")
		return ref, nil
	}

	param, err := parseSharedParameter(node, ctx)
	if err != nil {
		return nil, err
	}
	ref.Value = param

	return ref, nil
}
