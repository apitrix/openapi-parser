package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// parseSchemaRef parses a SchemaRef from a yaml.Node.
// In 3.1, $ref can have summary and description alongside it.
func parseSchemaRef(node *yaml.Node, ctx *ParseContext) (*shared.RefWithMeta[openapi31models.Schema], error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "schema must be an object")
	}

	ref := &shared.RefWithMeta[openapi31models.Schema]{}
	ref.Trix.Source = ctx.nodeSource(node)
	ref.VendorExtensions = parseNodeExtensions(node)

	// Check for $ref
	if nodeHasRef(node) {
		ref.Ref = nodeGetRef(node)
		ref.Summary = nodeGetString(node, "summary")
		ref.Description = nodeGetString(node, "description")
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
