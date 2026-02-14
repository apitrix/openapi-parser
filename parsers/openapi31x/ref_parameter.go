package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// parseParameterRef parses a ParameterRef from a yaml.Node.
func parseParameterRef(node *yaml.Node, ctx *ParseContext) (*shared.RefWithMeta[openapi31models.Parameter], error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "parameter must be an object")
	}

	ref := &shared.RefWithMeta[openapi31models.Parameter]{}
	ref.Trix.Source = ctx.nodeSource(node)
	ref.VendorExtensions = parseNodeExtensions(node)

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
	ref.SetValue(param)

	return ref, nil
}
