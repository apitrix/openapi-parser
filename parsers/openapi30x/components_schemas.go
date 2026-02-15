package openapi30x

import (
	openapi30models "github.com/apitrix/openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// parseComponentsSchemas parses the Components.Schemas field.
func parseComponentsSchemas(parent *yaml.Node, ctx *ParseContext) (map[string]*openapi30models.RefSchema, error) {
	node := nodeGetValue(parent, "schemas")
	if node == nil || !nodeIsMapping(node) {
		return nil, nil
	}

	schemas := make(map[string]*openapi30models.RefSchema)
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
