package openapi30

import (
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// parseOperationServers parses the Operation.Servers field.
func parseOperationServers(parent *yaml.Node, ctx *ParseContext) ([]*openapi30models.Server, error) {
	node := nodeGetValue(parent, "servers")
	if node == nil {
		return nil, nil
	}
	return parseSharedServers(node, ctx.push("servers"))
}
