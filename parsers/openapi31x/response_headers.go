package openapi31x

import (
	openapi31models "github.com/apitrix/openapi-parser/models/openapi31"
	"github.com/apitrix/openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// ParseHeaders parses the Response.Headers field.
func (p *responseParser) ParseHeaders(parent *yaml.Node, c *ParseContext) (map[string]*shared.RefWithMeta[openapi31models.Header], error) {
	node := nodeGetValue(parent, "headers")
	if node == nil || !nodeIsMapping(node) {
		return nil, nil
	}

	headers := make(map[string]*shared.RefWithMeta[openapi31models.Header])
	hctx := c.Push("headers")
	for name, headerNode := range nodeMapPairs(node) {
		headerRef, err := parseHeaderRef(headerNode, hctx.push(name))
		if err != nil {
			return nil, err
		}
		headers[name] = headerRef
	}
	return headers, nil
}
