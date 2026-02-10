package openapi20

import (
	openapi20models "openapi-parser/models/openapi20"

	"gopkg.in/yaml.v3"
)

// parseSchema parses a Schema object from a yaml.Node.
func parseSchema(node *yaml.Node, ctx *ParseContext) (*openapi20models.Schema, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "schema must be an object")
	}

	var err error
	var errors []error

	// Enum - array of any values
	var enum []interface{}
	if enumNode := nodeGetValue(node, "enum"); enumNode != nil && nodeIsSequence(enumNode) {
		enum = make([]interface{}, len(enumNode.Content))
		for i, item := range enumNode.Content {
			enum[i] = nodeToInterface(item)
		}
	}

	// Complex property - Items
	var items *openapi20models.SchemaRef
	if itemsNode := nodeGetValue(node, "items"); itemsNode != nil {
		items, err = parseSchemaRef(itemsNode, ctx.push("items"))
		if err != nil {
			errors = append(errors, err)
		}
	}

	// Complex property - Properties
	var properties map[string]*openapi20models.SchemaRef
	if propsNode := nodeGetValue(node, "properties"); propsNode != nil {
		properties, err = parseSchemaProperties(propsNode, ctx.push("properties"))
		if err != nil {
			errors = append(errors, err)
		}
	}

	// Complex property - AdditionalProperties
	var additionalProperties *openapi20models.SchemaRef
	var additionalPropertiesAllowed *bool
	if addPropsNode := nodeGetValue(node, "additionalProperties"); addPropsNode != nil {
		// Can be boolean or schema
		if nodeIsScalar(addPropsNode) {
			// Boolean value
			b := nodeGetBool(node, "additionalProperties")
			additionalPropertiesAllowed = &b
		} else {
			// Schema reference
			additionalProperties, err = parseSchemaRef(addPropsNode, ctx.push("additionalProperties"))
			if err != nil {
				errors = append(errors, err)
			}
		}
	}

	// Complex property - AllOf
	var allOf []*openapi20models.SchemaRef
	if allOfNode := nodeGetValue(node, "allOf"); allOfNode != nil {
		allOf, err = parseSchemaRefs(allOfNode, ctx.push("allOf"))
		if err != nil {
			errors = append(errors, err)
		}
	}

	// Complex property - XML
	var xml *openapi20models.XML
	if xmlNode := nodeGetValue(node, "xml"); xmlNode != nil {
		xml, err = parseXML(xmlNode, ctx.push("xml"))
		if err != nil {
			errors = append(errors, err)
		}
	}

	// Complex property - ExternalDocs
	var externalDocs *openapi20models.ExternalDocs
	if edNode := nodeGetValue(node, "externalDocs"); edNode != nil {
		externalDocs, err = parseExternalDocs(edNode, ctx.push("externalDocs"))
		if err != nil {
			errors = append(errors, err)
		}
	}

	schema := openapi20models.NewSchema(openapi20models.SchemaFields{
		Title:                       nodeGetString(node, "title"),
		Description:                 nodeGetString(node, "description"),
		Default:                     nodeGetAny(node, "default"),
		MultipleOf:                  nodeGetFloat64Ptr(node, "multipleOf"),
		Maximum:                     nodeGetFloat64Ptr(node, "maximum"),
		ExclusiveMaximum:            nodeGetBool(node, "exclusiveMaximum"),
		Minimum:                     nodeGetFloat64Ptr(node, "minimum"),
		ExclusiveMinimum:            nodeGetBool(node, "exclusiveMinimum"),
		MaxLength:                   nodeGetUint64Ptr(node, "maxLength"),
		MinLength:                   nodeGetUint64Ptr(node, "minLength"),
		Pattern:                     nodeGetString(node, "pattern"),
		MaxItems:                    nodeGetUint64Ptr(node, "maxItems"),
		MinItems:                    nodeGetUint64Ptr(node, "minItems"),
		UniqueItems:                 nodeGetBool(node, "uniqueItems"),
		MaxProperties:               nodeGetUint64Ptr(node, "maxProperties"),
		MinProperties:               nodeGetUint64Ptr(node, "minProperties"),
		Required:                    nodeGetStringSlice(node, "required"),
		Enum:                        enum,
		Type:                        nodeGetString(node, "type"),
		Format:                      nodeGetString(node, "format"),
		AllOf:                       allOf,
		Items:                       items,
		Properties:                  properties,
		AdditionalProperties:        additionalProperties,
		AdditionalPropertiesAllowed: additionalPropertiesAllowed,
		Discriminator:               nodeGetString(node, "discriminator"),
		ReadOnly:                    nodeGetBool(node, "readOnly"),
		XML:                         xml,
		ExternalDocs:                externalDocs,
		Example:                     nodeGetAny(node, "example"),
	})

	for _, e := range errors {
		schema.Trix.Errors = append(schema.Trix.Errors, toParseError(e))
	}

	schema.VendorExtensions = parseNodeExtensions(node)
	schema.Trix.Source = ctx.nodeSource(node)

	// Detect unknown fields
	schema.Trix.Errors = append(schema.Trix.Errors, unknownFieldParseErrors(ctx.detectUnknown(node, schemaKnownFieldsSet))...)

	return schema, nil
}

// parseSchemaProperties parses a map of schema properties.
func parseSchemaProperties(node *yaml.Node, ctx *ParseContext) (map[string]*openapi20models.SchemaRef, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "properties must be an object")
	}

	props := make(map[string]*openapi20models.SchemaRef)

	for key, propNode := range nodeMapPairs(node) {
		propRef, err := parseSchemaRef(propNode, ctx.push(key))
		if err != nil {
			return nil, err
		}
		props[key] = propRef
	}

	return props, nil
}

// parseSchemaRefs parses an array of SchemaRef objects.
func parseSchemaRefs(node *yaml.Node, ctx *ParseContext) ([]*openapi20models.SchemaRef, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsSequence(node) {
		return nil, ctx.errorAt(node, "must be an array of schemas")
	}

	refs := make([]*openapi20models.SchemaRef, 0, len(node.Content))
	for i, itemNode := range node.Content {
		ref, err := parseSchemaRef(itemNode, ctx.push(itoa(i)))
		if err != nil {
			return nil, err
		}
		refs = append(refs, ref)
	}

	return refs, nil
}
