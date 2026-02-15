package openapi20

import (
	openapi20models "github.com/apitrix/openapi-parser/models/openapi20"

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

	scheme := openapi20models.NewSecurityScheme(
		nodeGetString(node, "type"),
		nodeGetString(node, "description"),
		nodeGetString(node, "name"),
		nodeGetString(node, "in"),
		nodeGetString(node, "flow"),
		nodeGetString(node, "authorizationUrl"),
		nodeGetString(node, "tokenUrl"),
		nodeGetStringMap(node, "scopes"),
	)

	scheme.VendorExtensions = parseNodeExtensions(node)
	scheme.Trix.Source = ctx.nodeSource(node)

	// Detect unknown fields
	scheme.Trix.Errors = append(scheme.Trix.Errors, unknownFieldParseErrors(ctx.detectUnknown(node, securitySchemeKnownFieldsSet))...)

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
