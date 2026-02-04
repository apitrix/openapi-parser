package openapi30

import (
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// ParseServer parses the Link.Server field.
func (p *linkParser) ParseServer(parent *yaml.Node, c *ParseContext) (*openapi30models.Server, error) {
	node := nodeGetValue(parent, "server")
	if node == nil {
		return nil, nil
	}
	return parseSharedServer(node, c.Push("server"))
}
