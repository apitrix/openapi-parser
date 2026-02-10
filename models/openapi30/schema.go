package openapi30

// Schema represents the OpenAPI 3.0 Schema Object.
// https://spec.openapis.org/oas/v3.0.3#schema-object
type Schema struct {
	Node // embedded - provides VendorExtensions and Trix

	// JSON Schema fields
	Title            string                `json:"title,omitempty" yaml:"title,omitempty"`
	MultipleOf       *float64              `json:"multipleOf,omitempty" yaml:"multipleOf,omitempty"`
	Maximum          *float64              `json:"maximum,omitempty" yaml:"maximum,omitempty"`
	ExclusiveMaximum bool                  `json:"exclusiveMaximum,omitempty" yaml:"exclusiveMaximum,omitempty"`
	Minimum          *float64              `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	ExclusiveMinimum bool                  `json:"exclusiveMinimum,omitempty" yaml:"exclusiveMinimum,omitempty"`
	MaxLength        *uint64               `json:"maxLength,omitempty" yaml:"maxLength,omitempty"`
	MinLength        *uint64               `json:"minLength,omitempty" yaml:"minLength,omitempty"`
	Pattern          string                `json:"pattern,omitempty" yaml:"pattern,omitempty"`
	MaxItems         *uint64               `json:"maxItems,omitempty" yaml:"maxItems,omitempty"`
	MinItems         *uint64               `json:"minItems,omitempty" yaml:"minItems,omitempty"`
	UniqueItems      bool                  `json:"uniqueItems,omitempty" yaml:"uniqueItems,omitempty"`
	MaxProperties    *uint64               `json:"maxProperties,omitempty" yaml:"maxProperties,omitempty"`
	MinProperties    *uint64               `json:"minProperties,omitempty" yaml:"minProperties,omitempty"`
	Required         []string              `json:"required,omitempty" yaml:"required,omitempty"`
	Enum             []interface{}         `json:"enum,omitempty" yaml:"enum,omitempty"`
	Type             string                `json:"type,omitempty" yaml:"type,omitempty"`
	AllOf            []*SchemaRef          `json:"allOf,omitempty" yaml:"allOf,omitempty"`
	OneOf            []*SchemaRef          `json:"oneOf,omitempty" yaml:"oneOf,omitempty"`
	AnyOf            []*SchemaRef          `json:"anyOf,omitempty" yaml:"anyOf,omitempty"`
	Not              *SchemaRef            `json:"not,omitempty" yaml:"not,omitempty"`
	Items            *SchemaRef            `json:"items,omitempty" yaml:"items,omitempty"`
	Properties       map[string]*SchemaRef `json:"properties,omitempty" yaml:"properties,omitempty"`
	Description      string                `json:"description,omitempty" yaml:"description,omitempty"`
	Format           string                `json:"format,omitempty" yaml:"format,omitempty"`
	Default          interface{}           `json:"default,omitempty" yaml:"default,omitempty"`

	// AdditionalProperties can be a boolean or a schema.
	// When boolean: AdditionalPropertiesAllowed is set, AdditionalProperties is nil.
	// When schema: AdditionalProperties is set, AdditionalPropertiesAllowed is nil.
	AdditionalProperties        *SchemaRef `json:"additionalProperties,omitempty" yaml:"additionalProperties,omitempty"`
	AdditionalPropertiesAllowed *bool      `json:"-" yaml:"-"`

	// OpenAPI 3.0 specific fields
	Nullable      bool                   `json:"nullable,omitempty" yaml:"nullable,omitempty"`
	Discriminator *Discriminator         `json:"discriminator,omitempty" yaml:"discriminator,omitempty"`
	ReadOnly      bool                   `json:"readOnly,omitempty" yaml:"readOnly,omitempty"`
	WriteOnly     bool                   `json:"writeOnly,omitempty" yaml:"writeOnly,omitempty"`
	XML           *XML                   `json:"xml,omitempty" yaml:"xml,omitempty"`
	ExternalDocs  *ExternalDocumentation `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	Example       interface{}            `json:"example,omitempty" yaml:"example,omitempty"`
	Deprecated    bool                   `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
}

// NewSchema creates a new Schema instance.
func NewSchema() *Schema {
	return &Schema{}
}
