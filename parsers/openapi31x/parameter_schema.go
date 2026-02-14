package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// ParseSchema parses the Parameter.Schema field.
func (p *parameterParser) ParseSchema(parent *yaml.Node, c *ParseContext) (*shared.RefWithMeta[openapi31models.Schema], error) {
	node := nodeGetValue(parent, "schema")
	if node == nil {
		return nil, nil
	}
	return parseSchemaRef(node, c.Push("schema"))
}
