package openapi31x

import (
	"openapi-parser/models/shared"
	openapi31models "openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// ParseSchema parses the MediaType.Schema field.
func (p *mediaTypeParser) ParseSchema(parent *yaml.Node, c *ParseContext) (*shared.RefWithMeta[openapi31models.Schema], error) {
	node := nodeGetValue(parent, "schema")
	if node == nil {
		return nil, nil
	}
	return parseSchemaRef(node, c.Push("schema"))
}
