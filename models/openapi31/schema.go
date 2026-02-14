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
	allOf            []*RefSchema
	oneOf            []*RefSchema
	anyOf            []*RefSchema
	not              *RefSchema
	items            *RefSchema
	properties       map[string]*RefSchema
	description      string
	format           string
	defaultVal       interface{}

	// AdditionalProperties can be a boolean or a schema.
	additionalProperties        *RefSchema
	additionalPropertiesAllowed *bool

	// JSON Schema 2020-12 new keywords
	constVal              interface{}
	ifSchema              *RefSchema
	thenSchema            *RefSchema
	elseSchema            *RefSchema
	dependentSchemas      map[string]*RefSchema
	prefixItems           []*RefSchema
	anchor                string
	dynamicRef            string
	dynamicAnchor         string
	contentEncoding       string
	contentMediaType      string
	contentSchema         *RefSchema
	unevaluatedItems      *RefSchema
	unevaluatedProperties *RefSchema
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

func (s *Schema) Title() string                           { return s.title }
func (s *Schema) MultipleOf() *float64                    { return s.multipleOf }
func (s *Schema) Maximum() *float64                       { return s.maximum }
func (s *Schema) ExclusiveMaximum() *float64              { return s.exclusiveMaximum }
func (s *Schema) Minimum() *float64                       { return s.minimum }
func (s *Schema) ExclusiveMinimum() *float64              { return s.exclusiveMinimum }
func (s *Schema) MaxLength() *uint64                      { return s.maxLength }
func (s *Schema) MinLength() *uint64                      { return s.minLength }
func (s *Schema) Pattern() string                         { return s.pattern }
func (s *Schema) MaxItems() *uint64                       { return s.maxItems }
func (s *Schema) MinItems() *uint64                       { return s.minItems }
func (s *Schema) UniqueItems() bool                       { return s.uniqueItems }
func (s *Schema) MaxProperties() *uint64                  { return s.maxProperties }
func (s *Schema) MinProperties() *uint64                  { return s.minProperties }
func (s *Schema) Required() []string                      { return s.required }
func (s *Schema) Enum() []interface{}                     { return s.enum }
func (s *Schema) Type() SchemaType                        { return s.schemaType }
func (s *Schema) AllOf() []*RefSchema                     { return s.allOf }
func (s *Schema) OneOf() []*RefSchema                     { return s.oneOf }
func (s *Schema) AnyOf() []*RefSchema                     { return s.anyOf }
func (s *Schema) Not() *RefSchema                         { return s.not }
func (s *Schema) Items() *RefSchema                       { return s.items }
func (s *Schema) Properties() map[string]*RefSchema       { return s.properties }
func (s *Schema) Description() string                     { return s.description }
func (s *Schema) Format() string                          { return s.format }
func (s *Schema) Default() interface{}                    { return s.defaultVal }
func (s *Schema) AdditionalProperties() *RefSchema        { return s.additionalProperties }
func (s *Schema) AdditionalPropertiesAllowed() *bool      { return s.additionalPropertiesAllowed }
func (s *Schema) Const() interface{}                      { return s.constVal }
func (s *Schema) If() *RefSchema                          { return s.ifSchema }
func (s *Schema) Then() *RefSchema                        { return s.thenSchema }
func (s *Schema) Else() *RefSchema                        { return s.elseSchema }
func (s *Schema) DependentSchemas() map[string]*RefSchema { return s.dependentSchemas }
func (s *Schema) PrefixItems() []*RefSchema               { return s.prefixItems }
func (s *Schema) Anchor() string                          { return s.anchor }
func (s *Schema) DynamicRef() string                      { return s.dynamicRef }
func (s *Schema) DynamicAnchor() string                   { return s.dynamicAnchor }
func (s *Schema) ContentEncoding() string                 { return s.contentEncoding }
func (s *Schema) ContentMediaType() string                { return s.contentMediaType }
func (s *Schema) ContentSchema() *RefSchema               { return s.contentSchema }
func (s *Schema) UnevaluatedItems() *RefSchema            { return s.unevaluatedItems }
func (s *Schema) UnevaluatedProperties() *RefSchema       { return s.unevaluatedProperties }
func (s *Schema) Examples() []interface{}                 { return s.examples }
func (s *Schema) Discriminator() *Discriminator           { return s.discriminator }
func (s *Schema) ReadOnly() bool                          { return s.readOnly }
func (s *Schema) WriteOnly() bool                         { return s.writeOnly }
func (s *Schema) XML() *XML                               { return s.xml }
func (s *Schema) ExternalDocs() *ExternalDocumentation    { return s.externalDocs }
func (s *Schema) Example() interface{}                    { return s.example }
func (s *Schema) Deprecated() bool                        { return s.deprecated }

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
func (s *Schema) SetExclusiveMaximum(exclusiveMaximum *float64) error {
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
func (s *Schema) SetExclusiveMinimum(exclusiveMinimum *float64) error {
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
func (s *Schema) SetType(schemaType SchemaType) error {
	if err := s.Trix.RunHooks("type", s.schemaType, schemaType); err != nil {
		return err
	}
	s.schemaType = schemaType
	return nil
}
func (s *Schema) SetAllOf(allOf []*RefSchema) error {
	if err := s.Trix.RunHooks("allOf", s.allOf, allOf); err != nil {
		return err
	}
	s.allOf = allOf
	return nil
}
func (s *Schema) SetOneOf(oneOf []*RefSchema) error {
	if err := s.Trix.RunHooks("oneOf", s.oneOf, oneOf); err != nil {
		return err
	}
	s.oneOf = oneOf
	return nil
}
func (s *Schema) SetAnyOf(anyOf []*RefSchema) error {
	if err := s.Trix.RunHooks("anyOf", s.anyOf, anyOf); err != nil {
		return err
	}
	s.anyOf = anyOf
	return nil
}
func (s *Schema) SetNot(not *RefSchema) error {
	if err := s.Trix.RunHooks("not", s.not, not); err != nil {
		return err
	}
	s.not = not
	return nil
}
func (s *Schema) SetItems(items *RefSchema) error {
	if err := s.Trix.RunHooks("items", s.items, items); err != nil {
		return err
	}
	s.items = items
	return nil
}
func (s *Schema) SetProperties(properties map[string]*RefSchema) error {
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
func (s *Schema) SetAdditionalProperties(additionalProperties *RefSchema) error {
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
func (s *Schema) SetConst(constVal interface{}) error {
	if err := s.Trix.RunHooks("const", s.constVal, constVal); err != nil {
		return err
	}
	s.constVal = constVal
	return nil
}
func (s *Schema) SetIf(ifSchema *RefSchema) error {
	if err := s.Trix.RunHooks("if", s.ifSchema, ifSchema); err != nil {
		return err
	}
	s.ifSchema = ifSchema
	return nil
}
func (s *Schema) SetThen(thenSchema *RefSchema) error {
	if err := s.Trix.RunHooks("then", s.thenSchema, thenSchema); err != nil {
		return err
	}
	s.thenSchema = thenSchema
	return nil
}
func (s *Schema) SetElse(elseSchema *RefSchema) error {
	if err := s.Trix.RunHooks("else", s.elseSchema, elseSchema); err != nil {
		return err
	}
	s.elseSchema = elseSchema
	return nil
}
func (s *Schema) SetDependentSchemas(dependentSchemas map[string]*RefSchema) error {
	if err := s.Trix.RunHooks("dependentSchemas", s.dependentSchemas, dependentSchemas); err != nil {
		return err
	}
	s.dependentSchemas = dependentSchemas
	return nil
}
func (s *Schema) SetPrefixItems(prefixItems []*RefSchema) error {
	if err := s.Trix.RunHooks("prefixItems", s.prefixItems, prefixItems); err != nil {
		return err
	}
	s.prefixItems = prefixItems
	return nil
}
func (s *Schema) SetAnchor(anchor string) error {
	if err := s.Trix.RunHooks("$anchor", s.anchor, anchor); err != nil {
		return err
	}
	s.anchor = anchor
	return nil
}
func (s *Schema) SetDynamicRef(dynamicRef string) error {
	if err := s.Trix.RunHooks("$dynamicRef", s.dynamicRef, dynamicRef); err != nil {
		return err
	}
	s.dynamicRef = dynamicRef
	return nil
}
func (s *Schema) SetDynamicAnchor(dynamicAnchor string) error {
	if err := s.Trix.RunHooks("$dynamicAnchor", s.dynamicAnchor, dynamicAnchor); err != nil {
		return err
	}
	s.dynamicAnchor = dynamicAnchor
	return nil
}
func (s *Schema) SetContentEncoding(contentEncoding string) error {
	if err := s.Trix.RunHooks("contentEncoding", s.contentEncoding, contentEncoding); err != nil {
		return err
	}
	s.contentEncoding = contentEncoding
	return nil
}
func (s *Schema) SetContentMediaType(contentMediaType string) error {
	if err := s.Trix.RunHooks("contentMediaType", s.contentMediaType, contentMediaType); err != nil {
		return err
	}
	s.contentMediaType = contentMediaType
	return nil
}
func (s *Schema) SetContentSchema(contentSchema *RefSchema) error {
	if err := s.Trix.RunHooks("contentSchema", s.contentSchema, contentSchema); err != nil {
		return err
	}
	s.contentSchema = contentSchema
	return nil
}
func (s *Schema) SetUnevaluatedItems(unevaluatedItems *RefSchema) error {
	if err := s.Trix.RunHooks("unevaluatedItems", s.unevaluatedItems, unevaluatedItems); err != nil {
		return err
	}
	s.unevaluatedItems = unevaluatedItems
	return nil
}
func (s *Schema) SetUnevaluatedProperties(unevaluatedProperties *RefSchema) error {
	if err := s.Trix.RunHooks("unevaluatedProperties", s.unevaluatedProperties, unevaluatedProperties); err != nil {
		return err
	}
	s.unevaluatedProperties = unevaluatedProperties
	return nil
}
func (s *Schema) SetExamples(examples []interface{}) error {
	if err := s.Trix.RunHooks("examples", s.examples, examples); err != nil {
		return err
	}
	s.examples = examples
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
	AllOf            []*RefSchema
	OneOf            []*RefSchema
	AnyOf            []*RefSchema
	Not              *RefSchema
	Items            *RefSchema
	Properties       map[string]*RefSchema
	Description      string
	Format           string
	Default          interface{}

	AdditionalProperties        *RefSchema
	AdditionalPropertiesAllowed *bool

	// JSON Schema 2020-12
	Const                 interface{}
	If                    *RefSchema
	Then                  *RefSchema
	Else                  *RefSchema
	DependentSchemas      map[string]*RefSchema
	PrefixItems           []*RefSchema
	Anchor                string
	DynamicRef            string
	DynamicAnchor         string
	ContentEncoding       string
	ContentMediaType      string
	ContentSchema         *RefSchema
	UnevaluatedItems      *RefSchema
	UnevaluatedProperties *RefSchema
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

// MarshalFields implements shared.MarshalFieldsProvider for export.
func (s *Schema) MarshalFields() []shared.Field { return s.marshalFields() }

func (s *Schema) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(s.marshalFields())
}

func (s *Schema) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(s.marshalFields())
}

var _ yaml.Marshaler = (*Schema)(nil)
