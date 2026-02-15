package openapi30x

import (
	openapi30models "github.com/apitrix/openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

type oauthFlowsParser struct{}

// defaultOAuthFlowsParser is the singleton instance used by parsing functions.
var defaultOAuthFlowsParser = &oauthFlowsParser{}

// parseSharedOAuthFlows parses an OAuthFlows object from a yaml.Node.
func parseSharedOAuthFlows(node *yaml.Node, ctx *ParseContext) (*openapi30models.OAuthFlows, error) {
	return defaultOAuthFlowsParser.parse(node, ctx)
}

// Parse parses an OAuthFlows object.
func (p *oauthFlowsParser) parse(node *yaml.Node, ctx *ParseContext) (*openapi30models.OAuthFlows, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "oauthFlows must be an object")
	}

	var errors []openapi30models.ParseError

	// All properties are complex (nested OAuthFlow objects)
	implicit, err := p.ParseImplicit(node, ctx)
	if err != nil {
		errors = append(errors, toParseError(err))
	}

	password, err := p.ParsePassword(node, ctx)
	if err != nil {
		errors = append(errors, toParseError(err))
	}

	clientCredentials, err := p.ParseClientCredentials(node, ctx)
	if err != nil {
		errors = append(errors, toParseError(err))
	}

	authorizationCode, err := p.ParseAuthorizationCode(node, ctx)
	if err != nil {
		errors = append(errors, toParseError(err))
	}

	// Create via constructor
	flows := openapi30models.NewOAuthFlows(implicit, password, clientCredentials, authorizationCode)

	flows.VendorExtensions = parseNodeExtensions(node)
	flows.Trix.Source = ctx.nodeSource(node)
	flows.Trix.Errors = append(flows.Trix.Errors, errors...)

	// Detect unknown fields
	flows.Trix.Errors = append(flows.Trix.Errors, unknownFieldParseErrors(ctx.detectUnknown(node, oauthFlowsKnownFieldsSet))...)

	return flows, nil
}

func (p *oauthFlowsParser) ParseImplicit(parent *yaml.Node, c *ParseContext) (*openapi30models.OAuthFlow, error) {
	node := nodeGetValue(parent, "implicit")
	if node == nil {
		return nil, nil
	}
	return parseSharedOAuthFlow(node, c.Push("implicit"))
}

func (p *oauthFlowsParser) ParsePassword(parent *yaml.Node, c *ParseContext) (*openapi30models.OAuthFlow, error) {
	node := nodeGetValue(parent, "password")
	if node == nil {
		return nil, nil
	}
	return parseSharedOAuthFlow(node, c.Push("password"))
}

func (p *oauthFlowsParser) ParseClientCredentials(parent *yaml.Node, c *ParseContext) (*openapi30models.OAuthFlow, error) {
	node := nodeGetValue(parent, "clientCredentials")
	if node == nil {
		return nil, nil
	}
	return parseSharedOAuthFlow(node, c.Push("clientCredentials"))
}

func (p *oauthFlowsParser) ParseAuthorizationCode(parent *yaml.Node, c *ParseContext) (*openapi30models.OAuthFlow, error) {
	node := nodeGetValue(parent, "authorizationCode")
	if node == nil {
		return nil, nil
	}
	return parseSharedOAuthFlow(node, c.Push("authorizationCode"))
}
