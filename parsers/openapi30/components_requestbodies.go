package openapi30

import (
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// parseComponentsRequestBodies parses the Components.RequestBodies field.
func parseComponentsRequestBodies(parent *yaml.Node, ctx *ParseContext) (map[string]*openapi30models.RequestBodyRef, error) {
	node := nodeGetValue(parent, "requestBodies")
	if node == nil || !nodeIsMapping(node) {
		return nil, nil
	}

	requestBodies := make(map[string]*openapi30models.RequestBodyRef)
	rctx := ctx.push("requestBodies")
	for _, name := range nodeKeys(node) {
		rbNode := nodeGetValue(node, name)
		rbRef, err := parseRequestBodyRef(rbNode, rctx.push(name))
		if err != nil {
			return nil, err
		}
		requestBodies[name] = rbRef
	}
	return requestBodies, nil
}
