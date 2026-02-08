package openapi20

import (
	openapi20models "openapi-parser/models/openapi20"

	"gopkg.in/yaml.v3"
)

// parseItems parses an Items object from a yaml.Node.
func parseItems(node *yaml.Node, ctx *ParseContext) (*openapi20models.Items, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "items must be an object")
	}

	items := &openapi20models.Items{}
	var err error

	// Simple properties - inline
	items.Type = nodeGetString(node, "type")
	items.Format = nodeGetString(node, "format")
	items.CollectionFormat = nodeGetString(node, "collectionFormat")
	items.Default = nodeGetAny(node, "default")
	items.Maximum = nodeGetFloat64Ptr(node, "maximum")
	items.ExclusiveMaximum = nodeGetBool(node, "exclusiveMaximum")
	items.Minimum = nodeGetFloat64Ptr(node, "minimum")
	items.ExclusiveMinimum = nodeGetBool(node, "exclusiveMinimum")
	items.MaxLength = nodeGetUint64Ptr(node, "maxLength")
	items.MinLength = nodeGetUint64Ptr(node, "minLength")
	items.Pattern = nodeGetString(node, "pattern")
	items.MaxItems = nodeGetUint64Ptr(node, "maxItems")
	items.MinItems = nodeGetUint64Ptr(node, "minItems")
	items.UniqueItems = nodeGetBool(node, "uniqueItems")
	items.MultipleOf = nodeGetFloat64Ptr(node, "multipleOf")

	// Enum - array of any values
	if enumNode := nodeGetValue(node, "enum"); enumNode != nil && nodeIsSequence(enumNode) {
		items.Enum = make([]interface{}, len(enumNode.Content))
		for i, item := range enumNode.Content {
			items.Enum[i] = nodeToInterface(item)
		}
	}

	// Nested items - for nested arrays
	if nestedItemsNode := nodeGetValue(node, "items"); nestedItemsNode != nil {
		items.Items, err = parseItems(nestedItemsNode, ctx.push("items"))
		if err != nil {
			return nil, err
		}
	}

	items.VendorExtensions = parseNodeExtensions(node)
	items.NodeSource = ctx.nodeSource(node)

	// Detect unknown fields
	ctx.detectUnknown(node, itemsKnownFieldsSet)

	return items, nil
}
