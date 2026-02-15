package openapi30x

import (
	openapi30models "github.com/apitrix/openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// parseOperationExternalDocs parses the Operation.ExternalDocs field.
func parseOperationExternalDocs(parent *yaml.Node, ctx *ParseContext) (*openapi30models.ExternalDocumentation, error) {
	node := nodeGetValue(parent, "externalDocs")
	if node == nil {
		return nil, nil
	}
	return parseSharedExternalDocs(node, ctx.push("externalDocs"))
}
