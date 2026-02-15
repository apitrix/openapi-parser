package openapi30x

import (
	openapi30models "github.com/apitrix/openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// parseRequestBodyRef parses a RequestBodyRef from a yaml.Node.
func parseRequestBodyRef(node *yaml.Node, ctx *ParseContext) (*openapi30models.RefRequestBody, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "requestBody must be an object")
	}

	ref := &openapi30models.RefRequestBody{}
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
