package openapi30x

import (
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// parseComponentsHeaders parses the Components.Headers field.
func parseComponentsHeaders(parent *yaml.Node, ctx *ParseContext) (map[string]*openapi30models.RefHeader, error) {
	node := nodeGetValue(parent, "headers")
	if node == nil || !nodeIsMapping(node) {
		return nil, nil
	}

	headers := make(map[string]*openapi30models.RefHeader)
	hctx := ctx.push("headers")
	for name, headerNode := range nodeMapPairs(node) {
		headerRef, err := parseHeaderRef(headerNode, hctx.push(name))
		if err != nil {
			return nil, err
		}
		headers[name] = headerRef
	}
	return headers, nil
}
