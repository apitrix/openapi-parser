package openapi30x

import (
	openapi30models "github.com/apitrix/openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// ParseExamples parses the Parameter.Examples field.
func (p *parameterParser) ParseExamples(parent *yaml.Node, c *ParseContext) (map[string]*openapi30models.RefExample, error) {
	node := nodeGetValue(parent, "examples")
	if node == nil || !nodeIsMapping(node) {
		return nil, nil
	}

	examples := make(map[string]*openapi30models.RefExample)
	ectx := c.Push("examples")
	for name, exampleNode := range nodeMapPairs(node) {
		exampleRef, err := parseExampleRef(exampleNode, ectx.push(name))
		if err != nil {
			return nil, err
		}
		examples[name] = exampleRef
	}
	return examples, nil
}
