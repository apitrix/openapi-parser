package openapi30x

import (
	"openapi-parser/models/shared"
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// parseLinkRef parses a LinkRef from a yaml.Node.
func parseLinkRef(node *yaml.Node, ctx *ParseContext) (*shared.Ref[openapi30models.Link], error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "link must be an object")
	}

	ref := &shared.Ref[openapi30models.Link]{}
	ref.Trix.Source = ctx.nodeSource(node)
	ref.VendorExtensions = parseNodeExtensions(node)

	// Check for $ref
	if nodeHasRef(node) {
		ref.Ref = nodeGetRef(node)
		return ref, nil
	}

	// Parse inline link
	link, err := parseSharedLink(node, ctx)
	if err != nil {
		return nil, err
	}
	ref.SetValue(link)

	return ref, nil
}
