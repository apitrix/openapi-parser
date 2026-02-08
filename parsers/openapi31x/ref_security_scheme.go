package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

// parseSecuritySchemeRef parses a SecuritySchemeRef from a yaml.Node.
func parseSecuritySchemeRef(node *yaml.Node, ctx *ParseContext) (*openapi31models.SecuritySchemeRef, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "securityScheme must be an object")
	}

	ref := &openapi31models.SecuritySchemeRef{}
	ref.NodeSource = ctx.nodeSource(node)
	ref.VendorExtensions = parseNodeExtensions(node)

	if nodeHasRef(node) {
		ref.Ref = nodeGetRef(node)
		ref.Summary = nodeGetString(node, "summary")
		ref.Description = nodeGetString(node, "description")
		return ref, nil
	}

	scheme, err := parseSharedSecurityScheme(node, ctx)
	if err != nil {
		return nil, err
	}
	ref.Value = scheme

	return ref, nil
}
