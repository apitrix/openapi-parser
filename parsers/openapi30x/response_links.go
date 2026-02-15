package openapi30x

import (
	openapi30models "github.com/apitrix/openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// ParseLinks parses the Response.Links field.
func (p *responseParser) ParseLinks(parent *yaml.Node, c *ParseContext) (map[string]*openapi30models.RefLink, error) {
	node := nodeGetValue(parent, "links")
	if node == nil || !nodeIsMapping(node) {
		return nil, nil
	}

	links := make(map[string]*openapi30models.RefLink)
	lctx := c.Push("links")
	for name, linkNode := range nodeMapPairs(node) {
		linkRef, err := parseLinkRef(linkNode, lctx.push(name))
		if err != nil {
			return nil, err
		}
		links[name] = linkRef
	}
	return links, nil
}
