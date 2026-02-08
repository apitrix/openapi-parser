package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// parseComponentsPathItems parses the Components.PathItems field.
// New in OpenAPI 3.1.
func parseComponentsPathItems(parent *yaml.Node, ctx *ParseContext) (map[string]*openapi31models.PathItemRef, error) {
	node := nodeGetValue(parent, "pathItems")
	if node == nil || !nodeIsMapping(node) {
		return nil, nil
	}

	pathItems := make(map[string]*openapi31models.PathItemRef)
	pctx := ctx.push("pathItems")
	for name, pathItemNode := range nodeMapPairs(node) {
		ref, err := parsePathItemRef(pathItemNode, pctx.push(name))
		if err != nil {
			return nil, err
		}
		pathItems[name] = ref
	}
	return pathItems, nil
}
