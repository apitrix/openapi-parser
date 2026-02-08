package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// ParseEncoding parses the MediaType.Encoding field.
func (p *mediaTypeParser) ParseEncoding(parent *yaml.Node, c *ParseContext) (map[string]*openapi31models.Encoding, error) {
	node := nodeGetValue(parent, "encoding")
	if node == nil || !nodeIsMapping(node) {
		return nil, nil
	}

	encoding := make(map[string]*openapi31models.Encoding)
	ectx := c.Push("encoding")
	for propName, encNode := range nodeMapPairs(node) {
		enc, err := parseSharedEncoding(encNode, ectx.push(propName))
		if err != nil {
			return nil, err
		}
		encoding[propName] = enc
	}
	return encoding, nil
}
