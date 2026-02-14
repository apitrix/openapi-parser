package openapi30x

import (
	"openapi-parser/models/shared"
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// parseOperationParameters parses the Operation.Parameters field.
func parseOperationParameters(parent *yaml.Node, ctx *ParseContext) ([]*shared.Ref[openapi30models.Parameter], error) {
	node := nodeGetValue(parent, "parameters")
	if node == nil || !nodeIsSequence(node) {
		return nil, nil
	}

	params := make([]*shared.Ref[openapi30models.Parameter], 0, len(node.Content))
	pctx := ctx.push("parameters")
	for i, paramNode := range node.Content {
		paramRef, err := parseParameterRef(paramNode, pctx.push(itoa(i)))
		if err != nil {
			return nil, err
		}
		params = append(params, paramRef)
	}
	return params, nil
}
