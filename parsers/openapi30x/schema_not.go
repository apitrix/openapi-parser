package openapi30x

import (
	openapi30models "github.com/apitrix/openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// ParseNot parses the Schema.Not field.
// Complex property: recursive SchemaRef
func (p *schemaParser) ParseNot(parent *yaml.Node, c *ParseContext) (*openapi30models.RefSchema, error) {
	node := nodeGetValue(parent, "not")
	if node == nil {
		return nil, nil
	}
	return parseSchemaRef(node, c.Push("not"))
}
