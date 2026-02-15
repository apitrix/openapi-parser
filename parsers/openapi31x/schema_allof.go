package openapi31x

import (
	openapi31models "github.com/apitrix/openapi-parser/models/openapi31"
	"github.com/apitrix/openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// ParseAllOf parses the Schema.AllOf field.
// Complex property: array of SchemaRef
func (p *schemaParser) ParseAllOf(parent *yaml.Node, c *ParseContext) ([]*shared.RefWithMeta[openapi31models.Schema], error) {
	node := nodeGetValue(parent, "allOf")
	if node == nil || !nodeIsSequence(node) {
		return nil, nil
	}

	refs := make([]*shared.RefWithMeta[openapi31models.Schema], 0, len(node.Content))
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
