package openapi30x

import (
	openapi30models "github.com/apitrix/openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// ParseContent parses the Response.Content field.
func (p *responseParser) ParseContent(parent *yaml.Node, c *ParseContext) (map[string]*openapi30models.MediaType, error) {
	node := nodeGetValue(parent, "content")
	if node == nil || !nodeIsMapping(node) {
		return nil, nil
	}

	content := make(map[string]*openapi30models.MediaType)
	cctx := c.Push("content")
	for mediaTypeName, mtNode := range nodeMapPairs(node) {
		mt, err := parseSharedMediaType(mtNode, cctx.push(mediaTypeName))
		if err != nil {
			return nil, err
		}
		content[mediaTypeName] = mt
	}
	return content, nil
}
