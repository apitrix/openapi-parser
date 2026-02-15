package openapi20

import (
	openapi20models "github.com/apitrix/openapi-parser/models/openapi20"

	"gopkg.in/yaml.v3"
)

// parseSchemaRef parses a SchemaRef (either $ref or inline schema) from a yaml.Node.
func parseSchemaRef(node *yaml.Node, ctx *ParseContext) (*openapi20models.RefSchema, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "schema must be an object")
	}

	ref := &openapi20models.RefSchema{}

	// Check if it's a reference
	if nodeHasRef(node) {
		ref.Ref = nodeGetRef(node)
		ref.Trix.Source = ctx.nodeSource(node)
		return ref, nil
	}

	// Parse inline schema
	schema, err := parseSchema(node, ctx)
	if err != nil {
		return nil, err
	}
	ref.SetValue(schema)
	ref.Trix.Source = ctx.nodeSource(node)

	return ref, nil
}
