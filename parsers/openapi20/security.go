package openapi20

import (
	openapi20models "openapi-parser/models/openapi20"

	"gopkg.in/yaml.v3"
)

// parseSecurityScheme parses a SecurityScheme object from a yaml.Node.
func parseSecurityScheme(node *yaml.Node, ctx *ParseContext) (*openapi20models.SecurityScheme, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "securityScheme must be an object")
	}

	scheme := &openapi20models.SecurityScheme{}

	// Simple properties - inline
	scheme.Type = nodeGetString(node, "type")
	scheme.Description = nodeGetString(node, "description")
	scheme.Name = nodeGetString(node, "name")
	scheme.In = nodeGetString(node, "in")
	scheme.Flow = nodeGetString(node, "flow")
	scheme.AuthorizationURL = nodeGetString(node, "authorizationUrl")
	scheme.TokenURL = nodeGetString(node, "tokenUrl")
	scheme.Scopes = nodeGetStringMap(node, "scopes")

	scheme.VendorExtensions = parseNodeExtensions(node)
	scheme.NodeSource = ctx.nodeSource(node)

	// Detect unknown fields
	ctx.detectUnknown(node, securitySchemeKnownFieldsSet)

	return scheme, nil
}

// parseSecurityRequirements parses an array of SecurityRequirement objects.
func parseSecurityRequirements(node *yaml.Node, ctx *ParseContext) ([]openapi20models.SecurityRequirement, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsSequence(node) {
		return nil, ctx.errorAt(node, "security must be an array")
	}

	reqs := make([]openapi20models.SecurityRequirement, 0, len(node.Content))
	for i, itemNode := range node.Content {
		req, err := parseSecurityRequirement(itemNode, ctx.push(itoa(i)))
		if err != nil {
			return nil, err
		}
		reqs = append(reqs, req)
	}

	return reqs, nil
}

// parseSecurityRequirement parses a single SecurityRequirement from a yaml.Node.
func parseSecurityRequirement(node *yaml.Node, ctx *ParseContext) (openapi20models.SecurityRequirement, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "security requirement must be an object")
	}

	req := make(openapi20models.SecurityRequirement)

	for key, valNode := range nodeMapPairs(node) {
		if valNode == nil || !nodeIsSequence(valNode) {
			req[key] = []string{}
			continue
		}

		scopes := make([]string, 0, len(valNode.Content))
		for _, scopeNode := range valNode.Content {
			scopes = append(scopes, scopeNode.Value)
		}
		req[key] = scopes
	}

	return req, nil
}
