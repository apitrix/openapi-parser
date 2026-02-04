package openapi30

import (
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// parsePathItemParameters parses the PathItem.Parameters field.
func parsePathItemParameters(parent *yaml.Node, ctx *ParseContext) ([]*openapi30models.ParameterRef, error) {
	node := nodeGetValue(parent, "parameters")
	if node == nil || !nodeIsSequence(node) {
		return nil, nil
	}

	params := make([]*openapi30models.ParameterRef, 0, len(node.Content))
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
