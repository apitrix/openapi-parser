package openapi31x

import (
	openapi31models "github.com/apitrix/openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// parseOperationResponses parses the Operation.Responses field.
func parseOperationResponses(parent *yaml.Node, ctx *ParseContext) (*openapi31models.Responses, error) {
	node := nodeGetValue(parent, "responses")
	if node == nil {
		return nil, nil
	}
	return parseSharedResponses(node, ctx.push("responses"))
}
