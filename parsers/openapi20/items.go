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

	var err error

	// Enum - array of any values
	var enum []interface{}
	if enumNode := nodeGetValue(node, "enum"); enumNode != nil && nodeIsSequence(enumNode) {
		enum = make([]interface{}, len(enumNode.Content))
		for i, item := range enumNode.Content {
			enum[i] = nodeToInterface(item)
		}
	}

	// Nested items - for nested arrays (parsed first for constructor)
	var nestedItems *openapi20models.Items
	if nestedItemsNode := nodeGetValue(node, "items"); nestedItemsNode != nil {
		nestedItems, err = parseItems(nestedItemsNode, ctx.push("items"))
		if err != nil {
			return nil, err
		}
	}

	items := openapi20models.NewItems(openapi20models.ItemsFields{
		Type:             nodeGetString(node, "type"),
		Format:           nodeGetString(node, "format"),
		Items:            nestedItems,
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

	items.VendorExtensions = parseNodeExtensions(node)
	items.Trix.Source = ctx.nodeSource(node)

	// Detect unknown fields
	items.Trix.Errors = append(items.Trix.Errors, unknownFieldParseErrors(ctx.detectUnknown(node, itemsKnownFieldsSet))...)

	return items, nil
}
