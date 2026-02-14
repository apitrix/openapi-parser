package openapi30x

import (
	"openapi-parser/models/shared"
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// parseComponentsLinks parses the Components.Links field.
func parseComponentsLinks(parent *yaml.Node, ctx *ParseContext) (map[string]*shared.Ref[openapi30models.Link], error) {
	node := nodeGetValue(parent, "links")
	if node == nil || !nodeIsMapping(node) {
		return nil, nil
	}

	links := make(map[string]*shared.Ref[openapi30models.Link])
	lctx := ctx.push("links")
	for name, linkNode := range nodeMapPairs(node) {
		linkRef, err := parseLinkRef(linkNode, lctx.push(name))
		if err != nil {
			return nil, err
		}
		links[name] = linkRef
	}
	return links, nil
}
