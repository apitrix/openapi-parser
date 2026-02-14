package openapi30x

import (
	"openapi-parser/models/shared"
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// ParseExamples parses the MediaType.Examples field.
func (p *mediaTypeParser) ParseExamples(parent *yaml.Node, c *ParseContext) (map[string]*shared.Ref[openapi30models.Example], error) {
	node := nodeGetValue(parent, "examples")
	if node == nil || !nodeIsMapping(node) {
		return nil, nil
	}

	examples := make(map[string]*shared.Ref[openapi30models.Example])
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
