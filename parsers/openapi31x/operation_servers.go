package openapi31x

import (
	openapi31models "github.com/apitrix/openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// parseOperationServers parses the Operation.Servers field.
func parseOperationServers(parent *yaml.Node, ctx *ParseContext) ([]*openapi31models.Server, error) {
	node := nodeGetValue(parent, "servers")
	if node == nil {
		return nil, nil
	}
	return parseSharedServers(node, ctx.push("servers"))
}
