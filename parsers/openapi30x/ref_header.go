package openapi30x

import (
	"openapi-parser/models/shared"
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// parseHeaderRef parses a HeaderRef from a yaml.Node.
func parseHeaderRef(node *yaml.Node, ctx *ParseContext) (*shared.Ref[openapi30models.Header], error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "header must be an object")
	}

	ref := &shared.Ref[openapi30models.Header]{}
	ref.Trix.Source = ctx.nodeSource(node)
	ref.VendorExtensions = parseNodeExtensions(node)

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
	ref.SetValue(header)

	return ref, nil
}
