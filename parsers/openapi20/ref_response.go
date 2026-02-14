package openapi20

import (
	"openapi-parser/models/shared"
	openapi20models "openapi-parser/models/openapi20"

	"gopkg.in/yaml.v3"
)

// parseResponseRef parses a ResponseRef (either $ref or inline response) from a yaml.Node.
func parseResponseRef(node *yaml.Node, ctx *ParseContext) (*shared.Ref[openapi20models.Response], error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "response must be an object")
	}

	ref := &shared.Ref[openapi20models.Response]{}

	// Check if it's a reference
	if nodeHasRef(node) {
		ref.Ref = nodeGetRef(node)
		ref.Trix.Source = ctx.nodeSource(node)
		return ref, nil
	}

	// Parse inline response
	response, err := parseResponse(node, ctx)
	if err != nil {
		return nil, err
	}
	ref.SetValue(response)
	ref.Trix.Source = ctx.nodeSource(node)

	return ref, nil
}
