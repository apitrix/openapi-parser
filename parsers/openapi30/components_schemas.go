package openapi30

import (
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// parseComponentsSchemas parses the Components.Schemas field.
func parseComponentsSchemas(parent *yaml.Node, ctx *ParseContext) (map[string]*openapi30models.SchemaRef, error) {
	node := nodeGetValue(parent, "schemas")
	if node == nil || !nodeIsMapping(node) {
		return nil, nil
	}

	schemas := make(map[string]*openapi30models.SchemaRef)
	sctx := ctx.push("schemas")
	for _, name := range nodeKeys(node) {
		schemaNode := nodeGetValue(node, name)
		schemaRef, err := parseSchemaRef(schemaNode, sctx.push(name))
		if err != nil {
			return nil, err
		}
		schemas[name] = schemaRef
	}
	return schemas, nil
}
