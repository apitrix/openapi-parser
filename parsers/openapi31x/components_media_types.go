package openapi31x

import (
	openapi31models "github.com/apitrix/openapi-parser/models/openapi31"
	"github.com/apitrix/openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// parseComponentsMediaTypes parses the Components.MediaTypes field (OpenAPI 3.2).
func parseComponentsMediaTypes(parent *yaml.Node, ctx *ParseContext) (map[string]*shared.RefWithMeta[openapi31models.MediaType], error) {
	if parent == nil {
		return nil, nil
	}
	node := nodeGetValue(parent, "mediaTypes")
	if node == nil || !nodeIsMapping(node) {
		return nil, nil
	}

	mediaTypes := make(map[string]*shared.RefWithMeta[openapi31models.MediaType])
	mctx := ctx.push("mediaTypes")
	for name, mtNode := range nodeMapPairs(node) {
		mtRef, err := parseMediaTypeRef(mtNode, mctx.push(name))
		if err != nil {
			return nil, err
		}
		mediaTypes[name] = mtRef
	}
	return mediaTypes, nil
}
