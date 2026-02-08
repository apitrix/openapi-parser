package openapi30x

import (
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// parseExampleRef parses an ExampleRef from a yaml.Node.
func parseExampleRef(node *yaml.Node, ctx *ParseContext) (*openapi30models.ExampleRef, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "example must be an object")
	}

	ref := &openapi30models.ExampleRef{}
	ref.NodeSource = ctx.nodeSource(node)
	ref.VendorExtensions = parseNodeExtensions(node)

	// Check for $ref
	if nodeHasRef(node) {
		ref.Ref = nodeGetRef(node)
		return ref, nil
	}

	// Parse inline example
	example, err := parseSharedExample(node, ctx)
	if err != nil {
		return nil, err
	}
	ref.Value = example

	return ref, nil
}
