package openapi30x

import (
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// parseHeaderRef parses a HeaderRef from a yaml.Node.
func parseHeaderRef(node *yaml.Node, ctx *ParseContext) (*openapi30models.HeaderRef, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "header must be an object")
	}

	ref := &openapi30models.HeaderRef{}
	ref.NodeSource = ctx.nodeSource(node)
	ref.Extensions = parseNodeExtensions(node)

	// Check for $ref
	if nodeHasRef(node) {
		ref.Ref = nodeGetRef(node)
		return ref, nil
	}

	// Parse inline header
	header, err := parseSharedHeader(node, ctx)
	if err != nil {
		return nil, err
	}
	ref.Value = header

	return ref, nil
}
