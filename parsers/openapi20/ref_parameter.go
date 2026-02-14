package openapi20

import (
	"openapi-parser/models/shared"
	openapi20models "openapi-parser/models/openapi20"

	"gopkg.in/yaml.v3"
)

// parseParameterRef parses a ParameterRef (either $ref or inline parameter) from a yaml.Node.
func parseParameterRef(node *yaml.Node, ctx *ParseContext) (*shared.Ref[openapi20models.Parameter], error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "parameter must be an object")
	}

	ref := &shared.Ref[openapi20models.Parameter]{}

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
	ref.SetValue(param)
	ref.Trix.Source = ctx.nodeSource(node)

	return ref, nil
}
