package openapi30x

import (
	"openapi-parser/models/shared"
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// parseComponentsResponses parses the Components.Responses field.
func parseComponentsResponses(parent *yaml.Node, ctx *ParseContext) (map[string]*shared.Ref[openapi30models.Response], error) {
	node := nodeGetValue(parent, "responses")
	if node == nil || !nodeIsMapping(node) {
		return nil, nil
	}

	responses := make(map[string]*shared.Ref[openapi30models.Response])
	rctx := ctx.push("responses")
	for name, respNode := range nodeMapPairs(node) {
		respRef, err := parseResponseRef(respNode, rctx.push(name))
		if err != nil {
			return nil, err
		}
		responses[name] = respRef
	}
	return responses, nil
}
