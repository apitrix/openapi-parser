package openapi30x

import (
	openapi30models "github.com/apitrix/openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

type securitySchemeParser struct{}

// defaultSecuritySchemeParser is the singleton instance used by parsing functions.
var defaultSecuritySchemeParser = &securitySchemeParser{}

// parseSharedSecurityScheme parses a SecurityScheme object from a yaml.Node.
func parseSharedSecurityScheme(node *yaml.Node, ctx *ParseContext) (*openapi30models.SecurityScheme, error) {
	return defaultSecuritySchemeParser.parse(node, ctx)
}

// Parse parses a SecurityScheme object.
func (p *securitySchemeParser) parse(node *yaml.Node, ctx *ParseContext) (*openapi30models.SecurityScheme, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "securityScheme must be an object")
	}

	// Collect values
	secType := p.ParseType(node)
	description := p.ParseDescription(node)
	name := p.ParseName(node)
	in := p.ParseIn(node)
	scheme := p.ParseScheme(node)
	bearerFormat := p.ParseBearerFormat(node)
	openIDConnectURL := p.ParseOpenIDConnectURL(node)

	flows, err := p.ParseFlows(node, ctx)

	// Create via constructor
	ss := openapi30models.NewSecurityScheme(secType, description, name, in, scheme, bearerFormat, flows, openIDConnectURL)

	if err != nil {
		ss.Trix.Errors = append(ss.Trix.Errors, toParseError(err))
	}

	ss.VendorExtensions = parseNodeExtensions(node)
	ss.Trix.Source = ctx.nodeSource(node)

	// Detect unknown fields
	ss.Trix.Errors = append(ss.Trix.Errors, unknownFieldParseErrors(ctx.detectUnknown(node, securitySchemeKnownFieldsSet))...)

	return ss, nil
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
