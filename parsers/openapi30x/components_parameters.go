package openapi30x

import (
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// parseComponentsParameters parses the Components.Parameters field.
func parseComponentsParameters(parent *yaml.Node, ctx *ParseContext) (map[string]*openapi30models.ParameterRef, error) {
	node := nodeGetValue(parent, "parameters")
	if node == nil || !nodeIsMapping(node) {
		return nil, nil
	}

	params := make(map[string]*openapi30models.ParameterRef)
	pctx := ctx.push("parameters")
	for name, paramNode := range nodeMapPairs(node) {
		paramRef, err := parseParameterRef(paramNode, pctx.push(name))
		if err != nil {
			return nil, err
		}
		params[name] = paramRef
	}
	return params, nil
}
