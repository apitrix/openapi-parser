package openapi30x

import (
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// parseComponentsSecuritySchemes parses the Components.SecuritySchemes field.
func parseComponentsSecuritySchemes(parent *yaml.Node, ctx *ParseContext) (map[string]*openapi30models.SecuritySchemeRef, error) {
	node := nodeGetValue(parent, "securitySchemes")
	if node == nil || !nodeIsMapping(node) {
		return nil, nil
	}

	schemes := make(map[string]*openapi30models.SecuritySchemeRef)
	sctx := ctx.push("securitySchemes")
	for name, schemeNode := range nodeMapPairs(node) {
		schemeRef, err := parseSecuritySchemeRef(schemeNode, sctx.push(name))
		if err != nil {
			return nil, err
		}
		schemes[name] = schemeRef
	}
	return schemes, nil
}
