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

	var err error

	// Enum - array of any values
	var enum []interface{}
	if enumNode := nodeGetValue(node, "enum"); enumNode != nil && nodeIsSequence(enumNode) {
		enum = make([]interface{}, len(enumNode.Content))
		for i, item := range enumNode.Content {
			enum[i] = nodeToInterface(item)
		}
	}

	// Items - for array type headers (parsed first for constructor)
	var items *openapi20models.Items
	var itemsErr error
	if itemsNode := nodeGetValue(node, "items"); itemsNode != nil {
		items, err = parseItems(itemsNode, ctx.push("items"))
		if err != nil {
			itemsErr = err
		}
	}

	header := openapi20models.NewHeader(openapi20models.HeaderFields{
		Description:      nodeGetString(node, "description"),
		Type:             nodeGetString(node, "type"),
		Format:           nodeGetString(node, "format"),
		Items:            items,
		CollectionFormat: nodeGetString(node, "collectionFormat"),
		Default:          nodeGetAny(node, "default"),
		Maximum:          nodeGetFloat64Ptr(node, "maximum"),
		ExclusiveMaximum: nodeGetBool(node, "exclusiveMaximum"),
		Minimum:          nodeGetFloat64Ptr(node, "minimum"),
		ExclusiveMinimum: nodeGetBool(node, "exclusiveMinimum"),
		MaxLength:        nodeGetUint64Ptr(node, "maxLength"),
		MinLength:        nodeGetUint64Ptr(node, "minLength"),
		Pattern:          nodeGetString(node, "pattern"),
		MaxItems:         nodeGetUint64Ptr(node, "maxItems"),
		MinItems:         nodeGetUint64Ptr(node, "minItems"),
		UniqueItems:      nodeGetBool(node, "uniqueItems"),
		Enum:             enum,
		MultipleOf:       nodeGetFloat64Ptr(node, "multipleOf"),
	})

	if itemsErr != nil {
		header.Trix.Errors = append(header.Trix.Errors, toParseError(itemsErr))
	}

	header.VendorExtensions = parseNodeExtensions(node)
	header.Trix.Source = ctx.nodeSource(node)

	// Detect unknown fields
	header.Trix.Errors = append(header.Trix.Errors, unknownFieldParseErrors(ctx.detectUnknown(node, headerKnownFieldsSet))...)

	return header, nil
}
