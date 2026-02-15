package openapi30x

import (
	openapi30models "github.com/apitrix/openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// ParseVariables parses the Server.Variables field.
func (p *serverParser) ParseVariables(parent *yaml.Node, c *ParseContext) (map[string]*openapi30models.ServerVariable, error) {
	node := nodeGetValue(parent, "variables")
	if node == nil || !nodeIsMapping(node) {
		return nil, nil
	}

	variables := make(map[string]*openapi30models.ServerVariable)
	vctx := c.Push("variables")
	for name, varNode := range nodeMapPairs(node) {
		sv, err := parseSharedServerVariable(varNode, vctx.push(name))
		if err != nil {
			return nil, err
		}
		variables[name] = sv
	}
	return variables, nil
}
