package openapi30x

import (
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// ParseSchema parses the MediaType.Schema field.
func (p *mediaTypeParser) ParseSchema(parent *yaml.Node, c *ParseContext) (*openapi30models.SchemaRef, error) {
	node := nodeGetValue(parent, "schema")
	if node == nil {
		return nil, nil
	}
	return parseSchemaRef(node, c.Push("schema"))
}
