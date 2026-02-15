package openapi31x

import (
	openapi31models "github.com/apitrix/openapi-parser/models/openapi31"
	"github.com/apitrix/openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// parseOperationRequestBody parses the Operation.RequestBody field.
func parseOperationRequestBody(parent *yaml.Node, ctx *ParseContext) (*shared.RefWithMeta[openapi31models.RequestBody], error) {
	node := nodeGetValue(parent, "requestBody")
	if node == nil {
		return nil, nil
	}
	return parseRequestBodyRef(node, ctx.push("requestBody"))
}
