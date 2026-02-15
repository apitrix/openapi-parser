package openapi31x

import (
	openapi31models "github.com/apitrix/openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// ParseContent parses the Response.Content field.
func (p *responseParser) ParseContent(parent *yaml.Node, c *ParseContext) (map[string]*openapi31models.MediaType, error) {
	node := nodeGetValue(parent, "content")
	if node == nil || !nodeIsMapping(node) {
		return nil, nil
	}

	content := make(map[string]*openapi31models.MediaType)
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
