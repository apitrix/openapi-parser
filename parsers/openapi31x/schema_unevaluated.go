package openapi31x

import (
	"openapi-parser/models/shared"
	openapi31models "openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// ParseUnevaluatedItems parses the Schema.UnevaluatedItems field.
// JSON Schema 2020-12: applies to items not covered by items/prefixItems/contains
func (p *schemaParser) ParseUnevaluatedItems(parent *yaml.Node, c *ParseContext) (*shared.RefWithMeta[openapi31models.Schema], error) {
	node := nodeGetValue(parent, "unevaluatedItems")
	if node == nil {
		return nil, nil
	}
	return parseSchemaRef(node, c.Push("unevaluatedItems"))
}

// ParseUnevaluatedProperties parses the Schema.UnevaluatedProperties field.
// JSON Schema 2020-12: applies to properties not covered by properties/patternProperties/additionalProperties
func (p *schemaParser) ParseUnevaluatedProperties(parent *yaml.Node, c *ParseContext) (*shared.RefWithMeta[openapi31models.Schema], error) {
	node := nodeGetValue(parent, "unevaluatedProperties")
	if node == nil {
		return nil, nil
	}
	return parseSchemaRef(node, c.Push("unevaluatedProperties"))
}
