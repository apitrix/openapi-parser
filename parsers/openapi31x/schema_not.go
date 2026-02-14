package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// ParseNot parses the Schema.Not field.
// Complex property: recursive SchemaRef
func (p *schemaParser) ParseNot(parent *yaml.Node, c *ParseContext) (*shared.RefWithMeta[openapi31models.Schema], error) {
	node := nodeGetValue(parent, "not")
	if node == nil {
		return nil, nil
	}
	return parseSchemaRef(node, c.Push("not"))
}
