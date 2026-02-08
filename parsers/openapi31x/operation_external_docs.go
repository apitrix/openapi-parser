package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// parseOperationExternalDocs parses the Operation.ExternalDocs field.
func parseOperationExternalDocs(parent *yaml.Node, ctx *ParseContext) (*openapi31models.ExternalDocumentation, error) {
	node := nodeGetValue(parent, "externalDocs")
	if node == nil {
		return nil, nil
	}
	return parseSharedExternalDocs(node, ctx.push("externalDocs"))
}
