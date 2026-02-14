package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// parseComponentsRequestBodies parses the Components.RequestBodies field.
func parseComponentsRequestBodies(parent *yaml.Node, ctx *ParseContext) (map[string]*shared.RefWithMeta[openapi31models.RequestBody], error) {
	node := nodeGetValue(parent, "requestBodies")
	if node == nil || !nodeIsMapping(node) {
		return nil, nil
	}

	requestBodies := make(map[string]*shared.RefWithMeta[openapi31models.RequestBody])
	rctx := ctx.push("requestBodies")
	for name, rbNode := range nodeMapPairs(node) {
		rbRef, err := parseRequestBodyRef(rbNode, rctx.push(name))
		if err != nil {
			return nil, err
		}
		requestBodies[name] = rbRef
	}
	return requestBodies, nil
}
