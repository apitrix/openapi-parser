package openapi31x

import (
	openapi31models "github.com/apitrix/openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// ParseFlows parses the SecurityScheme.Flows field.
func (p *securitySchemeParser) ParseFlows(parent *yaml.Node, c *ParseContext) (*openapi31models.OAuthFlows, error) {
	node := nodeGetValue(parent, "flows")
	if node == nil {
		return nil, nil
	}
	return parseSharedOAuthFlows(node, c.Push("flows"))
}
