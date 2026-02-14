package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// parseComponentsParameters parses the Components.Parameters field.
func parseComponentsParameters(parent *yaml.Node, ctx *ParseContext) (map[string]*shared.RefWithMeta[openapi31models.Parameter], error) {
	node := nodeGetValue(parent, "parameters")
	if node == nil || !nodeIsMapping(node) {
		return nil, nil
	}

	params := make(map[string]*shared.RefWithMeta[openapi31models.Parameter])
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
