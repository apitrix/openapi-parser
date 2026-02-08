package openapi30x

import (
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// parsePathItemRef parses a PathItemRef from a yaml.Node.
func parsePathItemRef(node *yaml.Node, ctx *ParseContext) (*openapi30models.PathItemRef, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "pathItem must be an object")
	}

	ref := &openapi30models.PathItemRef{}
	ref.Trix.Source = ctx.nodeSource(node)
	ref.VendorExtensions = parseNodeExtensions(node)

	// Check for $ref
	if nodeHasRef(node) {
		ref.Ref = nodeGetRef(node)
		return ref, nil
	}

	// Parse inline path item
	pathItem, err := parseOpenAPIPathsPathItem(node, ctx)
	if err != nil {
		return nil, err
	}
	ref.Value = pathItem

	return ref, nil
}
