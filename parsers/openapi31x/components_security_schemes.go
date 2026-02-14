package openapi31x

import (
	"openapi-parser/models/shared"
	openapi31models "openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// parseComponentsSecuritySchemes parses the Components.SecuritySchemes field.
func parseComponentsSecuritySchemes(parent *yaml.Node, ctx *ParseContext) (map[string]*shared.RefWithMeta[openapi31models.SecurityScheme], error) {
	node := nodeGetValue(parent, "securitySchemes")
	if node == nil || !nodeIsMapping(node) {
		return nil, nil
	}

	schemes := make(map[string]*shared.RefWithMeta[openapi31models.SecurityScheme])
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
