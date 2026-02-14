package openapi30x

import (
	"openapi-parser/models/shared"
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// ParseSchema parses the Header.Schema field.
func (p *headerParser) ParseSchema(parent *yaml.Node, c *ParseContext) (*shared.Ref[openapi30models.Schema], error) {
	node := nodeGetValue(parent, "schema")
	if node == nil {
		return nil, nil
	}
	return parseSchemaRef(node, c.Push("schema"))
}
