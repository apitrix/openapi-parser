package openapi30x

import (
	"openapi-parser/models/shared"
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// parseResponseRef parses a ResponseRef from a yaml.Node.
func parseResponseRef(node *yaml.Node, ctx *ParseContext) (*shared.Ref[openapi30models.Response], error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "response must be an object")
	}

	ref := &shared.Ref[openapi30models.Response]{}
	ref.Trix.Source = ctx.nodeSource(node)
	ref.VendorExtensions = parseNodeExtensions(node)

	// Check for $ref
	if nodeHasRef(node) {
		ref.Ref = nodeGetRef(node)
		return ref, nil
	}

	// Parse inline response
	response, err := parseSharedResponse(node, ctx)
	if err != nil {
		return nil, err
	}
	ref.SetValue(response)

	return ref, nil
}
