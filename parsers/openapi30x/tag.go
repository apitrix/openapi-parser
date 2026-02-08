package openapi30x

import (
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

type tagParser struct{}

// defaultTagParser is the singleton instance used by parsing functions.
var defaultTagParser = &tagParser{}

// parseSharedTag parses a Tag object from a yaml.Node.
func parseSharedTag(node *yaml.Node, ctx *ParseContext) (*openapi30models.Tag, error) {
	return defaultTagParser.parse(node, ctx)
}

// parseSharedTags parses an array of Tag objects.
func parseSharedTags(node *yaml.Node, ctx *ParseContext) ([]*openapi30models.Tag, error) {
	if node == nil || !nodeIsSequence(node) {
		return nil, nil
	}

	tags := make([]*openapi30models.Tag, 0, len(node.Content))
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
func (p *tagParser) parse(node *yaml.Node, ctx *ParseContext) (*openapi30models.Tag, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "tag must be an object")
	}

	tag := &openapi30models.Tag{}
	var err error

	// Simple properties - inline
	tag.Name = p.ParseName(node)
	tag.Description = p.ParseDescription(node)

	// Complex properties - delegated to dedicated files
	tag.ExternalDocs, err = p.ParseExternalDocs(node, ctx)
	if err != nil {
		tag.Trix.Errors = append(tag.Trix.Errors, toParseError(err))
	}

	tag.VendorExtensions = parseNodeExtensions(node)
	tag.Trix.Source = ctx.nodeSource(node)

	// Detect unknown fields
	ctx.detectUnknown(node, tagKnownFieldsSet)

	return tag, nil
}

func (p *tagParser) ParseName(node *yaml.Node) string {
	return nodeGetString(node, "name")
}

func (p *tagParser) ParseDescription(node *yaml.Node) string {
	return nodeGetString(node, "description")
}
