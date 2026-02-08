package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// parseCallbackRef parses a CallbackRef from a yaml.Node.
func parseCallbackRef(node *yaml.Node, ctx *ParseContext) (*openapi31models.CallbackRef, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "callback must be an object")
	}

	ref := &openapi31models.CallbackRef{}
	ref.NodeSource = ctx.nodeSource(node)
	ref.VendorExtensions = parseNodeExtensions(node)

	if nodeHasRef(node) {
		ref.Ref = nodeGetRef(node)
		ref.Summary = nodeGetString(node, "summary")
		ref.Description = nodeGetString(node, "description")
		return ref, nil
	}

	cb, err := parseSharedCallback(node, ctx)
	if err != nil {
		return nil, err
	}
	ref.Value = cb

	return ref, nil
}
