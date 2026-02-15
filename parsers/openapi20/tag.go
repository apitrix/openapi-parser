package openapi20

import (
	openapi20models "github.com/apitrix/openapi-parser/models/openapi20"

	"gopkg.in/yaml.v3"
)

// parseTags parses an array of Tag objects from a yaml.Node.
func parseTags(node *yaml.Node, ctx *ParseContext) ([]*openapi20models.Tag, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsSequence(node) {
		return nil, ctx.errorAt(node, "tags must be an array")
	}

	tags := make([]*openapi20models.Tag, 0, len(node.Content))
	for i, tagNode := range node.Content {
		tag, err := parseTag(tagNode, ctx.push(itoa(i)))
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}

// parseTag parses a Tag object from a yaml.Node.
func parseTag(node *yaml.Node, ctx *ParseContext) (*openapi20models.Tag, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "tag must be an object")
	}

	var err error

	// Complex property - ExternalDocs (parsed first for constructor)
	var externalDocs *openapi20models.ExternalDocs
	var edErr error
	if edNode := nodeGetValue(node, "externalDocs"); edNode != nil {
		externalDocs, err = parseExternalDocs(edNode, ctx.push("externalDocs"))
		if err != nil {
			edErr = err
		}
	}

	tag := openapi20models.NewTag(
		nodeGetString(node, "name"),
		nodeGetString(node, "description"),
		externalDocs,
	)

	if edErr != nil {
		tag.Trix.Errors = append(tag.Trix.Errors, toParseError(edErr))
	}

	tag.VendorExtensions = parseNodeExtensions(node)
	tag.Trix.Source = ctx.nodeSource(node)

	// Detect unknown fields
	tag.Trix.Errors = append(tag.Trix.Errors, unknownFieldParseErrors(ctx.detectUnknown(node, tagKnownFieldsSet))...)

	return tag, nil
}
