package openapi30x

import (
	openapi30models "github.com/apitrix/openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// parseComponentsLinks parses the Components.Links field.
func parseComponentsLinks(parent *yaml.Node, ctx *ParseContext) (map[string]*openapi30models.RefLink, error) {
	node := nodeGetValue(parent, "links")
	if node == nil || !nodeIsMapping(node) {
		return nil, nil
	}

	links := make(map[string]*openapi30models.RefLink)
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
