package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// parseExampleRef parses an ExampleRef from a yaml.Node.
func parseExampleRef(node *yaml.Node, ctx *ParseContext) (*shared.RefWithMeta[openapi31models.Example], error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "example must be an object")
	}

	ref := &shared.RefWithMeta[openapi31models.Example]{}
	ref.Trix.Source = ctx.nodeSource(node)
	ref.VendorExtensions = parseNodeExtensions(node)

	if nodeHasRef(node) {
		ref.Ref = nodeGetRef(node)
		ref.Summary = nodeGetString(node, "summary")
		ref.Description = nodeGetString(node, "description")
		return ref, nil
	}

	example, err := parseSharedExample(node, ctx)
	if err != nil {
		return nil, err
	}
	ref.SetValue(example)

	return ref, nil
}
