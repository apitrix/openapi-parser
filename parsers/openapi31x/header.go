package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

type headerParser struct{}

// defaultHeaderParser is the singleton instance used by parsing functions.
var defaultHeaderParser = &headerParser{}

// parseSharedHeader parses a Header object from a yaml.Node.
func parseSharedHeader(node *yaml.Node, ctx *ParseContext) (*openapi31models.Header, error) {
	return defaultHeaderParser.parse(node, ctx)
}

// Parse parses a Header object.
func (p *headerParser) parse(node *yaml.Node, ctx *ParseContext) (*openapi31models.Header, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "header must be an object")
	}

	var errs []openapi31models.ParseError

	// Complex properties - delegated to dedicated files
	schema, err := p.ParseSchema(node, ctx)
	if err != nil {
		errs = append(errs, toParseError(err))
	}

	examples, err := p.ParseExamples(node, ctx)
	if err != nil {
		errs = append(errs, toParseError(err))
	}

	content, err := p.ParseContent(node, ctx)
	if err != nil {
		errs = append(errs, toParseError(err))
	}

	// Create via constructor with HeaderFields
	header := openapi31models.NewHeader(openapi31models.HeaderFields{
		Description:     p.ParseDescription(node),
		Required:        p.ParseRequired(node),
		Deprecated:      p.ParseDeprecated(node),
		AllowEmptyValue: p.ParseAllowEmptyValue(node),
		Style:           p.ParseStyle(node),
		Explode:         p.ParseExplode(node),
		AllowReserved:   p.ParseAllowReserved(node),
		Schema:          schema,
		Example:         p.ParseExample(node),
		Examples:        examples,
		Content:         content,
	})

	header.VendorExtensions = parseNodeExtensions(node)
	header.Trix.Source = ctx.nodeSource(node)
	header.Trix.Errors = append(header.Trix.Errors, errs...)

	// Detect unknown fields
	header.Trix.Errors = append(header.Trix.Errors, unknownFieldParseErrors(ctx.detectUnknown(node, headerKnownFieldsSet))...)

	return header, nil
}

func (p *headerParser) ParseDescription(node *yaml.Node) string {
	return nodeGetString(node, "description")
}

func (p *headerParser) ParseRequired(node *yaml.Node) bool {
	return nodeGetBool(node, "required")
}

func (p *headerParser) ParseDeprecated(node *yaml.Node) bool {
	return nodeGetBool(node, "deprecated")
}

func (p *headerParser) ParseAllowEmptyValue(node *yaml.Node) bool {
	return nodeGetBool(node, "allowEmptyValue")
}

func (p *headerParser) ParseStyle(node *yaml.Node) string {
	return nodeGetString(node, "style")
}

func (p *headerParser) ParseExplode(node *yaml.Node) *bool {
	return nodeGetBoolPtr(node, "explode")
}

func (p *headerParser) ParseAllowReserved(node *yaml.Node) bool {
	return nodeGetBool(node, "allowReserved")
}

func (p *headerParser) ParseExample(node *yaml.Node) interface{} {
	return nodeGetAny(node, "example")
}
