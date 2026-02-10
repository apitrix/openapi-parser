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

	var err error

	// Body parameter - schema (parsed first for constructor)
	var schema *openapi20models.SchemaRef
	var schemaErr error
	if schemaNode := nodeGetValue(node, "schema"); schemaNode != nil {
		schema, err = parseSchemaRef(schemaNode, ctx.push("schema"))
		if err != nil {
			schemaErr = err
		}
	}

	// Enum - array of any values
	var enum []interface{}
	if enumNode := nodeGetValue(node, "enum"); enumNode != nil && nodeIsSequence(enumNode) {
		enum = make([]interface{}, len(enumNode.Content))
		for i, item := range enumNode.Content {
			enum[i] = nodeToInterface(item)
		}
	}

	// Items - for array type parameters (parsed first for constructor)
	var items *openapi20models.Items
	var itemsErr error
	if itemsNode := nodeGetValue(node, "items"); itemsNode != nil {
		items, err = parseItems(itemsNode, ctx.push("items"))
		if err != nil {
			itemsErr = err
		}
	}

	param := openapi20models.NewParameter(openapi20models.ParameterFields{
		Name:             nodeGetString(node, "name"),
		In:               nodeGetString(node, "in"),
		Description:      nodeGetString(node, "description"),
		Required:         nodeGetBool(node, "required"),
		AllowEmptyValue:  nodeGetBool(node, "allowEmptyValue"),
		Schema:           schema,
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

	if schemaErr != nil {
		param.Trix.Errors = append(param.Trix.Errors, toParseError(schemaErr))
	}
	if itemsErr != nil {
		param.Trix.Errors = append(param.Trix.Errors, toParseError(itemsErr))
	}

	param.VendorExtensions = parseNodeExtensions(node)
	param.Trix.Source = ctx.nodeSource(node)

	// Detect unknown fields
	param.Trix.Errors = append(param.Trix.Errors, unknownFieldParseErrors(ctx.detectUnknown(node, parameterKnownFieldsSet))...)

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
