package openapi31x

import (
	"openapi-parser/models/shared"
	openapi31models "openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// parseComponentsCallbacks parses the Components.Callbacks field.
func parseComponentsCallbacks(parent *yaml.Node, ctx *ParseContext) (map[string]*shared.RefWithMeta[openapi31models.Callback], error) {
	node := nodeGetValue(parent, "callbacks")
	if node == nil || !nodeIsMapping(node) {
		return nil, nil
	}

	callbacks := make(map[string]*shared.RefWithMeta[openapi31models.Callback])
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
