package openapi30

import (
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// ParseLinks parses the Response.Links field.
func (p *responseParser) ParseLinks(parent *yaml.Node, c *ParseContext) (map[string]*openapi30models.LinkRef, error) {
	node := nodeGetValue(parent, "links")
	if node == nil || !nodeIsMapping(node) {
		return nil, nil
	}

	links := make(map[string]*openapi30models.LinkRef)
	lctx := c.Push("links")
	for _, name := range nodeKeys(node) {
		linkNode := nodeGetValue(node, name)
		linkRef, err := parseLinkRef(linkNode, lctx.push(name))
		if err != nil {
			return nil, err
		}
		links[name] = linkRef
	}
	return links, nil
}
