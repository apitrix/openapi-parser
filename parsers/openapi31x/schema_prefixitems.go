package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// ParsePrefixItems parses the Schema.PrefixItems field.
// JSON Schema 2020-12: tuple validation (replaces items-as-array from older drafts)
func (p *schemaParser) ParsePrefixItems(parent *yaml.Node, c *ParseContext) ([]*shared.RefWithMeta[openapi31models.Schema], error) {
	node := nodeGetValue(parent, "prefixItems")
	if node == nil || !nodeIsSequence(node) {
		return nil, nil
	}

	refs := make([]*shared.RefWithMeta[openapi31models.Schema], 0, len(node.Content))
	pctx := c.Push("prefixItems")
	for i, itemNode := range node.Content {
		ref, err := parseSchemaRef(itemNode, pctx.push(itoa(i)))
		if err != nil {
			return nil, err
		}
		refs = append(refs, ref)
	}
	return refs, nil
}
