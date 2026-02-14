package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// ParseItems parses the Schema.Items field.
// Complex property: recursive SchemaRef
func (p *schemaParser) ParseItems(parent *yaml.Node, c *ParseContext) (*shared.RefWithMeta[openapi31models.Schema], error) {
	node := nodeGetValue(parent, "items")
	if node == nil {
		return nil, nil
	}
	return parseSchemaRef(node, c.Push("items"))
}
