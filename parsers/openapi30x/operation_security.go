package openapi30x

import (
	openapi30models "github.com/apitrix/openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// parseOperationSecurity parses the Operation.Security field.
func parseOperationSecurity(parent *yaml.Node, ctx *ParseContext) ([]openapi30models.SecurityRequirement, error) {
	node := nodeGetValue(parent, "security")
	if node == nil {
		return nil, nil
	}
	return parseSharedSecurityRequirements(node, ctx.push("security"))
}
