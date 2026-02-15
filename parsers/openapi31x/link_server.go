package openapi31x

import (
	openapi31models "github.com/apitrix/openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// ParseServer parses the Link.Server field.
func (p *linkParser) ParseServer(parent *yaml.Node, c *ParseContext) (*openapi31models.Server, error) {
	node := nodeGetValue(parent, "server")
	if node == nil {
		return nil, nil
	}
	return parseSharedServer(node, c.Push("server"))
}
