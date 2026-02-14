package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// parseHeaderRef parses a HeaderRef from a yaml.Node.
func parseHeaderRef(node *yaml.Node, ctx *ParseContext) (*shared.RefWithMeta[openapi31models.Header], error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "header must be an object")
	}

	ref := &shared.RefWithMeta[openapi31models.Header]{}
	ref.Trix.Source = ctx.nodeSource(node)
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
	ref.SetValue(header)

	return ref, nil
}
