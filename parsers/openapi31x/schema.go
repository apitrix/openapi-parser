package openapi31x

import (
	openapi31models "openapi-parser/models/openapi31"

	"gopkg.in/yaml.v3"
)

type schemaParser struct{}

// defaultSchemaParser is the singleton instance used by parsing functions.
var defaultSchemaParser = &schemaParser{}

// parseSharedSchema parses a Schema object from a yaml.Node.
// OpenAPI 3.1.0 spec: https://spec.openapis.org/oas/v3.1.0#schema-object
// Uses JSON Schema Draft 2020-12.
func parseSharedSchema(node *yaml.Node, ctx *ParseContext) (*openapi31models.Schema, error) {
	return defaultSchemaParser.parse(node, ctx)
}

// Parse parses a Schema object from a yaml.Node.
func (p *schemaParser) parse(node *yaml.Node, ctx *ParseContext) (*openapi31models.Schema, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "schema must be an object")
	}

	schema := &openapi31models.Schema{}
	var err error

	// Simple properties - inline
	schema.Title = p.ParseTitle(node)
	schema.Description = p.ParseDescription(node)
	schema.Type = p.ParseType(node) // Returns SchemaType (string or []string)
	schema.Format = p.ParseFormat(node)
	schema.Pattern = p.ParsePattern(node)
	schema.MultipleOf = p.ParseMultipleOf(node)
	schema.Maximum = p.ParseMaximum(node)
	schema.Minimum = p.ParseMinimum(node)
	schema.ExclusiveMaximum = p.ParseExclusiveMaximum(node) // *float64 in 3.1 (was bool in 3.0)
	schema.ExclusiveMinimum = p.ParseExclusiveMinimum(node) // *float64 in 3.1 (was bool in 3.0)
	schema.MaxLength = p.ParseMaxLength(node)
	schema.MinLength = p.ParseMinLength(node)
	schema.MaxItems = p.ParseMaxItems(node)
	schema.MinItems = p.ParseMinItems(node)
	schema.MaxProperties = p.ParseMaxProperties(node)
	schema.MinProperties = p.ParseMinProperties(node)
	schema.UniqueItems = p.ParseUniqueItems(node)
	schema.ReadOnly = p.ParseReadOnly(node)
	schema.WriteOnly = p.ParseWriteOnly(node)
	schema.Deprecated = p.ParseDeprecated(node)
	schema.Required = p.ParseRequired(node)
	schema.Enum = p.ParseEnum(node)
	schema.Default = p.ParseDefault(node)
	schema.Example = p.ParseExample(node)

	// JSON Schema 2020-12 new simple properties
	schema.Const = p.ParseConst(node)
	schema.Anchor = p.ParseAnchor(node)
	schema.DynamicRef = p.ParseDynamicRef(node)
	schema.DynamicAnchor = p.ParseDynamicAnchor(node)
	schema.ContentEncoding = p.ParseContentEncoding(node)
	schema.ContentMediaType = p.ParseContentMediaType(node)
	schema.Examples = p.ParseExamples(node)

	// Complex properties - delegated to dedicated files
	schema.AllOf, err = p.ParseAllOf(node, ctx)
	if err != nil {
		return nil, err
	}

	schema.OneOf, err = p.ParseOneOf(node, ctx)
	if err != nil {
		return nil, err
	}

	schema.AnyOf, err = p.ParseAnyOf(node, ctx)
	if err != nil {
		return nil, err
	}

	schema.Not, err = p.ParseNot(node, ctx)
	if err != nil {
		return nil, err
	}

	schema.Items, err = p.ParseItems(node, ctx)
	if err != nil {
		return nil, err
	}

	schema.Properties, err = p.ParseProperties(node, ctx)
	if err != nil {
		return nil, err
	}

	// Additional properties (special handling for bool vs schema)
	addPropsResult, err := p.ParseAdditionalProperties(node, ctx)
	if err != nil {
		return nil, err
	}
	if addPropsResult != nil {
		schema.AdditionalPropertiesAllowed = addPropsResult.Allowed
		schema.AdditionalProperties = addPropsResult.SchemaRef
	}

	schema.Discriminator, err = p.ParseDiscriminator(node, ctx)
	if err != nil {
		return nil, err
	}

	schema.XML, err = p.ParseXML(node, ctx)
	if err != nil {
		return nil, err
	}

	schema.ExternalDocs, err = p.ParseExternalDocs(node, ctx)
	if err != nil {
		return nil, err
	}

	// JSON Schema 2020-12 new complex properties
	schema.If, err = p.ParseIf(node, ctx)
	if err != nil {
		return nil, err
	}

	schema.Then, err = p.ParseThen(node, ctx)
	if err != nil {
		return nil, err
	}

	schema.Else, err = p.ParseElse(node, ctx)
	if err != nil {
		return nil, err
	}

	schema.PrefixItems, err = p.ParsePrefixItems(node, ctx)
	if err != nil {
		return nil, err
	}

	schema.DependentSchemas, err = p.ParseDependentSchemas(node, ctx)
	if err != nil {
		return nil, err
	}

	schema.ContentSchema, err = p.ParseContentSchema(node, ctx)
	if err != nil {
		return nil, err
	}

	schema.UnevaluatedItems, err = p.ParseUnevaluatedItems(node, ctx)
	if err != nil {
		return nil, err
	}

	schema.UnevaluatedProperties, err = p.ParseUnevaluatedProperties(node, ctx)
	if err != nil {
		return nil, err
	}

	// Parse extensions
	schema.VendorExtensions = parseNodeExtensions(node)

	// Set node source info
	schema.NodeSource = ctx.nodeSource(node)

	// Detect unknown fields
	ctx.detectUnknown(node, schemaKnownFieldsSet)

	return schema, nil
}

// ============================================================================
// Simple property parsers - all inline in this file
// ============================================================================

func (p *schemaParser) ParseTitle(node *yaml.Node) string {
	return nodeGetString(node, "title")
}

func (p *schemaParser) ParseDescription(node *yaml.Node) string {
	return nodeGetString(node, "description")
}

// ParseType parses the "type" field which can be a string or an array of strings in 3.1.
func (p *schemaParser) ParseType(node *yaml.Node) openapi31models.SchemaType {
	typeNode := nodeGetValue(node, "type")
	if typeNode == nil {
		return openapi31models.SchemaType{}
	}

	if nodeIsScalar(typeNode) {
		return openapi31models.SchemaType{Single: typeNode.Value}
	}

	if nodeIsSequence(typeNode) {
		arr := make([]string, 0, len(typeNode.Content))
		for _, child := range typeNode.Content {
			if nodeIsScalar(child) {
				arr = append(arr, child.Value)
			}
		}
		return openapi31models.SchemaType{Array: arr}
	}

	return openapi31models.SchemaType{}
}

func (p *schemaParser) ParseFormat(node *yaml.Node) string {
	return nodeGetString(node, "format")
}

func (p *schemaParser) ParsePattern(node *yaml.Node) string {
	return nodeGetString(node, "pattern")
}

func (p *schemaParser) ParseMultipleOf(node *yaml.Node) *float64 {
	return nodeGetFloat64Ptr(node, "multipleOf")
}

func (p *schemaParser) ParseMaximum(node *yaml.Node) *float64 {
	return nodeGetFloat64Ptr(node, "maximum")
}

func (p *schemaParser) ParseMinimum(node *yaml.Node) *float64 {
	return nodeGetFloat64Ptr(node, "minimum")
}

// ParseExclusiveMaximum returns *float64 (JSON Schema 2020-12, was bool in 3.0).
func (p *schemaParser) ParseExclusiveMaximum(node *yaml.Node) *float64 {
	return nodeGetFloat64Ptr(node, "exclusiveMaximum")
}

// ParseExclusiveMinimum returns *float64 (JSON Schema 2020-12, was bool in 3.0).
func (p *schemaParser) ParseExclusiveMinimum(node *yaml.Node) *float64 {
	return nodeGetFloat64Ptr(node, "exclusiveMinimum")
}

func (p *schemaParser) ParseMaxLength(node *yaml.Node) *uint64 {
	return nodeGetUint64Ptr(node, "maxLength")
}

func (p *schemaParser) ParseMinLength(node *yaml.Node) *uint64 {
	return nodeGetUint64Ptr(node, "minLength")
}

func (p *schemaParser) ParseMaxItems(node *yaml.Node) *uint64 {
	return nodeGetUint64Ptr(node, "maxItems")
}

func (p *schemaParser) ParseMinItems(node *yaml.Node) *uint64 {
	return nodeGetUint64Ptr(node, "minItems")
}

func (p *schemaParser) ParseMaxProperties(node *yaml.Node) *uint64 {
	return nodeGetUint64Ptr(node, "maxProperties")
}

func (p *schemaParser) ParseMinProperties(node *yaml.Node) *uint64 {
	return nodeGetUint64Ptr(node, "minProperties")
}

func (p *schemaParser) ParseUniqueItems(node *yaml.Node) bool {
	return nodeGetBool(node, "uniqueItems")
}

func (p *schemaParser) ParseReadOnly(node *yaml.Node) bool {
	return nodeGetBool(node, "readOnly")
}

func (p *schemaParser) ParseWriteOnly(node *yaml.Node) bool {
	return nodeGetBool(node, "writeOnly")
}

func (p *schemaParser) ParseDeprecated(node *yaml.Node) bool {
	return nodeGetBool(node, "deprecated")
}

func (p *schemaParser) ParseRequired(node *yaml.Node) []string {
	return nodeGetStringSlice(node, "required")
}

func (p *schemaParser) ParseEnum(node *yaml.Node) []interface{} {
	enumNode := nodeGetValue(node, "enum")
	if enumNode == nil || !nodeIsSequence(enumNode) {
		return nil
	}
	result := make([]interface{}, len(enumNode.Content))
	for i, child := range enumNode.Content {
		result[i] = nodeToInterface(child)
	}
	return result
}

func (p *schemaParser) ParseDefault(node *yaml.Node) interface{} {
	return nodeGetAny(node, "default")
}

func (p *schemaParser) ParseExample(node *yaml.Node) interface{} {
	return nodeGetAny(node, "example")
}

// JSON Schema 2020-12 new simple property parsers

func (p *schemaParser) ParseConst(node *yaml.Node) interface{} {
	return nodeGetAny(node, "const")
}

func (p *schemaParser) ParseAnchor(node *yaml.Node) string {
	return nodeGetString(node, "$anchor")
}

func (p *schemaParser) ParseDynamicRef(node *yaml.Node) string {
	return nodeGetString(node, "$dynamicRef")
}

func (p *schemaParser) ParseDynamicAnchor(node *yaml.Node) string {
	return nodeGetString(node, "$dynamicAnchor")
}

func (p *schemaParser) ParseContentEncoding(node *yaml.Node) string {
	return nodeGetString(node, "contentEncoding")
}

func (p *schemaParser) ParseContentMediaType(node *yaml.Node) string {
	return nodeGetString(node, "contentMediaType")
}

// ParseExamples parses the "examples" array (JSON Schema 2020-12).
func (p *schemaParser) ParseExamples(node *yaml.Node) []interface{} {
	examplesNode := nodeGetValue(node, "examples")
	if examplesNode == nil || !nodeIsSequence(examplesNode) {
		return nil
	}
	result := make([]interface{}, len(examplesNode.Content))
	for i, child := range examplesNode.Content {
		result[i] = nodeToInterface(child)
	}
	return result
}
