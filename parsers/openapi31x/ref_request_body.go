package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// parseRequestBodyRef parses a RequestBodyRef from a yaml.Node.
func parseRequestBodyRef(node *yaml.Node, ctx *ParseContext) (*openapi31models.RequestBodyRef, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "requestBody must be an object")
	}

	ref := &openapi31models.RequestBodyRef{}
	ref.NodeSource = ctx.nodeSource(node)
	ref.VendorExtensions = parseNodeExtensions(node)

	if nodeHasRef(node) {
		ref.Ref = nodeGetRef(node)
		ref.Summary = nodeGetString(node, "summary")
		ref.Description = nodeGetString(node, "description")
		return ref, nil
	}

	body, err := parseSharedRequestBody(node, ctx)
	if err != nil {
		return nil, err
	}
	ref.Value = body

	return ref, nil
}
