package openapi31

// SchemaType represents a JSON Schema type field that can be either a single
// string or an array of strings (JSON Schema Draft 2020-12).
type SchemaType struct {
	// Single is set when the type is a single string value (e.g. "string").
	Single string
	// Array is set when the type is an array of strings (e.g. ["string", "null"]).
	Array []string
}

// IsEmpty returns true if no type was specified.
func (t SchemaType) IsEmpty() bool {
	return t.Single == "" && len(t.Array) == 0
}

// Values returns all type values as a slice, whether specified as single or array.
func (t SchemaType) Values() []string {
	if len(t.Array) > 0 {
		return t.Array
	}
	if t.Single != "" {
		return []string{t.Single}
	}
	return nil
}

// Schema represents the OpenAPI 3.1 Schema Object.
// Based on JSON Schema Draft 2020-12.
// https://spec.openapis.org/oas/v3.1.0#schema-object
type Schema struct {
	Node // embedded - provides NodeSource and Extensions

	// JSON Schema fields
	Title            string                `json:"title,omitempty" yaml:"title,omitempty"`
	MultipleOf       *float64              `json:"multipleOf,omitempty" yaml:"multipleOf,omitempty"`
	Maximum          *float64              `json:"maximum,omitempty" yaml:"maximum,omitempty"`
	ExclusiveMaximum *float64              `json:"exclusiveMaximum,omitempty" yaml:"exclusiveMaximum,omitempty"`
	Minimum          *float64              `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	ExclusiveMinimum *float64              `json:"exclusiveMinimum,omitempty" yaml:"exclusiveMinimum,omitempty"`
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
	Type             SchemaType            `json:"type,omitempty" yaml:"type,omitempty"`
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

	// JSON Schema 2020-12 new keywords
	Const                 interface{}           `json:"const,omitempty" yaml:"const,omitempty"`
	If                    *SchemaRef            `json:"if,omitempty" yaml:"if,omitempty"`
	Then                  *SchemaRef            `json:"then,omitempty" yaml:"then,omitempty"`
	Else                  *SchemaRef            `json:"else,omitempty" yaml:"else,omitempty"`
	DependentSchemas      map[string]*SchemaRef `json:"dependentSchemas,omitempty" yaml:"dependentSchemas,omitempty"`
	PrefixItems           []*SchemaRef          `json:"prefixItems,omitempty" yaml:"prefixItems,omitempty"`
	Anchor                string                `json:"$anchor,omitempty" yaml:"$anchor,omitempty"`
	DynamicRef            string                `json:"$dynamicRef,omitempty" yaml:"$dynamicRef,omitempty"`
	DynamicAnchor         string                `json:"$dynamicAnchor,omitempty" yaml:"$dynamicAnchor,omitempty"`
	ContentEncoding       string                `json:"contentEncoding,omitempty" yaml:"contentEncoding,omitempty"`
	ContentMediaType      string                `json:"contentMediaType,omitempty" yaml:"contentMediaType,omitempty"`
	ContentSchema         *SchemaRef            `json:"contentSchema,omitempty" yaml:"contentSchema,omitempty"`
	UnevaluatedItems      *SchemaRef            `json:"unevaluatedItems,omitempty" yaml:"unevaluatedItems,omitempty"`
	UnevaluatedProperties *SchemaRef            `json:"unevaluatedProperties,omitempty" yaml:"unevaluatedProperties,omitempty"`
	Examples              []interface{}         `json:"examples,omitempty" yaml:"examples,omitempty"`

	// OpenAPI extensions (still present in 3.1)
	Discriminator *Discriminator         `json:"discriminator,omitempty" yaml:"discriminator,omitempty"`
	ReadOnly      bool                   `json:"readOnly,omitempty" yaml:"readOnly,omitempty"`
	WriteOnly     bool                   `json:"writeOnly,omitempty" yaml:"writeOnly,omitempty"`
	XML           *XML                   `json:"xml,omitempty" yaml:"xml,omitempty"`
	ExternalDocs  *ExternalDocumentation `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	Example       interface{}            `json:"example,omitempty" yaml:"example,omitempty"`
	Deprecated    bool                   `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
}

// Discriminator is used for polymorphism support.
// https://spec.openapis.org/oas/v3.1.0#discriminator-object
type Discriminator struct {
	Node // embedded - provides NodeSource and Extensions

	PropertyName string            `json:"propertyName" yaml:"propertyName"`
	Mapping      map[string]string `json:"mapping,omitempty" yaml:"mapping,omitempty"`
}

// XML provides additional metadata for XML serialization.
// https://spec.openapis.org/oas/v3.1.0#xml-object
type XML struct {
	Node // embedded - provides NodeSource and Extensions

	Name      string `json:"name,omitempty" yaml:"name,omitempty"`
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	Prefix    string `json:"prefix,omitempty" yaml:"prefix,omitempty"`
	Attribute bool   `json:"attribute,omitempty" yaml:"attribute,omitempty"`
	Wrapped   bool   `json:"wrapped,omitempty" yaml:"wrapped,omitempty"`
}
