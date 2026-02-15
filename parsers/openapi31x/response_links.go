package openapi31x

import (
	openapi31models "github.com/apitrix/openapi-parser/models/openapi31"
	"github.com/apitrix/openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// ParseLinks parses the Response.Links field.
func (p *responseParser) ParseLinks(parent *yaml.Node, c *ParseContext) (map[string]*shared.RefWithMeta[openapi31models.Link], error) {
	node := nodeGetValue(parent, "links")
	if node == nil || !nodeIsMapping(node) {
		return nil, nil
	}

	links := make(map[string]*shared.RefWithMeta[openapi31models.Link])
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
