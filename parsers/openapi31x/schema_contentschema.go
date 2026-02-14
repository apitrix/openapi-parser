package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// ParseContentSchema parses the Schema.ContentSchema field.
// JSON Schema 2020-12: schema for content described by contentMediaType/contentEncoding
func (p *schemaParser) ParseContentSchema(parent *yaml.Node, c *ParseContext) (*shared.RefWithMeta[openapi31models.Schema], error) {
	node := nodeGetValue(parent, "contentSchema")
	if node == nil {
		return nil, nil
	}
	return parseSchemaRef(node, c.Push("contentSchema"))
}
