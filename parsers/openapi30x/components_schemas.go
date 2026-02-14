package openapi30x

import (
	"openapi-parser/models/shared"
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// parseComponentsSchemas parses the Components.Schemas field.
func parseComponentsSchemas(parent *yaml.Node, ctx *ParseContext) (map[string]*shared.Ref[openapi30models.Schema], error) {
	node := nodeGetValue(parent, "schemas")
	if node == nil || !nodeIsMapping(node) {
		return nil, nil
	}

	schemas := make(map[string]*shared.Ref[openapi30models.Schema])
	sctx := ctx.push("schemas")
	for name, schemaNode := range nodeMapPairs(node) {
		schemaRef, err := parseSchemaRef(schemaNode, sctx.push(name))
		if err != nil {
			return nil, err
		}
		schemas[name] = schemaRef
	}
	return schemas, nil
}
