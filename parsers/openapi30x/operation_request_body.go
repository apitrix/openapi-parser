package openapi30x

import (
	"openapi-parser/models/shared"
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// parseOperationRequestBody parses the Operation.RequestBody field.
func parseOperationRequestBody(parent *yaml.Node, ctx *ParseContext) (*shared.Ref[openapi30models.RequestBody], error) {
	node := nodeGetValue(parent, "requestBody")
	if node == nil {
		return nil, nil
	}
	return parseRequestBodyRef(node, ctx.push("requestBody"))
}
