package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

type oauthFlowParser struct{}

// defaultOAuthFlowParser is the singleton instance used by parsing functions.
var defaultOAuthFlowParser = &oauthFlowParser{}

// parseSharedOAuthFlow parses an OAuthFlow object from a yaml.Node.
func parseSharedOAuthFlow(node *yaml.Node, ctx *ParseContext) (*openapi31models.OAuthFlow, error) {
	return defaultOAuthFlowParser.parse(node, ctx)
}

// Parse parses an OAuthFlow object.
func (p *oauthFlowParser) parse(node *yaml.Node, ctx *ParseContext) (*openapi31models.OAuthFlow, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "oauthFlow must be an object")
	}

	flow := &openapi31models.OAuthFlow{}

	// All properties are simple - inline
	flow.AuthorizationURL = p.ParseAuthorizationURL(node)
	flow.TokenURL = p.ParseTokenURL(node)
	flow.RefreshURL = p.ParseRefreshURL(node)
	flow.Scopes = p.ParseScopes(node)

	flow.Extensions = parseNodeExtensions(node)
	flow.NodeSource = ctx.nodeSource(node)

	// Detect unknown fields
	ctx.detectUnknown(node, oauthFlowKnownFieldsSet)

	return flow, nil
}

func (p *oauthFlowParser) ParseAuthorizationURL(node *yaml.Node) string {
	return nodeGetString(node, "authorizationUrl")
}

func (p *oauthFlowParser) ParseTokenURL(node *yaml.Node) string {
	return nodeGetString(node, "tokenUrl")
}

func (p *oauthFlowParser) ParseRefreshURL(node *yaml.Node) string {
	return nodeGetString(node, "refreshUrl")
}

func (p *oauthFlowParser) ParseScopes(node *yaml.Node) map[string]string {
	return nodeGetStringMap(node, "scopes")
}
