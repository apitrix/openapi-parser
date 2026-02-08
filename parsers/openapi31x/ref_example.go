package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// parseExampleRef parses an ExampleRef from a yaml.Node.
func parseExampleRef(node *yaml.Node, ctx *ParseContext) (*openapi31models.ExampleRef, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "example must be an object")
	}

	ref := &openapi31models.ExampleRef{}
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
	ref.Value = example

	return ref, nil
}
