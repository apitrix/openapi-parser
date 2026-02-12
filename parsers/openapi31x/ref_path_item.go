package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// parsePathItemRef parses a PathItemRef from a yaml.Node.
func parsePathItemRef(node *yaml.Node, ctx *ParseContext) (*openapi31models.PathItemRef, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "pathItem must be an object")
	}

	ref := &openapi31models.PathItemRef{}
	ref.Trix.Source = ctx.nodeSource(node)
	ref.VendorExtensions = parseNodeExtensions(node)

	if nodeHasRef(node) {
		ref.Ref = nodeGetRef(node)
		ref.Summary = nodeGetString(node, "summary")
		ref.Description = nodeGetString(node, "description")
		return ref, nil
	}

	pathItem, err := parseOpenAPIPathsPathItem(node, ctx)
	if err != nil {
		return nil, err
	}
	ref.SetValue(pathItem)

	return ref, nil
}
