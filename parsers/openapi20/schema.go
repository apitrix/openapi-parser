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

	schema := &openapi20models.Schema{}
	var err error

	// Basic metadata
	schema.Title = nodeGetString(node, "title")
	schema.Description = nodeGetString(node, "description")
	schema.Default = nodeGetAny(node, "default")
	schema.Example = nodeGetAny(node, "example")

	// Type constraints
	schema.Type = nodeGetString(node, "type")
	schema.Format = nodeGetString(node, "format")

	// Numeric validation
	schema.MultipleOf = nodeGetFloat64Ptr(node, "multipleOf")
	schema.Maximum = nodeGetFloat64Ptr(node, "maximum")
	schema.ExclusiveMaximum = nodeGetBool(node, "exclusiveMaximum")
	schema.Minimum = nodeGetFloat64Ptr(node, "minimum")
	schema.ExclusiveMinimum = nodeGetBool(node, "exclusiveMinimum")

	// String validation
	schema.MaxLength = nodeGetUint64Ptr(node, "maxLength")
	schema.MinLength = nodeGetUint64Ptr(node, "minLength")
	schema.Pattern = nodeGetString(node, "pattern")

	// Array validation
	schema.MaxItems = nodeGetUint64Ptr(node, "maxItems")
	schema.MinItems = nodeGetUint64Ptr(node, "minItems")
	schema.UniqueItems = nodeGetBool(node, "uniqueItems")

	// Object validation
	schema.MaxProperties = nodeGetUint64Ptr(node, "maxProperties")
	schema.MinProperties = nodeGetUint64Ptr(node, "minProperties")
	schema.Required = nodeGetStringSlice(node, "required")

	// Swagger 2.0 specific
	schema.Discriminator = nodeGetString(node, "discriminator")
	schema.ReadOnly = nodeGetBool(node, "readOnly")

	// Enum - array of any values
	if enumNode := nodeGetValue(node, "enum"); enumNode != nil && nodeIsSequence(enumNode) {
		schema.Enum = make([]interface{}, len(enumNode.Content))
		for i, item := range enumNode.Content {
			schema.Enum[i] = nodeToInterface(item)
		}
	}

	// Complex property - Items
	if itemsNode := nodeGetValue(node, "items"); itemsNode != nil {
		schema.Items, err = parseSchemaRef(itemsNode, ctx.push("items"))
		if err != nil {
			return nil, err
		}
	}

	// Complex property - Properties
	if propsNode := nodeGetValue(node, "properties"); propsNode != nil {
		schema.Properties, err = parseSchemaProperties(propsNode, ctx.push("properties"))
		if err != nil {
			return nil, err
		}
	}

	// Complex property - AdditionalProperties
	if addPropsNode := nodeGetValue(node, "additionalProperties"); addPropsNode != nil {
		// Can be boolean or schema
		if nodeIsScalar(addPropsNode) {
			// Boolean value
			b := nodeGetBool(node, "additionalProperties")
			schema.AdditionalPropertiesAllowed = &b
		} else {
			// Schema reference
			schema.AdditionalProperties, err = parseSchemaRef(addPropsNode, ctx.push("additionalProperties"))
			if err != nil {
				return nil, err
			}
		}
	}

	// Complex property - AllOf
	if allOfNode := nodeGetValue(node, "allOf"); allOfNode != nil {
		schema.AllOf, err = parseSchemaRefs(allOfNode, ctx.push("allOf"))
		if err != nil {
			return nil, err
		}
	}

	// Complex property - XML
	if xmlNode := nodeGetValue(node, "xml"); xmlNode != nil {
		schema.XML, err = parseXML(xmlNode, ctx.push("xml"))
		if err != nil {
			return nil, err
		}
	}

	// Complex property - ExternalDocs
	if edNode := nodeGetValue(node, "externalDocs"); edNode != nil {
		schema.ExternalDocs, err = parseExternalDocs(edNode, ctx.push("externalDocs"))
		if err != nil {
			return nil, err
		}
	}

	schema.VendorExtensions = parseNodeExtensions(node)
	schema.Trix.Source = ctx.nodeSource(node)

	// Detect unknown fields
	ctx.detectUnknown(node, schemaKnownFieldsSet)

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
