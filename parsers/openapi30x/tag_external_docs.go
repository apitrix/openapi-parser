package openapi30x

import (
	openapi30models "github.com/apitrix/openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// ParseExternalDocs parses the Tag.ExternalDocs field.
func (p *tagParser) ParseExternalDocs(parent *yaml.Node, c *ParseContext) (*openapi30models.ExternalDocumentation, error) {
	node := nodeGetValue(parent, "externalDocs")
	if node == nil {
		return nil, nil
	}
	return parseSharedExternalDocs(node, c.Push("externalDocs"))
}
