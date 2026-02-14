package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// parseOperationCallbacks parses the Operation.Callbacks field.
func parseOperationCallbacks(parent *yaml.Node, ctx *ParseContext) (map[string]*shared.RefWithMeta[openapi31models.Callback], error) {
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
