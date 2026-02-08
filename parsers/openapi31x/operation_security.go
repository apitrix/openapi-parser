package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// parseOperationSecurity parses the Operation.Security field.
func parseOperationSecurity(parent *yaml.Node, ctx *ParseContext) ([]openapi31models.SecurityRequirement, error) {
	node := nodeGetValue(parent, "security")
	if node == nil {
		return nil, nil
	}
	return parseSharedSecurityRequirements(node, ctx.push("security"))
}
