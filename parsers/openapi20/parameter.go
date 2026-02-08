package openapi20

import (
	openapi20models "openapi-parser/models/openapi20"

	"gopkg.in/yaml.v3"
)

// parseParameter parses a Parameter object from a yaml.Node.
func parseParameter(node *yaml.Node, ctx *ParseContext) (*openapi20models.Parameter, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "parameter must be an object")
	}

	param := &openapi20models.Parameter{}
	var err error

	// Simple properties - inline
	param.Name = nodeGetString(node, "name")
	param.In = nodeGetString(node, "in")
	param.Description = nodeGetString(node, "description")
	param.Required = nodeGetBool(node, "required")
	param.AllowEmptyValue = nodeGetBool(node, "allowEmptyValue")

	// Body parameter - schema
	if schemaNode := nodeGetValue(node, "schema"); schemaNode != nil {
		param.Schema, err = parseSchemaRef(schemaNode, ctx.push("schema"))
		if err != nil {
			return nil, err
		}
	}

	// Non-body parameters - type info
	param.Type = nodeGetString(node, "type")
	param.Format = nodeGetString(node, "format")
	param.CollectionFormat = nodeGetString(node, "collectionFormat")
	param.Default = nodeGetAny(node, "default")
	param.Maximum = nodeGetFloat64Ptr(node, "maximum")
	param.ExclusiveMaximum = nodeGetBool(node, "exclusiveMaximum")
	param.Minimum = nodeGetFloat64Ptr(node, "minimum")
	param.ExclusiveMinimum = nodeGetBool(node, "exclusiveMinimum")
	param.MaxLength = nodeGetUint64Ptr(node, "maxLength")
	param.MinLength = nodeGetUint64Ptr(node, "minLength")
	param.Pattern = nodeGetString(node, "pattern")
	param.MaxItems = nodeGetUint64Ptr(node, "maxItems")
	param.MinItems = nodeGetUint64Ptr(node, "minItems")
	param.UniqueItems = nodeGetBool(node, "uniqueItems")
	param.MultipleOf = nodeGetFloat64Ptr(node, "multipleOf")

	// Enum - array of any values
	if enumNode := nodeGetValue(node, "enum"); enumNode != nil && nodeIsSequence(enumNode) {
		param.Enum = make([]interface{}, len(enumNode.Content))
		for i, item := range enumNode.Content {
			param.Enum[i] = nodeToInterface(item)
		}
	}

	// Items - for array type parameters
	if itemsNode := nodeGetValue(node, "items"); itemsNode != nil {
		param.Items, err = parseItems(itemsNode, ctx.push("items"))
		if err != nil {
			return nil, err
		}
	}

	param.VendorExtensions = parseNodeExtensions(node)
	param.NodeSource = ctx.nodeSource(node)

	// Detect unknown fields
	ctx.detectUnknown(node, parameterKnownFieldsSet)

	return param, nil
}

// parseParameterRefs parses an array of ParameterRef objects.
func parseParameterRefs(node *yaml.Node, ctx *ParseContext) ([]*openapi20models.ParameterRef, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsSequence(node) {
		return nil, ctx.errorAt(node, "parameters must be an array")
	}

	params := make([]*openapi20models.ParameterRef, 0, len(node.Content))
	for i, itemNode := range node.Content {
		paramRef, err := parseParameterRef(itemNode, ctx.push(itoa(i)))
		if err != nil {
			return nil, err
		}
		params = append(params, paramRef)
	}

	return params, nil
}
