package openapi31x

import (
	openapi31models "github.com/apitrix/openapi-parser/models/openapi31"
	"github.com/apitrix/openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// ParseDependentSchemas parses the Schema.DependentSchemas field.
// JSON Schema 2020-12: schema-based conditional dependencies
func (p *schemaParser) ParseDependentSchemas(parent *yaml.Node, c *ParseContext) (map[string]*shared.RefWithMeta[openapi31models.Schema], error) {
	node := nodeGetValue(parent, "dependentSchemas")
	if node == nil || !nodeIsMapping(node) {
		return nil, nil
	}

	schemas := make(map[string]*shared.RefWithMeta[openapi31models.Schema])
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
