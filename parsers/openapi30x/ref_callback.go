package openapi30x

import (
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// parseCallbackRef parses a CallbackRef from a yaml.Node.
func parseCallbackRef(node *yaml.Node, ctx *ParseContext) (*openapi30models.CallbackRef, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "callback must be an object")
	}

	ref := &openapi30models.CallbackRef{}
	ref.NodeSource = ctx.nodeSource(node)
	ref.Extensions = parseNodeExtensions(node)

	// Check for $ref
	if nodeHasRef(node) {
		ref.Ref = nodeGetRef(node)
		return ref, nil
	}

	// Parse inline callback
	callback, err := parseSharedCallback(node, ctx)
	if err != nil {
		return nil, err
	}
	ref.Value = callback

	return ref, nil
}
