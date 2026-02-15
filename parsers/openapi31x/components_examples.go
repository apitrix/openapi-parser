package openapi31x

import (
	openapi31models "github.com/apitrix/openapi-parser/models/openapi31"
	"github.com/apitrix/openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// parseComponentsExamples parses the Components.Examples field.
func parseComponentsExamples(parent *yaml.Node, ctx *ParseContext) (map[string]*shared.RefWithMeta[openapi31models.Example], error) {
	node := nodeGetValue(parent, "examples")
	if node == nil || !nodeIsMapping(node) {
		return nil, nil
	}

	examples := make(map[string]*shared.RefWithMeta[openapi31models.Example])
	ectx := ctx.push("examples")
	for name, exampleNode := range nodeMapPairs(node) {
		exampleRef, err := parseExampleRef(exampleNode, ectx.push(name))
		if err != nil {
			return nil, err
		}
		examples[name] = exampleRef
	}
	return examples, nil
}
