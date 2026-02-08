package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// ParseAnyOf parses the Schema.AnyOf field.
// Complex property: array of SchemaRef
func (p *schemaParser) ParseAnyOf(parent *yaml.Node, c *ParseContext) ([]*openapi31models.SchemaRef, error) {
	node := nodeGetValue(parent, "anyOf")
	if node == nil || !nodeIsSequence(node) {
		return nil, nil
	}

	refs := make([]*openapi31models.SchemaRef, 0, len(node.Content))
	actx := c.Push("anyOf")
	for i, itemNode := range node.Content {
		ref, err := parseSchemaRef(itemNode, actx.push(itoa(i)))
		if err != nil {
			return nil, err
		}
		refs = append(refs, ref)
	}
	return refs, nil
}
