package openapi20

import (
	openapi20models "openapi-parser/models/openapi20"

	"gopkg.in/yaml.v3"
)

// parseHeaders parses a map of Header objects from a yaml.Node.
func parseHeaders(node *yaml.Node, ctx *ParseContext) (map[string]*openapi20models.Header, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "headers must be an object")
	}

	headers := make(map[string]*openapi20models.Header)

	for key, headerNode := range nodeMapPairs(node) {
		// Skip extensions
		if len(key) > 2 && key[0] == 'x' && key[1] == '-' {
			continue
		}

		header, err := parseHeader(headerNode, ctx.push(key))
		if err != nil {
			return nil, err
		}
		headers[key] = header
	}

	return headers, nil
}

// parseHeader parses a Header object from a yaml.Node.
func parseHeader(node *yaml.Node, ctx *ParseContext) (*openapi20models.Header, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "header must be an object")
	}

	header := &openapi20models.Header{}
	var err error

	// Simple properties - inline
	header.Description = nodeGetString(node, "description")
	header.Type = nodeGetString(node, "type")
	header.Format = nodeGetString(node, "format")
	header.CollectionFormat = nodeGetString(node, "collectionFormat")
	header.Default = nodeGetAny(node, "default")
	header.Maximum = nodeGetFloat64Ptr(node, "maximum")
	header.ExclusiveMaximum = nodeGetBool(node, "exclusiveMaximum")
	header.Minimum = nodeGetFloat64Ptr(node, "minimum")
	header.ExclusiveMinimum = nodeGetBool(node, "exclusiveMinimum")
	header.MaxLength = nodeGetUint64Ptr(node, "maxLength")
	header.MinLength = nodeGetUint64Ptr(node, "minLength")
	header.Pattern = nodeGetString(node, "pattern")
	header.MaxItems = nodeGetUint64Ptr(node, "maxItems")
	header.MinItems = nodeGetUint64Ptr(node, "minItems")
	header.UniqueItems = nodeGetBool(node, "uniqueItems")
	header.MultipleOf = nodeGetFloat64Ptr(node, "multipleOf")

	// Enum - array of any values
	if enumNode := nodeGetValue(node, "enum"); enumNode != nil && nodeIsSequence(enumNode) {
		header.Enum = make([]interface{}, len(enumNode.Content))
		for i, item := range enumNode.Content {
			header.Enum[i] = nodeToInterface(item)
		}
	}

	// Items - for array type headers
	if itemsNode := nodeGetValue(node, "items"); itemsNode != nil {
		header.Items, err = parseItems(itemsNode, ctx.push("items"))
		if err != nil {
			header.Trix.Errors = append(header.Trix.Errors, toParseError(err))
		}
	}

	header.VendorExtensions = parseNodeExtensions(node)
	header.Trix.Source = ctx.nodeSource(node)

	// Detect unknown fields
	ctx.detectUnknown(node, headerKnownFieldsSet)

	return header, nil
}
