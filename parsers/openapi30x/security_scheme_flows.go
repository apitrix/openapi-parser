package openapi30x

import (
	openapi30models "github.com/apitrix/openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// ParseFlows parses the SecurityScheme.Flows field.
func (p *securitySchemeParser) ParseFlows(parent *yaml.Node, c *ParseContext) (*openapi30models.OAuthFlows, error) {
	node := nodeGetValue(parent, "flows")
	if node == nil {
		return nil, nil
	}
	return parseSharedOAuthFlows(node, c.Push("flows"))
}
