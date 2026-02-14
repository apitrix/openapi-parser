package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// ParseOneOf parses the Schema.OneOf field.
// Complex property: array of SchemaRef
func (p *schemaParser) ParseOneOf(parent *yaml.Node, c *ParseContext) ([]*shared.RefWithMeta[openapi31models.Schema], error) {
	node := nodeGetValue(parent, "oneOf")
	if node == nil || !nodeIsSequence(node) {
		return nil, nil
	}

	refs := make([]*shared.RefWithMeta[openapi31models.Schema], 0, len(node.Content))
	octx := c.Push("oneOf")
	for i, itemNode := range node.Content {
		ref, err := parseSchemaRef(itemNode, octx.push(itoa(i)))
		if err != nil {
			return nil, err
		}
		refs = append(refs, ref)
	}
	return refs, nil
}
