package openapi30x

import (
	"openapi-parser/models/shared"
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// ParseOneOf parses the Schema.OneOf field.
// Complex property: array of SchemaRef
func (p *schemaParser) ParseOneOf(parent *yaml.Node, c *ParseContext) ([]*shared.Ref[openapi30models.Schema], error) {
	node := nodeGetValue(parent, "oneOf")
	if node == nil || !nodeIsSequence(node) {
		return nil, nil
	}

	refs := make([]*shared.Ref[openapi30models.Schema], 0, len(node.Content))
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
