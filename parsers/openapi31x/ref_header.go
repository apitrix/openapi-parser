package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// parseHeaderRef parses a HeaderRef from a yaml.Node.
func parseHeaderRef(node *yaml.Node, ctx *ParseContext) (*openapi31models.HeaderRef, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "header must be an object")
	}

	ref := &openapi31models.HeaderRef{}
	ref.NodeSource = ctx.nodeSource(node)
	ref.VendorExtensions = parseNodeExtensions(node)

	if nodeHasRef(node) {
		ref.Ref = nodeGetRef(node)
		ref.Summary = nodeGetString(node, "summary")
		ref.Description = nodeGetString(node, "description")
		return ref, nil
	}

	header, err := parseSharedHeader(node, ctx)
	if err != nil {
		return nil, err
	}
	ref.Value = header

	return ref, nil
}
