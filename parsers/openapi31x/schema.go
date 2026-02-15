package openapi31x

import (
	openapi31models "github.com/apitrix/openapi-parser/models/openapi31"
	"github.com/apitrix/openapi-parser/models/shared"

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

	var errs []openapi31models.ParseError

	// Complex properties - parse first
	allOf, err := p.ParseAllOf(node, ctx)
	if err != nil {
		errs = append(errs, toParseError(err))
	}

	oneOf, err := p.ParseOneOf(node, ctx)
	if err != nil {
		errs = append(errs, toParseError(err))
	}

	anyOf, err := p.ParseAnyOf(node, ctx)
	if err != nil {
		errs = append(errs, toParseError(err))
	}

	not, err := p.ParseNot(node, ctx)
	if err != nil {
		errs = append(errs, toParseError(err))
	}

	items, err := p.ParseItems(node, ctx)
	if err != nil {
		errs = append(errs, toParseError(err))
	}

	properties, err := p.ParseProperties(node, ctx)
	if err != nil {
		errs = append(errs, toParseError(err))
	}

	// Additional properties (special handling for bool vs schema)
	var additionalProperties *shared.RefWithMeta[openapi31models.Schema]
	var additionalPropertiesAllowed *bool
	addPropsResult, err := p.ParseAdditionalProperties(node, ctx)
	if err != nil {
		errs = append(errs, toParseError(err))
	}
	if addPropsResult != nil {
		additionalPropertiesAllowed = addPropsResult.Allowed
		additionalProperties = addPropsResult.SchemaRef
	}

	discriminator, err := p.ParseDiscriminator(node, ctx)
	if err != nil {
		errs = append(errs, toParseError(err))
	}

	xml, err := p.ParseXML(node, ctx)
	if err != nil {
		errs = append(errs, toParseError(err))
	}

	externalDocs, err := p.ParseExternalDocs(node, ctx)
	if err != nil {
		errs = append(errs, toParseError(err))
	}

	// JSON Schema 2020-12 new complex properties
	ifSchema, err := p.ParseIf(node, ctx)
	if err != nil {
		errs = append(errs, toParseError(err))
	}

	thenSchema, err := p.ParseThen(node, ctx)
	if err != nil {
		errs = append(errs, toParseError(err))
	}

	elseSchema, err := p.ParseElse(node, ctx)
	if err != nil {
		errs = append(errs, toParseError(err))
	}

	prefixItems, err := p.ParsePrefixItems(node, ctx)
	if err != nil {
		errs = append(errs, toParseError(err))
	}

	dependentSchemas, err := p.ParseDependentSchemas(node, ctx)
	if err != nil {
		errs = append(errs, toParseError(err))
	}

	contentSchema, err := p.ParseContentSchema(node, ctx)
	if err != nil {
		errs = append(errs, toParseError(err))
	}

	unevaluatedItems, err := p.ParseUnevaluatedItems(node, ctx)
	if err != nil {
		errs = append(errs, toParseError(err))
	}

	unevaluatedProperties, err := p.ParseUnevaluatedProperties(node, ctx)
	if err != nil {
		errs = append(errs, toParseError(err))
	}

	// Create via constructor with SchemaFields
	schema := openapi31models.NewSchema(openapi31models.SchemaFields{
		Title:                       p.ParseTitle(node),
		MultipleOf:                  p.ParseMultipleOf(node),
		Maximum:                     p.ParseMaximum(node),
		ExclusiveMaximum:            p.ParseExclusiveMaximum(node),
		Minimum:                     p.ParseMinimum(node),
		ExclusiveMinimum:            p.ParseExclusiveMinimum(node),
		MaxLength:                   p.ParseMaxLength(node),
		MinLength:                   p.ParseMinLength(node),
		Pattern:                     p.ParsePattern(node),
		MaxItems:                    p.ParseMaxItems(node),
		MinItems:                    p.ParseMinItems(node),
		UniqueItems:                 p.ParseUniqueItems(node),
		MaxProperties:               p.ParseMaxProperties(node),
		MinProperties:               p.ParseMinProperties(node),
		Required:                    p.ParseRequired(node),
		Enum:                        p.ParseEnum(node),
		Type:                        p.ParseType(node),
		AllOf:                       allOf,
		OneOf:                       oneOf,
		AnyOf:                       anyOf,
		Not:                         not,
		Items:                       items,
		Properties:                  properties,
		Description:                 p.ParseDescription(node),
		Format:                      p.ParseFormat(node),
		Default:                     p.ParseDefault(node),
		AdditionalProperties:        additionalProperties,
		AdditionalPropertiesAllowed: additionalPropertiesAllowed,
		Const:                       p.ParseConst(node),
		If:                          ifSchema,
		Then:                        thenSchema,
		Else:                        elseSchema,
		DependentSchemas:            dependentSchemas,
		PrefixItems:                 prefixItems,
		Anchor:                      p.ParseAnchor(node),
		DynamicRef:                  p.ParseDynamicRef(node),
		DynamicAnchor:               p.ParseDynamicAnchor(node),
		ContentEncoding:             p.ParseContentEncoding(node),
		ContentMediaType:            p.ParseContentMediaType(node),
		ContentSchema:               contentSchema,
		UnevaluatedItems:            unevaluatedItems,
		UnevaluatedProperties:       unevaluatedProperties,
		Examples:                    p.ParseExamples(node),
		Discriminator:               discriminator,
		ReadOnly:                    p.ParseReadOnly(node),
		WriteOnly:                   p.ParseWriteOnly(node),
		XML:                         xml,
		ExternalDocs:                externalDocs,
		Example:                     p.ParseExample(node),
		Deprecated:                  p.ParseDeprecated(node),
	})

	// Parse extensions
	schema.VendorExtensions = parseNodeExtensions(node)

	// Set node source info
	schema.Trix.Source = ctx.nodeSource(node)
	schema.Trix.Errors = append(schema.Trix.Errors, errs...)

	// Detect unknown fields
	schema.Trix.Errors = append(schema.Trix.Errors, unknownFieldParseErrors(ctx.detectUnknown(node, schemaKnownFieldsSet))...)

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
