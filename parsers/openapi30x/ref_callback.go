package openapi30x

import (
	"openapi-parser/models/shared"
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// parseCallbackRef parses a CallbackRef from a yaml.Node.
func parseCallbackRef(node *yaml.Node, ctx *ParseContext) (*shared.Ref[openapi30models.Callback], error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "callback must be an object")
	}

	ref := &shared.Ref[openapi30models.Callback]{}
	ref.Trix.Source = ctx.nodeSource(node)
	ref.VendorExtensions = parseNodeExtensions(node)

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
	ref.SetValue(callback)

	return ref, nil
}
