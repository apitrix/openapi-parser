package openapi30x

import (
	"openapi-parser/models/shared"
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// parseRequestBodyRef parses a RequestBodyRef from a yaml.Node.
func parseRequestBodyRef(node *yaml.Node, ctx *ParseContext) (*shared.Ref[openapi30models.RequestBody], error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "requestBody must be an object")
	}

	ref := &shared.Ref[openapi30models.RequestBody]{}
	ref.Trix.Source = ctx.nodeSource(node)
	ref.VendorExtensions = parseNodeExtensions(node)

	// Check for $ref
	if nodeHasRef(node) {
		ref.Ref = nodeGetRef(node)
		return ref, nil
	}

	// Parse inline request body
	requestBody, err := parseSharedRequestBody(node, ctx)
	if err != nil {
		return nil, err
	}
	ref.SetValue(requestBody)

	return ref, nil
}
