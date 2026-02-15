package openapi31x

import (
	openapi31models "github.com/apitrix/openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

type tagParser struct{}

// defaultTagParser is the singleton instance used by parsing functions.
var defaultTagParser = &tagParser{}

// parseSharedTag parses a Tag object from a yaml.Node.
func parseSharedTag(node *yaml.Node, ctx *ParseContext) (*openapi31models.Tag, error) {
	return defaultTagParser.parse(node, ctx)
}

// parseSharedTags parses an array of Tag objects.
func parseSharedTags(node *yaml.Node, ctx *ParseContext) ([]*openapi31models.Tag, error) {
	if node == nil || !nodeIsSequence(node) {
		return nil, nil
	}

	tags := make([]*openapi31models.Tag, 0, len(node.Content))
	for i, tagNode := range node.Content {
		tag, err := parseSharedTag(tagNode, ctx.push(itoa(i)))
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

// Parse parses a Tag object.
func (p *tagParser) parse(node *yaml.Node, ctx *ParseContext) (*openapi31models.Tag, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "tag must be an object")
	}

	var errs []openapi31models.ParseError

	// Complex properties - delegated to dedicated files
	externalDocs, err := p.ParseExternalDocs(node, ctx)
	if err != nil {
		errs = append(errs, toParseError(err))
	}

	// Create via constructor
	tag := openapi31models.NewTag(
		p.ParseName(node),
		p.ParseDescription(node),
		externalDocs,
	)

	tag.VendorExtensions = parseNodeExtensions(node)
	tag.Trix.Source = ctx.nodeSource(node)
	tag.Trix.Errors = append(tag.Trix.Errors, errs...)

	// Detect unknown fields
	tag.Trix.Errors = append(tag.Trix.Errors, unknownFieldParseErrors(ctx.detectUnknown(node, tagKnownFieldsSet))...)

	return tag, nil
}

func (p *tagParser) ParseName(node *yaml.Node) string {
	return nodeGetString(node, "name")
}

func (p *tagParser) ParseDescription(node *yaml.Node) string {
	return nodeGetString(node, "description")
}
