package openapi31x

import (
	"openapi-parser/models/shared"
	openapi31models "openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// parseResponseRef parses a ResponseRef from a yaml.Node.
func parseResponseRef(node *yaml.Node, ctx *ParseContext) (*shared.RefWithMeta[openapi31models.Response], error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "response must be an object")
	}

	ref := &shared.RefWithMeta[openapi31models.Response]{}
	ref.Trix.Source = ctx.nodeSource(node)
	ref.VendorExtensions = parseNodeExtensions(node)

	if nodeHasRef(node) {
		ref.Ref = nodeGetRef(node)
		ref.Summary = nodeGetString(node, "summary")
		ref.Description = nodeGetString(node, "description")
		return ref, nil
	}

	response, err := parseSharedResponse(node, ctx)
	if err != nil {
		return nil, err
	}
	ref.SetValue(response)

	return ref, nil
}
