package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// ParseNot parses the Schema.Not field.
// Complex property: recursive SchemaRef
func (p *schemaParser) ParseNot(parent *yaml.Node, c *ParseContext) (*openapi31models.SchemaRef, error) {
	node := nodeGetValue(parent, "not")
	if node == nil {
		return nil, nil
	}
	return parseSchemaRef(node, c.Push("not"))
}
