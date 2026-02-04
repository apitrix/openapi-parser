package openapi30

import (
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

// parseSecuritySchemeRef parses a SecuritySchemeRef from a yaml.Node.
func parseSecuritySchemeRef(node *yaml.Node, ctx *ParseContext) (*openapi30models.SecuritySchemeRef, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "securityScheme must be an object")
	}

	ref := &openapi30models.SecuritySchemeRef{}
	ref.NodeSource = ctx.nodeSource(node)
	ref.Extensions = parseNodeExtensions(node)

	// Check for $ref
	if nodeHasRef(node) {
		ref.Ref = nodeGetRef(node)
		return ref, nil
	}

	// Parse inline security scheme
	scheme, err := parseSharedSecurityScheme(node, ctx)
	if err != nil {
		return nil, err
	}
	ref.Value = scheme

	return ref, nil
}
