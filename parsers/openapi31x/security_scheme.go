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

	var errs []openapi31models.ParseError

	// Complex properties - delegated to dedicated files
	flows, err := p.ParseFlows(node, ctx)
	if err != nil {
		errs = append(errs, toParseError(err))
	}

	// Create via constructor
	scheme := openapi31models.NewSecurityScheme(
		p.ParseType(node),
		p.ParseDescription(node),
		p.ParseName(node),
		p.ParseIn(node),
		p.ParseScheme(node),
		p.ParseBearerFormat(node),
		p.ParseOpenIDConnectURL(node),
		flows,
	)

	scheme.VendorExtensions = parseNodeExtensions(node)
	scheme.Trix.Source = ctx.nodeSource(node)
	scheme.Trix.Errors = append(scheme.Trix.Errors, errs...)

	// Detect unknown fields
	scheme.Trix.Errors = append(scheme.Trix.Errors, unknownFieldParseErrors(ctx.detectUnknown(node, securitySchemeKnownFieldsSet))...)

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
