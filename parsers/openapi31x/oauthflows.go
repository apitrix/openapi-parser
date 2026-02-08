package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

type oauthFlowsParser struct{}

// defaultOAuthFlowsParser is the singleton instance used by parsing functions.
var defaultOAuthFlowsParser = &oauthFlowsParser{}

// parseSharedOAuthFlows parses an OAuthFlows object from a yaml.Node.
func parseSharedOAuthFlows(node *yaml.Node, ctx *ParseContext) (*openapi31models.OAuthFlows, error) {
	return defaultOAuthFlowsParser.parse(node, ctx)
}

// Parse parses an OAuthFlows object.
func (p *oauthFlowsParser) parse(node *yaml.Node, ctx *ParseContext) (*openapi31models.OAuthFlows, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "oauthFlows must be an object")
	}

	flows := &openapi31models.OAuthFlows{}
	var err error

	// All properties are complex (nested OAuthFlow objects)
	flows.Implicit, err = p.ParseImplicit(node, ctx)
	if err != nil {
		return nil, err
	}

	flows.Password, err = p.ParsePassword(node, ctx)
	if err != nil {
		return nil, err
	}

	flows.ClientCredentials, err = p.ParseClientCredentials(node, ctx)
	if err != nil {
		return nil, err
	}

	flows.AuthorizationCode, err = p.ParseAuthorizationCode(node, ctx)
	if err != nil {
		return nil, err
	}

	flows.VendorExtensions = parseNodeExtensions(node)
	flows.Trix.Source = ctx.nodeSource(node)

	// Detect unknown fields
	ctx.detectUnknown(node, oauthFlowsKnownFieldsSet)

	return flows, nil
}

func (p *oauthFlowsParser) ParseImplicit(parent *yaml.Node, c *ParseContext) (*openapi31models.OAuthFlow, error) {
	node := nodeGetValue(parent, "implicit")
	if node == nil {
		return nil, nil
	}
	return parseSharedOAuthFlow(node, c.Push("implicit"))
}

func (p *oauthFlowsParser) ParsePassword(parent *yaml.Node, c *ParseContext) (*openapi31models.OAuthFlow, error) {
	node := nodeGetValue(parent, "password")
	if node == nil {
		return nil, nil
	}
	return parseSharedOAuthFlow(node, c.Push("password"))
}

func (p *oauthFlowsParser) ParseClientCredentials(parent *yaml.Node, c *ParseContext) (*openapi31models.OAuthFlow, error) {
	node := nodeGetValue(parent, "clientCredentials")
	if node == nil {
		return nil, nil
	}
	return parseSharedOAuthFlow(node, c.Push("clientCredentials"))
}

func (p *oauthFlowsParser) ParseAuthorizationCode(parent *yaml.Node, c *ParseContext) (*openapi31models.OAuthFlow, error) {
	node := nodeGetValue(parent, "authorizationCode")
	if node == nil {
		return nil, nil
	}
	return parseSharedOAuthFlow(node, c.Push("authorizationCode"))
}
