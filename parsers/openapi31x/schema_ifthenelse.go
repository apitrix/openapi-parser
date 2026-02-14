package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// ParseIf parses the Schema.If field.
// JSON Schema 2020-12: conditional subschema
func (p *schemaParser) ParseIf(parent *yaml.Node, c *ParseContext) (*shared.RefWithMeta[openapi31models.Schema], error) {
	node := nodeGetValue(parent, "if")
	if node == nil {
		return nil, nil
	}
	return parseSchemaRef(node, c.Push("if"))
}

// ParseThen parses the Schema.Then field.
// JSON Schema 2020-12: conditional subschema
func (p *schemaParser) ParseThen(parent *yaml.Node, c *ParseContext) (*shared.RefWithMeta[openapi31models.Schema], error) {
	node := nodeGetValue(parent, "then")
	if node == nil {
		return nil, nil
	}
	return parseSchemaRef(node, c.Push("then"))
}

// ParseElse parses the Schema.Else field.
// JSON Schema 2020-12: conditional subschema
func (p *schemaParser) ParseElse(parent *yaml.Node, c *ParseContext) (*shared.RefWithMeta[openapi31models.Schema], error) {
	node := nodeGetValue(parent, "else")
	if node == nil {
		return nil, nil
	}
	return parseSchemaRef(node, c.Push("else"))
}
