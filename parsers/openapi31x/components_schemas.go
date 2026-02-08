package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// parseComponentsSchemas parses the Components.Schemas field.
func parseComponentsSchemas(parent *yaml.Node, ctx *ParseContext) (map[string]*openapi31models.SchemaRef, error) {
	node := nodeGetValue(parent, "schemas")
	if node == nil || !nodeIsMapping(node) {
		return nil, nil
	}

	schemas := make(map[string]*openapi31models.SchemaRef)
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
