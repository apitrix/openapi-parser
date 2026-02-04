package openapi30

import (
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// parseComponentsExamples parses the Components.Examples field.
func parseComponentsExamples(parent *yaml.Node, ctx *ParseContext) (map[string]*openapi30models.ExampleRef, error) {
	node := nodeGetValue(parent, "examples")
	if node == nil || !nodeIsMapping(node) {
		return nil, nil
	}

	examples := make(map[string]*openapi30models.ExampleRef)
	ectx := ctx.push("examples")
	for _, name := range nodeKeys(node) {
		exampleNode := nodeGetValue(node, name)
		exampleRef, err := parseExampleRef(exampleNode, ectx.push(name))
		if err != nil {
			return nil, err
		}
		examples[name] = exampleRef
	}
	return examples, nil
}
