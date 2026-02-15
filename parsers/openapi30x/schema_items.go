package openapi30x

import (
	openapi30models "github.com/apitrix/openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// ParseItems parses the Schema.Items field.
// Complex property: recursive SchemaRef
func (p *schemaParser) ParseItems(parent *yaml.Node, c *ParseContext) (*openapi30models.RefSchema, error) {
	node := nodeGetValue(parent, "items")
	if node == nil {
		return nil, nil
	}
	return parseSchemaRef(node, c.Push("items"))
}
