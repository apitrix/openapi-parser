package openapi30x

import (
	"openapi-parser/models/shared"
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// ParseAllOf parses the Schema.AllOf field.
// Complex property: array of SchemaRef
func (p *schemaParser) ParseAllOf(parent *yaml.Node, c *ParseContext) ([]*shared.Ref[openapi30models.Schema], error) {
	node := nodeGetValue(parent, "allOf")
	if node == nil || !nodeIsSequence(node) {
		return nil, nil
	}

	refs := make([]*shared.Ref[openapi30models.Schema], 0, len(node.Content))
	actx := c.Push("allOf")
	for i, itemNode := range node.Content {
		ref, err := parseSchemaRef(itemNode, actx.push(itoa(i)))
		if err != nil {
			return nil, err
		}
		refs = append(refs, ref)
	}
	return refs, nil
}
