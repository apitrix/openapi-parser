package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// ParseExternalDocs parses the Tag.ExternalDocs field.
func (p *tagParser) ParseExternalDocs(parent *yaml.Node, c *ParseContext) (*openapi31models.ExternalDocumentation, error) {
	node := nodeGetValue(parent, "externalDocs")
	if node == nil {
		return nil, nil
	}
	return parseSharedExternalDocs(node, c.Push("externalDocs"))
}
