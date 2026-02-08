package openapi30x

import (
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// ParseProperties parses the Schema.Properties field.
// Complex property: map of SchemaRef
func (p *schemaParser) ParseProperties(parent *yaml.Node, c *ParseContext) (map[string]*openapi30models.SchemaRef, error) {
	node := nodeGetValue(parent, "properties")
	if node == nil || !nodeIsMapping(node) {
		return nil, nil
	}

	props := make(map[string]*openapi30models.SchemaRef)
	pctx := c.Push("properties")
	for name, propNode := range nodeMapPairs(node) {
		ref, err := parseSchemaRef(propNode, pctx.push(name))
		if err != nil {
			return nil, err
		}
		props[name] = ref
	}
	return props, nil
}
