package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// ParseLinks parses the Response.Links field.
func (p *responseParser) ParseLinks(parent *yaml.Node, c *ParseContext) (map[string]*openapi31models.LinkRef, error) {
	node := nodeGetValue(parent, "links")
	if node == nil || !nodeIsMapping(node) {
		return nil, nil
	}

	links := make(map[string]*openapi31models.LinkRef)
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
