package openapi30x

import (
	openapi30models "github.com/apitrix/openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// parseOperationResponses parses the Operation.Responses field.
func parseOperationResponses(parent *yaml.Node, ctx *ParseContext) (*openapi30models.Responses, error) {
	node := nodeGetValue(parent, "responses")
	if node == nil {
		return nil, nil
	}
	return parseSharedResponses(node, ctx.push("responses"))
}
