package openapi31x

import (
	openapi31models "github.com/apitrix/openapi-parser/models/openapi31"
	"github.com/apitrix/openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// parsePathItemParameters parses the PathItem.Parameters field.
func parsePathItemParameters(parent *yaml.Node, ctx *ParseContext) ([]*shared.RefWithMeta[openapi31models.Parameter], error) {
	node := nodeGetValue(parent, "parameters")
	if node == nil || !nodeIsSequence(node) {
		return nil, nil
	}

	params := make([]*shared.RefWithMeta[openapi31models.Parameter], 0, len(node.Content))
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
