package openapi20

// Schema represents the Swagger 2.0 Schema Object (JSON Schema subset).
// https://swagger.io/specification/v2/#schema-object
type Schema struct {
	Node // embedded - provides VendorExtensions and Trix

	// JSON Schema fields
	title                       string
	description                 string
	defaultVal                  interface{}
	multipleOf                  *float64
	maximum                     *float64
	exclusiveMaximum            bool
	minimum                     *float64
	exclusiveMinimum            bool
	maxLength                   *uint64
	minLength                   *uint64
	pattern                     string
	maxItems                    *uint64
	minItems                    *uint64
	uniqueItems                 bool
	maxProperties               *uint64
	minProperties               *uint64
	required                    []string
	enum                        []interface{}
	schemaType                  string
	format                      string
	allOf                       []*SchemaRef
	items                       *SchemaRef
	properties                  map[string]*SchemaRef
	additionalProperties        *SchemaRef
	additionalPropertiesAllowed *bool

	// Swagger 2.0 specific fields
	discriminator string
	readOnly      bool
	xml           *XML
	externalDocs  *ExternalDocs
	example       interface{}
}

func (s *Schema) Title() string                      { return s.title }
func (s *Schema) Description() string                { return s.description }
func (s *Schema) Default() interface{}               { return s.defaultVal }
func (s *Schema) MultipleOf() *float64               { return s.multipleOf }
func (s *Schema) Maximum() *float64                  { return s.maximum }
func (s *Schema) ExclusiveMaximum() bool             { return s.exclusiveMaximum }
func (s *Schema) Minimum() *float64                  { return s.minimum }
func (s *Schema) ExclusiveMinimum() bool             { return s.exclusiveMinimum }
func (s *Schema) MaxLength() *uint64                 { return s.maxLength }
func (s *Schema) MinLength() *uint64                 { return s.minLength }
func (s *Schema) Pattern() string                    { return s.pattern }
func (s *Schema) MaxItems() *uint64                  { return s.maxItems }
func (s *Schema) MinItems() *uint64                  { return s.minItems }
func (s *Schema) UniqueItems() bool                  { return s.uniqueItems }
func (s *Schema) MaxProperties() *uint64             { return s.maxProperties }
func (s *Schema) MinProperties() *uint64             { return s.minProperties }
func (s *Schema) Required() []string                 { return s.required }
func (s *Schema) Enum() []interface{}                { return s.enum }
func (s *Schema) Type() string                       { return s.schemaType }
func (s *Schema) Format() string                     { return s.format }
func (s *Schema) AllOf() []*SchemaRef                { return s.allOf }
func (s *Schema) Items() *SchemaRef                  { return s.items }
func (s *Schema) Properties() map[string]*SchemaRef  { return s.properties }
func (s *Schema) AdditionalProperties() *SchemaRef   { return s.additionalProperties }
func (s *Schema) AdditionalPropertiesAllowed() *bool { return s.additionalPropertiesAllowed }
func (s *Schema) Discriminator() string              { return s.discriminator }
func (s *Schema) ReadOnly() bool                     { return s.readOnly }
func (s *Schema) XML() *XML                          { return s.xml }
func (s *Schema) ExternalDocs() *ExternalDocs        { return s.externalDocs }
func (s *Schema) Example() interface{}               { return s.example }

// SchemaFields groups all schema properties for the constructor.
type SchemaFields struct {
	Title                       string
	Description                 string
	Default                     interface{}
	MultipleOf                  *float64
	Maximum                     *float64
	ExclusiveMaximum            bool
	Minimum                     *float64
	ExclusiveMinimum            bool
	MaxLength                   *uint64
	MinLength                   *uint64
	Pattern                     string
	MaxItems                    *uint64
	MinItems                    *uint64
	UniqueItems                 bool
	MaxProperties               *uint64
	MinProperties               *uint64
	Required                    []string
	Enum                        []interface{}
	Type                        string
	Format                      string
	AllOf                       []*SchemaRef
	Items                       *SchemaRef
	Properties                  map[string]*SchemaRef
	AdditionalProperties        *SchemaRef
	AdditionalPropertiesAllowed *bool
	Discriminator               string
	ReadOnly                    bool
	XML                         *XML
	ExternalDocs                *ExternalDocs
	Example                     interface{}
}

// NewSchema creates a new Schema instance.
func NewSchema(f SchemaFields) *Schema {
	return &Schema{
		title: f.Title, description: f.Description, defaultVal: f.Default,
		multipleOf: f.MultipleOf, maximum: f.Maximum,
		exclusiveMaximum: f.ExclusiveMaximum, minimum: f.Minimum,
		exclusiveMinimum: f.ExclusiveMinimum, maxLength: f.MaxLength,
		minLength: f.MinLength, pattern: f.Pattern, maxItems: f.MaxItems,
		minItems: f.MinItems, uniqueItems: f.UniqueItems,
		maxProperties: f.MaxProperties, minProperties: f.MinProperties,
		required: f.Required, enum: f.Enum, schemaType: f.Type,
		format: f.Format, allOf: f.AllOf, items: f.Items,
		properties: f.Properties, additionalProperties: f.AdditionalProperties,
		additionalPropertiesAllowed: f.AdditionalPropertiesAllowed,
		discriminator:               f.Discriminator, readOnly: f.ReadOnly,
		xml: f.XML, externalDocs: f.ExternalDocs, example: f.Example,
	}
}
