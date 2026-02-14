package openapi30x

import (
	"openapi-parser/models/shared"
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// parseParameterRef parses a ParameterRef from a yaml.Node.
func parseParameterRef(node *yaml.Node, ctx *ParseContext) (*shared.Ref[openapi30models.Parameter], error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "parameter must be an object")
	}

	ref := &shared.Ref[openapi30models.Parameter]{}
	ref.Trix.Source = ctx.nodeSource(node)
	ref.VendorExtensions = parseNodeExtensions(node)

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
	ref.SetValue(parameter)

	return ref, nil
}
