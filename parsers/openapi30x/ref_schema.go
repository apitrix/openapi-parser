package openapi30x

import (
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// parseSchemaRef parses a SchemaRef from a yaml.Node.
func parseSchemaRef(node *yaml.Node, ctx *ParseContext) (*openapi30models.RefSchema, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "schema must be an object")
	}

	ref := &openapi30models.RefSchema{}
	ref.Trix.Source = ctx.nodeSource(node)
	ref.VendorExtensions = parseNodeExtensions(node)

	// Check for $ref
	if nodeHasRef(node) {
		ref.Ref = nodeGetRef(node)
		return ref, nil
	}

	// Parse inline schema
	schema, err := parseSharedSchema(node, ctx)
	if err != nil {
		return nil, err
	}
	ref.SetValue(schema)

	return ref, nil
}
