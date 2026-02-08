package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// parseSharedSecurityRequirement parses a SecurityRequirement from a yaml.Node.
// OpenAPI 3.0.3 spec: https://spec.openapis.org/oas/v3.0.3#security-requirement-object
func parseSharedSecurityRequirement(node *yaml.Node, ctx *ParseContext) (openapi31models.SecurityRequirement, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "securityRequirement must be an object")
	}

	req := make(openapi31models.SecurityRequirement)

	// Each key is a security scheme name, value is array of scope names
	for name, scopeNode := range nodeMapPairs(node) {
		// Skip extensions
		if len(name) > 2 && name[0] == 'x' && name[1] == '-' {
			continue
		}

		scopes := make([]string, 0)
		if scopeNode != nil && nodeIsSequence(scopeNode) {
			for _, s := range scopeNode.Content {
				scopes = append(scopes, s.Value)
			}
		}
		req[name] = scopes
	}

	return req, nil
}

// parseSharedSecurityRequirements parses a slice of SecurityRequirement from a yaml.Node.
func parseSharedSecurityRequirements(node *yaml.Node, ctx *ParseContext) ([]openapi31models.SecurityRequirement, error) {
	if node == nil || !nodeIsSequence(node) {
		return nil, nil
	}

	requirements := make([]openapi31models.SecurityRequirement, 0, len(node.Content))
	for i, reqNode := range node.Content {
		req, err := parseSharedSecurityRequirement(reqNode, ctx.push(itoa(i)))
		if err != nil {
			return nil, err
		}
		if req != nil {
			requirements = append(requirements, req)
		}
	}

	return requirements, nil
}
