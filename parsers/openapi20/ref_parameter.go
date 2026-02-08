package openapi20

import (
	openapi20models "openapi-parser/models/openapi20"

	"gopkg.in/yaml.v3"
)

// parseParameterRef parses a ParameterRef (either $ref or inline parameter) from a yaml.Node.
func parseParameterRef(node *yaml.Node, ctx *ParseContext) (*openapi20models.ParameterRef, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "parameter must be an object")
	}

	ref := &openapi20models.ParameterRef{}

	// Check if it's a reference
	if nodeHasRef(node) {
		ref.Ref = nodeGetRef(node)
		ref.Trix.Source = ctx.nodeSource(node)
		return ref, nil
	}

	// Parse inline parameter
	param, err := parseParameter(node, ctx)
	if err != nil {
		return nil, err
	}
	ref.Value = param
	ref.Trix.Source = ctx.nodeSource(node)

	return ref, nil
}
