package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// ParseDependentSchemas parses the Schema.DependentSchemas field.
// JSON Schema 2020-12: schema-based conditional dependencies
func (p *schemaParser) ParseDependentSchemas(parent *yaml.Node, c *ParseContext) (map[string]*openapi31models.SchemaRef, error) {
	node := nodeGetValue(parent, "dependentSchemas")
	if node == nil || !nodeIsMapping(node) {
		return nil, nil
	}

	schemas := make(map[string]*openapi31models.SchemaRef)
	dctx := c.Push("dependentSchemas")
	for name, schemaNode := range nodeMapPairs(node) {
		ref, err := parseSchemaRef(schemaNode, dctx.push(name))
		if err != nil {
			return nil, err
		}
		schemas[name] = ref
	}
	return schemas, nil
}
