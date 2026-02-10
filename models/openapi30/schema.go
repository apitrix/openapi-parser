package openapi30

// Schema represents the OpenAPI 3.0 Schema Object.
// https://spec.openapis.org/oas/v3.0.3#schema-object
type Schema struct {
	Node // embedded - provides VendorExtensions and Trix

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
	allOf            []*SchemaRef
	oneOf            []*SchemaRef
	anyOf            []*SchemaRef
	not              *SchemaRef
	items            *SchemaRef
	properties       map[string]*SchemaRef
	description      string
	format           string
	defaultVal       interface{}

	// AdditionalProperties can be a boolean or a schema.
	additionalProperties        *SchemaRef
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

func (s *Schema) Title() string                        { return s.title }
func (s *Schema) MultipleOf() *float64                 { return s.multipleOf }
func (s *Schema) Maximum() *float64                    { return s.maximum }
func (s *Schema) ExclusiveMaximum() bool               { return s.exclusiveMaximum }
func (s *Schema) Minimum() *float64                    { return s.minimum }
func (s *Schema) ExclusiveMinimum() bool               { return s.exclusiveMinimum }
func (s *Schema) MaxLength() *uint64                   { return s.maxLength }
func (s *Schema) MinLength() *uint64                   { return s.minLength }
func (s *Schema) Pattern() string                      { return s.pattern }
func (s *Schema) MaxItems() *uint64                    { return s.maxItems }
func (s *Schema) MinItems() *uint64                    { return s.minItems }
func (s *Schema) UniqueItems() bool                    { return s.uniqueItems }
func (s *Schema) MaxProperties() *uint64               { return s.maxProperties }
func (s *Schema) MinProperties() *uint64               { return s.minProperties }
func (s *Schema) Required() []string                   { return s.required }
func (s *Schema) Enum() []interface{}                  { return s.enum }
func (s *Schema) Type() string                         { return s.schemaType }
func (s *Schema) AllOf() []*SchemaRef                  { return s.allOf }
func (s *Schema) OneOf() []*SchemaRef                  { return s.oneOf }
func (s *Schema) AnyOf() []*SchemaRef                  { return s.anyOf }
func (s *Schema) Not() *SchemaRef                      { return s.not }
func (s *Schema) Items() *SchemaRef                    { return s.items }
func (s *Schema) Properties() map[string]*SchemaRef    { return s.properties }
func (s *Schema) Description() string                  { return s.description }
func (s *Schema) Format() string                       { return s.format }
func (s *Schema) Default() interface{}                 { return s.defaultVal }
func (s *Schema) AdditionalProperties() *SchemaRef     { return s.additionalProperties }
func (s *Schema) AdditionalPropertiesAllowed() *bool   { return s.additionalPropertiesAllowed }
func (s *Schema) Nullable() bool                       { return s.nullable }
func (s *Schema) Discriminator() *Discriminator        { return s.discriminator }
func (s *Schema) ReadOnly() bool                       { return s.readOnly }
func (s *Schema) WriteOnly() bool                      { return s.writeOnly }
func (s *Schema) XML() *XML                            { return s.xml }
func (s *Schema) ExternalDocs() *ExternalDocumentation { return s.externalDocs }
func (s *Schema) Example() interface{}                 { return s.example }
func (s *Schema) Deprecated() bool                     { return s.deprecated }

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
	AllOf            []*SchemaRef
	OneOf            []*SchemaRef
	AnyOf            []*SchemaRef
	Not              *SchemaRef
	Items            *SchemaRef
	Properties       map[string]*SchemaRef
	Description      string
	Format           string
	Default          interface{}

	AdditionalProperties        *SchemaRef
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
