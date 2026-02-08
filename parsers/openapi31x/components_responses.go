package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// parseComponentsResponses parses the Components.Responses field.
func parseComponentsResponses(parent *yaml.Node, ctx *ParseContext) (map[string]*openapi31models.ResponseRef, error) {
	node := nodeGetValue(parent, "responses")
	if node == nil || !nodeIsMapping(node) {
		return nil, nil
	}

	responses := make(map[string]*openapi31models.ResponseRef)
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
