package openapi30x

import (
	openapi30models "github.com/apitrix/openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// parsePathItemServers parses the PathItem.Servers field.
func parsePathItemServers(parent *yaml.Node, ctx *ParseContext) ([]*openapi30models.Server, error) {
	node := nodeGetValue(parent, "servers")
	if node == nil {
		return nil, nil
	}
	return parseSharedServers(node, ctx.push("servers"))
}
