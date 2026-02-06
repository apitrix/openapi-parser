package openapi30

import (
	"gopkg.in/yaml.v3"
)

// ParseParameters parses the Link.Parameters field.
func (p *linkParser) ParseParameters(parent *yaml.Node, c *ParseContext) (map[string]interface{}, error) {
	node := nodeGetValue(parent, "parameters")
	if node == nil || !nodeIsMapping(node) {
		return nil, nil
	}

	params := make(map[string]interface{})
	for name, paramNode := range nodeMapPairs(node) {
		params[name] = nodeToInterface(paramNode)
	}
	return params, nil
}
