package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

type securitySchemeParser struct{}

// defaultSecuritySchemeParser is the singleton instance used by parsing functions.
var defaultSecuritySchemeParser = &securitySchemeParser{}

// parseSharedSecurityScheme parses a SecurityScheme object from a yaml.Node.
func parseSharedSecurityScheme(node *yaml.Node, ctx *ParseContext) (*openapi31models.SecurityScheme, error) {
	return defaultSecuritySchemeParser.parse(node, ctx)
}

// Parse parses a SecurityScheme object.
func (p *securitySchemeParser) parse(node *yaml.Node, ctx *ParseContext) (*openapi31models.SecurityScheme, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "securityScheme must be an object")
	}

	scheme := &openapi31models.SecurityScheme{}
	var err error

	// Simple properties - inline
	scheme.Type = p.ParseType(node)
	scheme.Description = p.ParseDescription(node)
	scheme.Name = p.ParseName(node)
	scheme.In = p.ParseIn(node)
	scheme.Scheme = p.ParseScheme(node)
	scheme.BearerFormat = p.ParseBearerFormat(node)
	scheme.OpenIDConnectURL = p.ParseOpenIDConnectURL(node)

	// Complex properties - delegated to dedicated files
	scheme.Flows, err = p.ParseFlows(node, ctx)
	if err != nil {
		return nil, err
	}

	scheme.Extensions = parseNodeExtensions(node)
	scheme.NodeSource = ctx.nodeSource(node)

	// Detect unknown fields
	ctx.detectUnknown(node, securitySchemeKnownFieldsSet)

	return scheme, nil
}

func (p *securitySchemeParser) ParseType(node *yaml.Node) string {
	return nodeGetString(node, "type")
}

func (p *securitySchemeParser) ParseDescription(node *yaml.Node) string {
	return nodeGetString(node, "description")
}

func (p *securitySchemeParser) ParseName(node *yaml.Node) string {
	return nodeGetString(node, "name")
}

func (p *securitySchemeParser) ParseIn(node *yaml.Node) string {
	return nodeGetString(node, "in")
}

func (p *securitySchemeParser) ParseScheme(node *yaml.Node) string {
	return nodeGetString(node, "scheme")
}

func (p *securitySchemeParser) ParseBearerFormat(node *yaml.Node) string {
	return nodeGetString(node, "bearerFormat")
}

func (p *securitySchemeParser) ParseOpenIDConnectURL(node *yaml.Node) string {
	return nodeGetString(node, "openIdConnectUrl")
}
