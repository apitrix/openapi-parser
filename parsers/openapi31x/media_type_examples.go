package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// ParseExamples parses the MediaType.Examples field.
func (p *mediaTypeParser) ParseExamples(parent *yaml.Node, c *ParseContext) (map[string]*shared.RefWithMeta[openapi31models.Example], error) {
	node := nodeGetValue(parent, "examples")
	if node == nil || !nodeIsMapping(node) {
		return nil, nil
	}

	examples := make(map[string]*shared.RefWithMeta[openapi31models.Example])
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
