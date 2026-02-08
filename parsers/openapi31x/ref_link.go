package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// parseLinkRef parses a LinkRef from a yaml.Node.
func parseLinkRef(node *yaml.Node, ctx *ParseContext) (*openapi31models.LinkRef, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "link must be an object")
	}

	ref := &openapi31models.LinkRef{}
	ref.Trix.Source = ctx.nodeSource(node)
	ref.VendorExtensions = parseNodeExtensions(node)

	if nodeHasRef(node) {
		ref.Ref = nodeGetRef(node)
		ref.Summary = nodeGetString(node, "summary")
		ref.Description = nodeGetString(node, "description")
		return ref, nil
	}

	link, err := parseSharedLink(node, ctx)
	if err != nil {
		return nil, err
	}
	ref.Value = link

	return ref, nil
}
