package openapi30

import (
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// parseOperationRequestBody parses the Operation.RequestBody field.
func parseOperationRequestBody(parent *yaml.Node, ctx *ParseContext) (*openapi30models.RequestBodyRef, error) {
	node := nodeGetValue(parent, "requestBody")
	if node == nil {
		return nil, nil
	}
	return parseRequestBodyRef(node, ctx.push("requestBody"))
}
