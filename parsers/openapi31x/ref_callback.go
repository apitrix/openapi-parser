package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// parseCallbackRef parses a CallbackRef from a yaml.Node.
func parseCallbackRef(node *yaml.Node, ctx *ParseContext) (*shared.RefWithMeta[openapi31models.Callback], error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "callback must be an object")
	}

	ref := &shared.RefWithMeta[openapi31models.Callback]{}
	ref.Trix.Source = ctx.nodeSource(node)
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
	ref.SetValue(cb)

	return ref, nil
}
