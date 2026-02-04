package openapi30

import (
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// ParseEncoding parses the MediaType.Encoding field.
func (p *mediaTypeParser) ParseEncoding(parent *yaml.Node, c *ParseContext) (map[string]*openapi30models.Encoding, error) {
	node := nodeGetValue(parent, "encoding")
	if node == nil || !nodeIsMapping(node) {
		return nil, nil
	}

	encoding := make(map[string]*openapi30models.Encoding)
	ectx := c.Push("encoding")
	for _, propName := range nodeKeys(node) {
		encNode := nodeGetValue(node, propName)
		enc, err := parseSharedEncoding(encNode, ectx.push(propName))
		if err != nil {
			return nil, err
		}
		encoding[propName] = enc
	}
	return encoding, nil
}
