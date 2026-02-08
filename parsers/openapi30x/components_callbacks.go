package openapi30x

import (
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// parseComponentsCallbacks parses the Components.Callbacks field.
func parseComponentsCallbacks(parent *yaml.Node, ctx *ParseContext) (map[string]*openapi30models.CallbackRef, error) {
	node := nodeGetValue(parent, "callbacks")
	if node == nil || !nodeIsMapping(node) {
		return nil, nil
	}

	callbacks := make(map[string]*openapi30models.CallbackRef)
	cctx := ctx.push("callbacks")
	for name, cbNode := range nodeMapPairs(node) {
		cbRef, err := parseCallbackRef(cbNode, cctx.push(name))
		if err != nil {
			return nil, err
		}
		callbacks[name] = cbRef
	}
	return callbacks, nil
}
