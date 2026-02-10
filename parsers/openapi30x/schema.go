package openapi30x

import (
	openapi30models "openapi-parser/models/openapi30"

	"gopkg.in/yaml.v3"
)

type schemaParser struct{}

// defaultSchemaParser is the singleton instance used by parsing functions.
var defaultSchemaParser = &schemaParser{}

// parseSharedSchema parses a Schema object from a yaml.Node.
// OpenAPI 3.0.3 spec: https://spec.openapis.org/oas/v3.0.3#schema-object
func parseSharedSchema(node *yaml.Node, ctx *ParseContext) (*openapi30models.Schema, error) {
	return defaultSchemaParser.parse(node, ctx)
}

// Parse parses a Schema object from a yaml.Node.
func (p *schemaParser) parse(node *yaml.Node, ctx *ParseContext) (*openapi30models.Schema, error) {
	if node == nil {
		return nil, nil
	}

	if !nodeIsMapping(node) {
		return nil, ctx.errorAt(node, "schema must be an object")
	}

	var errors []openapi30models.ParseError

	// Build SchemaFields for the constructor
	f := openapi30models.SchemaFields{
		Title:            p.ParseTitle(node),
		Description:      p.ParseDescription(node),
		Type:             p.ParseType(node),
		Format:           p.ParseFormat(node),
		Pattern:          p.ParsePattern(node),
		MultipleOf:       p.ParseMultipleOf(node),
		Maximum:          p.ParseMaximum(node),
		Minimum:          p.ParseMinimum(node),
		ExclusiveMaximum: p.ParseExclusiveMaximum(node),
		ExclusiveMinimum: p.ParseExclusiveMinimum(node),
		MaxLength:        p.ParseMaxLength(node),
		MinLength:        p.ParseMinLength(node),
		MaxItems:         p.ParseMaxItems(node),
		MinItems:         p.ParseMinItems(node),
		MaxProperties:    p.ParseMaxProperties(node),
		MinProperties:    p.ParseMinProperties(node),
		UniqueItems:      p.ParseUniqueItems(node),
		Nullable:         p.ParseNullable(node),
		ReadOnly:         p.ParseReadOnly(node),
		WriteOnly:        p.ParseWriteOnly(node),
		Deprecated:       p.ParseDeprecated(node),
		Required:         p.ParseRequired(node),
		Enum:             p.ParseEnum(node),
		Default:          p.ParseDefault(node),
		Example:          p.ParseExample(node),
	}

	// Complex properties - delegated to dedicated files
	var err error

	f.AllOf, err = p.ParseAllOf(node, ctx)
	if err != nil {
		errors = append(errors, toParseError(err))
	}

	f.OneOf, err = p.ParseOneOf(node, ctx)
	if err != nil {
		errors = append(errors, toParseError(err))
	}

	f.AnyOf, err = p.ParseAnyOf(node, ctx)
	if err != nil {
		errors = append(errors, toParseError(err))
	}

	f.Not, err = p.ParseNot(node, ctx)
	if err != nil {
		errors = append(errors, toParseError(err))
	}

	f.Items, err = p.ParseItems(node, ctx)
	if err != nil {
		errors = append(errors, toParseError(err))
	}

	f.Properties, err = p.ParseProperties(node, ctx)
	if err != nil {
		errors = append(errors, toParseError(err))
	}

	// Additional properties (special handling for bool vs schema)
	addPropsResult, err := p.ParseAdditionalProperties(node, ctx)
	if err != nil {
		errors = append(errors, toParseError(err))
	}
	if addPropsResult != nil {
		f.AdditionalPropertiesAllowed = addPropsResult.Allowed
		f.AdditionalProperties = addPropsResult.SchemaRef
	}

	f.Discriminator, err = p.ParseDiscriminator(node, ctx)
	if err != nil {
		errors = append(errors, toParseError(err))
	}

	f.XML, err = p.ParseXML(node, ctx)
	if err != nil {
		errors = append(errors, toParseError(err))
	}

	f.ExternalDocs, err = p.ParseExternalDocs(node, ctx)
	if err != nil {
		errors = append(errors, toParseError(err))
	}

	// Create via constructor
	schema := openapi30models.NewSchema(f)

	// Parse extensions
	schema.VendorExtensions = parseNodeExtensions(node)

	// Set node source info
	schema.Trix.Source = ctx.nodeSource(node)
	schema.Trix.Errors = append(schema.Trix.Errors, errors...)

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

func (p *schemaParser) ParseType(node *yaml.Node) string {
	return nodeGetString(node, "type")
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

func (p *schemaParser) ParseExclusiveMaximum(node *yaml.Node) bool {
	return nodeGetBool(node, "exclusiveMaximum")
}

func (p *schemaParser) ParseExclusiveMinimum(node *yaml.Node) bool {
	return nodeGetBool(node, "exclusiveMinimum")
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

func (p *schemaParser) ParseNullable(node *yaml.Node) bool {
	return nodeGetBool(node, "nullable")
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
