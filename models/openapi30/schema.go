package openapi30

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// Schema represents the OpenAPI 3.0 Schema Object.
// https://spec.openapis.org/oas/v3.0.3#schema-object
type Schema struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	// JSON Schema fields
	title            string
	multipleOf       *float64
	maximum          *float64
	exclusiveMaximum bool
	minimum          *float64
	exclusiveMinimum bool
	maxLength        *uint64
	minLength        *uint64
	pattern          string
	maxItems         *uint64
	minItems         *uint64
	uniqueItems      bool
	maxProperties    *uint64
	minProperties    *uint64
	required         []string
	enum             []interface{}
	schemaType       string
	allOf            []*shared.Ref[Schema]
	oneOf            []*shared.Ref[Schema]
	anyOf            []*shared.Ref[Schema]
	not              *shared.Ref[Schema]
	items            *shared.Ref[Schema]
	properties       map[string]*shared.Ref[Schema]
	description      string
	format           string
	defaultVal       interface{}

	// AdditionalProperties can be a boolean or a schema.
	additionalProperties        *shared.Ref[Schema]
	additionalPropertiesAllowed *bool

	// OpenAPI 3.0 specific fields
	nullable      bool
	discriminator *Discriminator
	readOnly      bool
	writeOnly     bool
	xml           *XML
	externalDocs  *ExternalDocumentation
	example       interface{}
	deprecated    bool
}

func (s *Schema) Title() string                              { return s.title }
func (s *Schema) MultipleOf() *float64                       { return s.multipleOf }
func (s *Schema) Maximum() *float64                          { return s.maximum }
func (s *Schema) ExclusiveMaximum() bool                     { return s.exclusiveMaximum }
func (s *Schema) Minimum() *float64                          { return s.minimum }
func (s *Schema) ExclusiveMinimum() bool                     { return s.exclusiveMinimum }
func (s *Schema) MaxLength() *uint64                         { return s.maxLength }
func (s *Schema) MinLength() *uint64                         { return s.minLength }
func (s *Schema) Pattern() string                            { return s.pattern }
func (s *Schema) MaxItems() *uint64                          { return s.maxItems }
func (s *Schema) MinItems() *uint64                          { return s.minItems }
func (s *Schema) UniqueItems() bool                          { return s.uniqueItems }
func (s *Schema) MaxProperties() *uint64                     { return s.maxProperties }
func (s *Schema) MinProperties() *uint64                     { return s.minProperties }
func (s *Schema) Required() []string                         { return s.required }
func (s *Schema) Enum() []interface{}                        { return s.enum }
func (s *Schema) Type() string                               { return s.schemaType }
func (s *Schema) AllOf() []*shared.Ref[Schema]               { return s.allOf }
func (s *Schema) OneOf() []*shared.Ref[Schema]               { return s.oneOf }
func (s *Schema) AnyOf() []*shared.Ref[Schema]               { return s.anyOf }
func (s *Schema) Not() *shared.Ref[Schema]                   { return s.not }
func (s *Schema) Items() *shared.Ref[Schema]                 { return s.items }
func (s *Schema) Properties() map[string]*shared.Ref[Schema] { return s.properties }
func (s *Schema) Description() string                        { return s.description }
func (s *Schema) Format() string                             { return s.format }
func (s *Schema) Default() interface{}                       { return s.defaultVal }
func (s *Schema) AdditionalProperties() *shared.Ref[Schema]  { return s.additionalProperties }
func (s *Schema) AdditionalPropertiesAllowed() *bool         { return s.additionalPropertiesAllowed }
func (s *Schema) Nullable() bool                             { return s.nullable }
func (s *Schema) Discriminator() *Discriminator              { return s.discriminator }
func (s *Schema) ReadOnly() bool                             { return s.readOnly }
func (s *Schema) WriteOnly() bool                            { return s.writeOnly }
func (s *Schema) XML() *XML                                  { return s.xml }
func (s *Schema) ExternalDocs() *ExternalDocumentation       { return s.externalDocs }
func (s *Schema) Example() interface{}                       { return s.example }
func (s *Schema) Deprecated() bool                           { return s.deprecated }

// NewSchema creates a new Schema instance.
// Due to the large number of fields, callers should use NewSchemaFields.
func NewSchema(f SchemaFields) *Schema {
	return &Schema{
		title: f.Title, multipleOf: f.MultipleOf, maximum: f.Maximum,
		exclusiveMaximum: f.ExclusiveMaximum, minimum: f.Minimum,
		exclusiveMinimum: f.ExclusiveMinimum, maxLength: f.MaxLength,
		minLength: f.MinLength, pattern: f.Pattern, maxItems: f.MaxItems,
		minItems: f.MinItems, uniqueItems: f.UniqueItems,
		maxProperties: f.MaxProperties, minProperties: f.MinProperties,
		required: f.Required, enum: f.Enum, schemaType: f.Type,
		allOf: f.AllOf, oneOf: f.OneOf, anyOf: f.AnyOf,
		not: f.Not, items: f.Items, properties: f.Properties,
		description: f.Description, format: f.Format, defaultVal: f.Default,
		additionalProperties:        f.AdditionalProperties,
		additionalPropertiesAllowed: f.AdditionalPropertiesAllowed,
		nullable:                    f.Nullable, discriminator: f.Discriminator,
		readOnly: f.ReadOnly, writeOnly: f.WriteOnly, xml: f.XML,
		externalDocs: f.ExternalDocs, example: f.Example,
		deprecated: f.Deprecated,
	}
}

// SchemaFields holds all fields for constructing a Schema.
// Using a struct avoids a 44-parameter constructor.
type SchemaFields struct {
	Title            string
	MultipleOf       *float64
	Maximum          *float64
	ExclusiveMaximum bool
	Minimum          *float64
	ExclusiveMinimum bool
	MaxLength        *uint64
	MinLength        *uint64
	Pattern          string
	MaxItems         *uint64
	MinItems         *uint64
	UniqueItems      bool
	MaxProperties    *uint64
	MinProperties    *uint64
	Required         []string
	Enum             []interface{}
	Type             string
	AllOf            []*shared.Ref[Schema]
	OneOf            []*shared.Ref[Schema]
	AnyOf            []*shared.Ref[Schema]
	Not              *shared.Ref[Schema]
	Items            *shared.Ref[Schema]
	Properties       map[string]*shared.Ref[Schema]
	Description      string
	Format           string
	Default          interface{}

	AdditionalProperties        *shared.Ref[Schema]
	AdditionalPropertiesAllowed *bool

	Nullable      bool
	Discriminator *Discriminator
	ReadOnly      bool
	WriteOnly     bool
	XML           *XML
	ExternalDocs  *ExternalDocumentation
	Example       interface{}
	Deprecated    bool
}

func (s *Schema) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "title", Value: s.title},
		{Key: "multipleOf", Value: s.multipleOf},
		{Key: "maximum", Value: s.maximum},
		{Key: "exclusiveMaximum", Value: s.exclusiveMaximum},
		{Key: "minimum", Value: s.minimum},
		{Key: "exclusiveMinimum", Value: s.exclusiveMinimum},
		{Key: "maxLength", Value: s.maxLength},
		{Key: "minLength", Value: s.minLength},
		{Key: "pattern", Value: s.pattern},
		{Key: "maxItems", Value: s.maxItems},
		{Key: "minItems", Value: s.minItems},
		{Key: "uniqueItems", Value: s.uniqueItems},
		{Key: "maxProperties", Value: s.maxProperties},
		{Key: "minProperties", Value: s.minProperties},
		{Key: "required", Value: s.required},
		{Key: "enum", Value: s.enum},
		{Key: "type", Value: s.schemaType},
		{Key: "allOf", Value: s.allOf},
		{Key: "oneOf", Value: s.oneOf},
		{Key: "anyOf", Value: s.anyOf},
		{Key: "not", Value: s.not},
		{Key: "items", Value: s.items},
		{Key: "properties", Value: s.properties},
		{Key: "description", Value: s.description},
		{Key: "format", Value: s.format},
		{Key: "default", Value: s.defaultVal},
	}

	// additionalProperties: either a boolean or a schema reference
	if s.additionalPropertiesAllowed != nil {
		fields = append(fields, shared.Field{Key: "additionalProperties", Value: s.additionalPropertiesAllowed})
	} else if s.additionalProperties != nil {
		fields = append(fields, shared.Field{Key: "additionalProperties", Value: s.additionalProperties})
	}

	fields = append(fields,
		shared.Field{Key: "nullable", Value: s.nullable},
		shared.Field{Key: "discriminator", Value: s.discriminator},
		shared.Field{Key: "readOnly", Value: s.readOnly},
		shared.Field{Key: "writeOnly", Value: s.writeOnly},
		shared.Field{Key: "xml", Value: s.xml},
		shared.Field{Key: "externalDocs", Value: s.externalDocs},
		shared.Field{Key: "example", Value: s.example},
		shared.Field{Key: "deprecated", Value: s.deprecated},
	)

	return shared.AppendExtensions(fields, s.VendorExtensions)
}

func (s *Schema) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(s.marshalFields())
}

func (s *Schema) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(s.marshalFields())
}

var _ yaml.Marshaler = (*Schema)(nil)
