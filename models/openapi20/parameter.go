package openapi20

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// Parameter describes a single operation parameter.
// https://swagger.io/specification/v2/#parameter-object
type Parameter struct {
	Node // embedded - provides VendorExtensions and Trix

	name            string
	in              string
	description     string
	required        bool
	allowEmptyValue bool

	// For body parameters only
	schema *SchemaRef

	// For non-body parameters (query, header, path, formData)
	paramType        string
	format           string
	items            *Items
	collectionFormat string
	defaultVal       interface{}
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
	enum             []interface{}
	multipleOf       *float64
}

func (p *Parameter) Name() string             { return p.name }
func (p *Parameter) In() string               { return p.in }
func (p *Parameter) Description() string      { return p.description }
func (p *Parameter) Required() bool           { return p.required }
func (p *Parameter) AllowEmptyValue() bool    { return p.allowEmptyValue }
func (p *Parameter) Schema() *SchemaRef       { return p.schema }
func (p *Parameter) Type() string             { return p.paramType }
func (p *Parameter) Format() string           { return p.format }
func (p *Parameter) Items() *Items            { return p.items }
func (p *Parameter) CollectionFormat() string { return p.collectionFormat }
func (p *Parameter) Default() interface{}     { return p.defaultVal }
func (p *Parameter) Maximum() *float64        { return p.maximum }
func (p *Parameter) ExclusiveMaximum() bool   { return p.exclusiveMaximum }
func (p *Parameter) Minimum() *float64        { return p.minimum }
func (p *Parameter) ExclusiveMinimum() bool   { return p.exclusiveMinimum }
func (p *Parameter) MaxLength() *uint64       { return p.maxLength }
func (p *Parameter) MinLength() *uint64       { return p.minLength }
func (p *Parameter) Pattern() string          { return p.pattern }
func (p *Parameter) MaxItems() *uint64        { return p.maxItems }
func (p *Parameter) MinItems() *uint64        { return p.minItems }
func (p *Parameter) UniqueItems() bool        { return p.uniqueItems }
func (p *Parameter) Enum() []interface{}      { return p.enum }
func (p *Parameter) MultipleOf() *float64     { return p.multipleOf }

// ParameterFields groups all parameter properties for the constructor.
type ParameterFields struct {
	Name             string
	In               string
	Description      string
	Required         bool
	AllowEmptyValue  bool
	Schema           *SchemaRef
	Type             string
	Format           string
	Items            *Items
	CollectionFormat string
	Default          interface{}
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
	Enum             []interface{}
	MultipleOf       *float64
}

// NewParameter creates a new Parameter instance.
func NewParameter(f ParameterFields) *Parameter {
	return &Parameter{
		name: f.Name, in: f.In, description: f.Description,
		required: f.Required, allowEmptyValue: f.AllowEmptyValue,
		schema: f.Schema, paramType: f.Type, format: f.Format,
		items: f.Items, collectionFormat: f.CollectionFormat,
		defaultVal: f.Default, maximum: f.Maximum,
		exclusiveMaximum: f.ExclusiveMaximum, minimum: f.Minimum,
		exclusiveMinimum: f.ExclusiveMinimum, maxLength: f.MaxLength,
		minLength: f.MinLength, pattern: f.Pattern, maxItems: f.MaxItems,
		minItems: f.MinItems, uniqueItems: f.UniqueItems, enum: f.Enum,
		multipleOf: f.MultipleOf,
	}
}

func (p *Parameter) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "name", Value: p.name},
		{Key: "in", Value: p.in},
		{Key: "description", Value: p.description},
		{Key: "required", Value: p.required},
		{Key: "allowEmptyValue", Value: p.allowEmptyValue},
		{Key: "schema", Value: p.schema},
		{Key: "type", Value: p.paramType},
		{Key: "format", Value: p.format},
		{Key: "items", Value: p.items},
		{Key: "collectionFormat", Value: p.collectionFormat},
		{Key: "default", Value: p.defaultVal},
		{Key: "maximum", Value: p.maximum},
		{Key: "exclusiveMaximum", Value: p.exclusiveMaximum},
		{Key: "minimum", Value: p.minimum},
		{Key: "exclusiveMinimum", Value: p.exclusiveMinimum},
		{Key: "maxLength", Value: p.maxLength},
		{Key: "minLength", Value: p.minLength},
		{Key: "pattern", Value: p.pattern},
		{Key: "maxItems", Value: p.maxItems},
		{Key: "minItems", Value: p.minItems},
		{Key: "uniqueItems", Value: p.uniqueItems},
		{Key: "enum", Value: p.enum},
		{Key: "multipleOf", Value: p.multipleOf},
	}
	return shared.AppendExtensions(fields, p.VendorExtensions)
}

func (p *Parameter) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(p.marshalFields())
}

func (p *Parameter) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(p.marshalFields())
}

var _ yaml.Marshaler = (*Parameter)(nil)
