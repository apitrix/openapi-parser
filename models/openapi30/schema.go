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

func (s *Schema) SetTitle(title string) error {
	if err := s.Trix.RunHooks("title", s.title, title); err != nil {
		return err
	}
	s.title = title
	return nil
}
func (s *Schema) SetMultipleOf(multipleOf *float64) error {
	if err := s.Trix.RunHooks("multipleOf", s.multipleOf, multipleOf); err != nil {
		return err
	}
	s.multipleOf = multipleOf
	return nil
}
func (s *Schema) SetMaximum(maximum *float64) error {
	if err := s.Trix.RunHooks("maximum", s.maximum, maximum); err != nil {
		return err
	}
	s.maximum = maximum
	return nil
}
func (s *Schema) SetExclusiveMaximum(exclusiveMaximum bool) error {
	if err := s.Trix.RunHooks("exclusiveMaximum", s.exclusiveMaximum, exclusiveMaximum); err != nil {
		return err
	}
	s.exclusiveMaximum = exclusiveMaximum
	return nil
}
func (s *Schema) SetMinimum(minimum *float64) error {
	if err := s.Trix.RunHooks("minimum", s.minimum, minimum); err != nil {
		return err
	}
	s.minimum = minimum
	return nil
}
func (s *Schema) SetExclusiveMinimum(exclusiveMinimum bool) error {
	if err := s.Trix.RunHooks("exclusiveMinimum", s.exclusiveMinimum, exclusiveMinimum); err != nil {
		return err
	}
	s.exclusiveMinimum = exclusiveMinimum
	return nil
}
func (s *Schema) SetMaxLength(maxLength *uint64) error {
	if err := s.Trix.RunHooks("maxLength", s.maxLength, maxLength); err != nil {
		return err
	}
	s.maxLength = maxLength
	return nil
}
func (s *Schema) SetMinLength(minLength *uint64) error {
	if err := s.Trix.RunHooks("minLength", s.minLength, minLength); err != nil {
		return err
	}
	s.minLength = minLength
	return nil
}
func (s *Schema) SetPattern(pattern string) error {
	if err := s.Trix.RunHooks("pattern", s.pattern, pattern); err != nil {
		return err
	}
	s.pattern = pattern
	return nil
}
func (s *Schema) SetMaxItems(maxItems *uint64) error {
	if err := s.Trix.RunHooks("maxItems", s.maxItems, maxItems); err != nil {
		return err
	}
	s.maxItems = maxItems
	return nil
}
func (s *Schema) SetMinItems(minItems *uint64) error {
	if err := s.Trix.RunHooks("minItems", s.minItems, minItems); err != nil {
		return err
	}
	s.minItems = minItems
	return nil
}
func (s *Schema) SetUniqueItems(uniqueItems bool) error {
	if err := s.Trix.RunHooks("uniqueItems", s.uniqueItems, uniqueItems); err != nil {
		return err
	}
	s.uniqueItems = uniqueItems
	return nil
}
func (s *Schema) SetMaxProperties(maxProperties *uint64) error {
	if err := s.Trix.RunHooks("maxProperties", s.maxProperties, maxProperties); err != nil {
		return err
	}
	s.maxProperties = maxProperties
	return nil
}
func (s *Schema) SetMinProperties(minProperties *uint64) error {
	if err := s.Trix.RunHooks("minProperties", s.minProperties, minProperties); err != nil {
		return err
	}
	s.minProperties = minProperties
	return nil
}
func (s *Schema) SetRequired(required []string) error {
	if err := s.Trix.RunHooks("required", s.required, required); err != nil {
		return err
	}
	s.required = required
	return nil
}
func (s *Schema) SetEnum(enum []interface{}) error {
	if err := s.Trix.RunHooks("enum", s.enum, enum); err != nil {
		return err
	}
	s.enum = enum
	return nil
}
func (s *Schema) SetType(schemaType string) error {
	if err := s.Trix.RunHooks("type", s.schemaType, schemaType); err != nil {
		return err
	}
	s.schemaType = schemaType
	return nil
}
func (s *Schema) SetAllOf(allOf []*shared.Ref[Schema]) error {
	if err := s.Trix.RunHooks("allOf", s.allOf, allOf); err != nil {
		return err
	}
	s.allOf = allOf
	return nil
}
func (s *Schema) SetOneOf(oneOf []*shared.Ref[Schema]) error {
	if err := s.Trix.RunHooks("oneOf", s.oneOf, oneOf); err != nil {
		return err
	}
	s.oneOf = oneOf
	return nil
}
func (s *Schema) SetAnyOf(anyOf []*shared.Ref[Schema]) error {
	if err := s.Trix.RunHooks("anyOf", s.anyOf, anyOf); err != nil {
		return err
	}
	s.anyOf = anyOf
	return nil
}
func (s *Schema) SetNot(not *shared.Ref[Schema]) error {
	if err := s.Trix.RunHooks("not", s.not, not); err != nil {
		return err
	}
	s.not = not
	return nil
}
func (s *Schema) SetItems(items *shared.Ref[Schema]) error {
	if err := s.Trix.RunHooks("items", s.items, items); err != nil {
		return err
	}
	s.items = items
	return nil
}
func (s *Schema) SetProperties(properties map[string]*shared.Ref[Schema]) error {
	if err := s.Trix.RunHooks("properties", s.properties, properties); err != nil {
		return err
	}
	s.properties = properties
	return nil
}
func (s *Schema) SetDescription(description string) error {
	if err := s.Trix.RunHooks("description", s.description, description); err != nil {
		return err
	}
	s.description = description
	return nil
}
func (s *Schema) SetFormat(format string) error {
	if err := s.Trix.RunHooks("format", s.format, format); err != nil {
		return err
	}
	s.format = format
	return nil
}
func (s *Schema) SetDefault(defaultVal interface{}) error {
	if err := s.Trix.RunHooks("default", s.defaultVal, defaultVal); err != nil {
		return err
	}
	s.defaultVal = defaultVal
	return nil
}
func (s *Schema) SetAdditionalProperties(additionalProperties *shared.Ref[Schema]) error {
	if err := s.Trix.RunHooks("additionalProperties", s.additionalProperties, additionalProperties); err != nil {
		return err
	}
	s.additionalProperties = additionalProperties
	return nil
}
func (s *Schema) SetAdditionalPropertiesAllowed(additionalPropertiesAllowed *bool) error {
	if err := s.Trix.RunHooks("additionalProperties", s.additionalPropertiesAllowed, additionalPropertiesAllowed); err != nil {
		return err
	}
	s.additionalPropertiesAllowed = additionalPropertiesAllowed
	return nil
}
func (s *Schema) SetNullable(nullable bool) error {
	if err := s.Trix.RunHooks("nullable", s.nullable, nullable); err != nil {
		return err
	}
	s.nullable = nullable
	return nil
}
func (s *Schema) SetDiscriminator(discriminator *Discriminator) error {
	if err := s.Trix.RunHooks("discriminator", s.discriminator, discriminator); err != nil {
		return err
	}
	s.discriminator = discriminator
	return nil
}
func (s *Schema) SetReadOnly(readOnly bool) error {
	if err := s.Trix.RunHooks("readOnly", s.readOnly, readOnly); err != nil {
		return err
	}
	s.readOnly = readOnly
	return nil
}
func (s *Schema) SetWriteOnly(writeOnly bool) error {
	if err := s.Trix.RunHooks("writeOnly", s.writeOnly, writeOnly); err != nil {
		return err
	}
	s.writeOnly = writeOnly
	return nil
}
func (s *Schema) SetXML(xml *XML) error {
	if err := s.Trix.RunHooks("xml", s.xml, xml); err != nil {
		return err
	}
	s.xml = xml
	return nil
}
func (s *Schema) SetExternalDocs(externalDocs *ExternalDocumentation) error {
	if err := s.Trix.RunHooks("externalDocs", s.externalDocs, externalDocs); err != nil {
		return err
	}
	s.externalDocs = externalDocs
	return nil
}
func (s *Schema) SetExample(example interface{}) error {
	if err := s.Trix.RunHooks("example", s.example, example); err != nil {
		return err
	}
	s.example = example
	return nil
}
func (s *Schema) SetDeprecated(deprecated bool) error {
	if err := s.Trix.RunHooks("deprecated", s.deprecated, deprecated); err != nil {
		return err
	}
	s.deprecated = deprecated
	return nil
}

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
