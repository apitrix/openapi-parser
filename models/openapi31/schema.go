package openapi31

import (
	"openapi-parser/models/shared"
	"sort"

	"gopkg.in/yaml.v3"
)

// Schema represents the OpenAPI 3.1 Schema Object.
// Based on JSON Schema Draft 2020-12.
// https://spec.openapis.org/oas/v3.1.0#schema-object
type Schema struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	// JSON Schema fields
	title            string
	multipleOf       *float64
	maximum          *float64
	exclusiveMaximum *float64
	minimum          *float64
	exclusiveMinimum *float64
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
	schemaType       SchemaType
	allOf            []*shared.RefWithMeta[Schema]
	oneOf            []*shared.RefWithMeta[Schema]
	anyOf            []*shared.RefWithMeta[Schema]
	not              *shared.RefWithMeta[Schema]
	items            *shared.RefWithMeta[Schema]
	properties       map[string]*shared.RefWithMeta[Schema]
	description      string
	format           string
	defaultVal       interface{}

	// AdditionalProperties can be a boolean or a schema.
	additionalProperties        *shared.RefWithMeta[Schema]
	additionalPropertiesAllowed *bool

	// JSON Schema 2020-12 new keywords
	constVal              interface{}
	ifSchema              *shared.RefWithMeta[Schema]
	thenSchema            *shared.RefWithMeta[Schema]
	elseSchema            *shared.RefWithMeta[Schema]
	dependentSchemas      map[string]*shared.RefWithMeta[Schema]
	prefixItems           []*shared.RefWithMeta[Schema]
	anchor                string
	dynamicRef            string
	dynamicAnchor         string
	contentEncoding       string
	contentMediaType      string
	contentSchema         *shared.RefWithMeta[Schema]
	unevaluatedItems      *shared.RefWithMeta[Schema]
	unevaluatedProperties *shared.RefWithMeta[Schema]
	examples              []interface{}

	// OpenAPI extensions (still present in 3.1)
	discriminator *Discriminator
	readOnly      bool
	writeOnly     bool
	xml           *XML
	externalDocs  *ExternalDocumentation
	example       interface{}
	deprecated    bool
}

func (s *Schema) Title() string                                            { return s.title }
func (s *Schema) MultipleOf() *float64                                     { return s.multipleOf }
func (s *Schema) Maximum() *float64                                        { return s.maximum }
func (s *Schema) ExclusiveMaximum() *float64                               { return s.exclusiveMaximum }
func (s *Schema) Minimum() *float64                                        { return s.minimum }
func (s *Schema) ExclusiveMinimum() *float64                               { return s.exclusiveMinimum }
func (s *Schema) MaxLength() *uint64                                       { return s.maxLength }
func (s *Schema) MinLength() *uint64                                       { return s.minLength }
func (s *Schema) Pattern() string                                          { return s.pattern }
func (s *Schema) MaxItems() *uint64                                        { return s.maxItems }
func (s *Schema) MinItems() *uint64                                        { return s.minItems }
func (s *Schema) UniqueItems() bool                                        { return s.uniqueItems }
func (s *Schema) MaxProperties() *uint64                                   { return s.maxProperties }
func (s *Schema) MinProperties() *uint64                                   { return s.minProperties }
func (s *Schema) Required() []string                                       { return s.required }
func (s *Schema) Enum() []interface{}                                      { return s.enum }
func (s *Schema) Type() SchemaType                                         { return s.schemaType }
func (s *Schema) AllOf() []*shared.RefWithMeta[Schema]                     { return s.allOf }
func (s *Schema) OneOf() []*shared.RefWithMeta[Schema]                     { return s.oneOf }
func (s *Schema) AnyOf() []*shared.RefWithMeta[Schema]                     { return s.anyOf }
func (s *Schema) Not() *shared.RefWithMeta[Schema]                         { return s.not }
func (s *Schema) Items() *shared.RefWithMeta[Schema]                       { return s.items }
func (s *Schema) Properties() map[string]*shared.RefWithMeta[Schema]       { return s.properties }
func (s *Schema) Description() string                                      { return s.description }
func (s *Schema) Format() string                                           { return s.format }
func (s *Schema) Default() interface{}                                     { return s.defaultVal }
func (s *Schema) AdditionalProperties() *shared.RefWithMeta[Schema]        { return s.additionalProperties }
func (s *Schema) AdditionalPropertiesAllowed() *bool                       { return s.additionalPropertiesAllowed }
func (s *Schema) Const() interface{}                                       { return s.constVal }
func (s *Schema) If() *shared.RefWithMeta[Schema]                          { return s.ifSchema }
func (s *Schema) Then() *shared.RefWithMeta[Schema]                        { return s.thenSchema }
func (s *Schema) Else() *shared.RefWithMeta[Schema]                        { return s.elseSchema }
func (s *Schema) DependentSchemas() map[string]*shared.RefWithMeta[Schema] { return s.dependentSchemas }
func (s *Schema) PrefixItems() []*shared.RefWithMeta[Schema]               { return s.prefixItems }
func (s *Schema) Anchor() string                                           { return s.anchor }
func (s *Schema) DynamicRef() string                                       { return s.dynamicRef }
func (s *Schema) DynamicAnchor() string                                    { return s.dynamicAnchor }
func (s *Schema) ContentEncoding() string                                  { return s.contentEncoding }
func (s *Schema) ContentMediaType() string                                 { return s.contentMediaType }
func (s *Schema) ContentSchema() *shared.RefWithMeta[Schema]               { return s.contentSchema }
func (s *Schema) UnevaluatedItems() *shared.RefWithMeta[Schema]            { return s.unevaluatedItems }
func (s *Schema) UnevaluatedProperties() *shared.RefWithMeta[Schema]       { return s.unevaluatedProperties }
func (s *Schema) Examples() []interface{}                                  { return s.examples }
func (s *Schema) Discriminator() *Discriminator                            { return s.discriminator }
func (s *Schema) ReadOnly() bool                                           { return s.readOnly }
func (s *Schema) WriteOnly() bool                                          { return s.writeOnly }
func (s *Schema) XML() *XML                                                { return s.xml }
func (s *Schema) ExternalDocs() *ExternalDocumentation                     { return s.externalDocs }
func (s *Schema) Example() interface{}                                     { return s.example }
func (s *Schema) Deprecated() bool                                         { return s.deprecated }

// NewSchema creates a new Schema instance.
// Due to the large number of fields, callers should use SchemaFields.
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
		constVal:                    f.Const, ifSchema: f.If, thenSchema: f.Then,
		elseSchema: f.Else, dependentSchemas: f.DependentSchemas,
		prefixItems: f.PrefixItems, anchor: f.Anchor,
		dynamicRef: f.DynamicRef, dynamicAnchor: f.DynamicAnchor,
		contentEncoding: f.ContentEncoding, contentMediaType: f.ContentMediaType,
		contentSchema:    f.ContentSchema,
		unevaluatedItems: f.UnevaluatedItems, unevaluatedProperties: f.UnevaluatedProperties,
		examples:      f.Examples,
		discriminator: f.Discriminator, readOnly: f.ReadOnly, writeOnly: f.WriteOnly,
		xml: f.XML, externalDocs: f.ExternalDocs, example: f.Example,
		deprecated: f.Deprecated,
	}
}

// SchemaFields holds all fields for constructing a Schema.
// Using a struct avoids a 55+ parameter constructor.
type SchemaFields struct {
	Title            string
	MultipleOf       *float64
	Maximum          *float64
	ExclusiveMaximum *float64
	Minimum          *float64
	ExclusiveMinimum *float64
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
	Type             SchemaType
	AllOf            []*shared.RefWithMeta[Schema]
	OneOf            []*shared.RefWithMeta[Schema]
	AnyOf            []*shared.RefWithMeta[Schema]
	Not              *shared.RefWithMeta[Schema]
	Items            *shared.RefWithMeta[Schema]
	Properties       map[string]*shared.RefWithMeta[Schema]
	Description      string
	Format           string
	Default          interface{}

	AdditionalProperties        *shared.RefWithMeta[Schema]
	AdditionalPropertiesAllowed *bool

	// JSON Schema 2020-12
	Const                 interface{}
	If                    *shared.RefWithMeta[Schema]
	Then                  *shared.RefWithMeta[Schema]
	Else                  *shared.RefWithMeta[Schema]
	DependentSchemas      map[string]*shared.RefWithMeta[Schema]
	PrefixItems           []*shared.RefWithMeta[Schema]
	Anchor                string
	DynamicRef            string
	DynamicAnchor         string
	ContentEncoding       string
	ContentMediaType      string
	ContentSchema         *shared.RefWithMeta[Schema]
	UnevaluatedItems      *shared.RefWithMeta[Schema]
	UnevaluatedProperties *shared.RefWithMeta[Schema]
	Examples              []interface{}

	// OpenAPI extensions
	Discriminator *Discriminator
	ReadOnly      bool
	WriteOnly     bool
	XML           *XML
	ExternalDocs  *ExternalDocumentation
	Example       interface{}
	Deprecated    bool
}

func (s *Schema) marshalFields() []shared.Field {
	// additionalProperties: bool pointer or schema ref
	var addProps interface{}
	if s.additionalPropertiesAllowed != nil {
		addProps = s.additionalPropertiesAllowed
	} else if s.additionalProperties != nil {
		addProps = s.additionalProperties
	}

	// properties: sorted keys
	var propsFields []shared.Field
	if len(s.properties) > 0 {
		keys := make([]string, 0, len(s.properties))
		for k := range s.properties {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		propsFields = make([]shared.Field, 0, len(keys))
		for _, k := range keys {
			propsFields = append(propsFields, shared.Field{Key: k, Value: s.properties[k]})
		}
	}

	// dependentSchemas: sorted keys
	var depSchemaFields []shared.Field
	if len(s.dependentSchemas) > 0 {
		keys := make([]string, 0, len(s.dependentSchemas))
		for k := range s.dependentSchemas {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		depSchemaFields = make([]shared.Field, 0, len(keys))
		for _, k := range keys {
			depSchemaFields = append(depSchemaFields, shared.Field{Key: k, Value: s.dependentSchemas[k]})
		}
	}

	// type: use SchemaType's marshal (string or array)
	var schemaTypeVal interface{}
	if !s.schemaType.IsEmpty() {
		schemaTypeVal = s.schemaType
	}

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
		{Key: "type", Value: schemaTypeVal},
		{Key: "allOf", Value: s.allOf},
		{Key: "oneOf", Value: s.oneOf},
		{Key: "anyOf", Value: s.anyOf},
		{Key: "not", Value: s.not},
		{Key: "items", Value: s.items},
		{Key: "prefixItems", Value: s.prefixItems},
		{Key: "properties", Value: propsFields},
		{Key: "additionalProperties", Value: addProps},
		{Key: "description", Value: s.description},
		{Key: "format", Value: s.format},
		{Key: "default", Value: s.defaultVal},
		{Key: "const", Value: s.constVal},
		{Key: "if", Value: s.ifSchema},
		{Key: "then", Value: s.thenSchema},
		{Key: "else", Value: s.elseSchema},
		{Key: "dependentSchemas", Value: depSchemaFields},
		{Key: "$anchor", Value: s.anchor},
		{Key: "$dynamicRef", Value: s.dynamicRef},
		{Key: "$dynamicAnchor", Value: s.dynamicAnchor},
		{Key: "contentEncoding", Value: s.contentEncoding},
		{Key: "contentMediaType", Value: s.contentMediaType},
		{Key: "contentSchema", Value: s.contentSchema},
		{Key: "unevaluatedItems", Value: s.unevaluatedItems},
		{Key: "unevaluatedProperties", Value: s.unevaluatedProperties},
		{Key: "examples", Value: s.examples},
		{Key: "discriminator", Value: s.discriminator},
		{Key: "readOnly", Value: s.readOnly},
		{Key: "writeOnly", Value: s.writeOnly},
		{Key: "xml", Value: s.xml},
		{Key: "externalDocs", Value: s.externalDocs},
		{Key: "example", Value: s.example},
		{Key: "deprecated", Value: s.deprecated},
	}
	return shared.AppendExtensions(fields, s.VendorExtensions)
}

func (s *Schema) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(s.marshalFields())
}

func (s *Schema) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(s.marshalFields())
}

var _ yaml.Marshaler = (*Schema)(nil)
