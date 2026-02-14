package openapi20

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// Parameter describes a single operation parameter.
// https://swagger.io/specification/v2/#parameter-object
type Parameter struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	name            string
	in              string
	description     string
	required        bool
	allowEmptyValue bool

	// For body parameters only
	schema *shared.Ref[Schema]

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

func (p *Parameter) Name() string                 { return p.name }
func (p *Parameter) In() string                  { return p.in }
func (p *Parameter) Description() string       { return p.description }
func (p *Parameter) Required() bool             { return p.required }
func (p *Parameter) AllowEmptyValue() bool     { return p.allowEmptyValue }
func (p *Parameter) Schema() *shared.Ref[Schema] { return p.schema }
func (p *Parameter) Type() string               { return p.paramType }
func (p *Parameter) Format() string             { return p.format }
func (p *Parameter) Items() *Items              { return p.items }
func (p *Parameter) CollectionFormat() string  { return p.collectionFormat }
func (p *Parameter) Default() interface{}      { return p.defaultVal }
func (p *Parameter) Maximum() *float64          { return p.maximum }
func (p *Parameter) ExclusiveMaximum() bool    { return p.exclusiveMaximum }
func (p *Parameter) Minimum() *float64          { return p.minimum }
func (p *Parameter) ExclusiveMinimum() bool   { return p.exclusiveMinimum }
func (p *Parameter) MaxLength() *uint64         { return p.maxLength }
func (p *Parameter) MinLength() *uint64         { return p.minLength }
func (p *Parameter) Pattern() string            { return p.pattern }
func (p *Parameter) MaxItems() *uint64          { return p.maxItems }
func (p *Parameter) MinItems() *uint64          { return p.minItems }
func (p *Parameter) UniqueItems() bool          { return p.uniqueItems }
func (p *Parameter) Enum() []interface{}        { return p.enum }
func (p *Parameter) MultipleOf() *float64      { return p.multipleOf }

func (p *Parameter) SetName(name string) error {
	if err := p.Trix.RunHooks("name", p.name, name); err != nil {
		return err
	}
	p.name = name
	return nil
}
func (p *Parameter) SetIn(in string) error {
	if err := p.Trix.RunHooks("in", p.in, in); err != nil {
		return err
	}
	p.in = in
	return nil
}
func (p *Parameter) SetDescription(description string) error {
	if err := p.Trix.RunHooks("description", p.description, description); err != nil {
		return err
	}
	p.description = description
	return nil
}
func (p *Parameter) SetRequired(required bool) error {
	if err := p.Trix.RunHooks("required", p.required, required); err != nil {
		return err
	}
	p.required = required
	return nil
}
func (p *Parameter) SetAllowEmptyValue(allowEmptyValue bool) error {
	if err := p.Trix.RunHooks("allowEmptyValue", p.allowEmptyValue, allowEmptyValue); err != nil {
		return err
	}
	p.allowEmptyValue = allowEmptyValue
	return nil
}
func (p *Parameter) SetSchema(schema *shared.Ref[Schema]) error {
	if err := p.Trix.RunHooks("schema", p.schema, schema); err != nil {
		return err
	}
	p.schema = schema
	return nil
}
func (p *Parameter) SetType(paramType string) error {
	if err := p.Trix.RunHooks("type", p.paramType, paramType); err != nil {
		return err
	}
	p.paramType = paramType
	return nil
}
func (p *Parameter) SetFormat(format string) error {
	if err := p.Trix.RunHooks("format", p.format, format); err != nil {
		return err
	}
	p.format = format
	return nil
}
func (p *Parameter) SetItems(items *Items) error {
	if err := p.Trix.RunHooks("items", p.items, items); err != nil {
		return err
	}
	p.items = items
	return nil
}
func (p *Parameter) SetCollectionFormat(collectionFormat string) error {
	if err := p.Trix.RunHooks("collectionFormat", p.collectionFormat, collectionFormat); err != nil {
		return err
	}
	p.collectionFormat = collectionFormat
	return nil
}
func (p *Parameter) SetDefault(defaultVal interface{}) error {
	if err := p.Trix.RunHooks("default", p.defaultVal, defaultVal); err != nil {
		return err
	}
	p.defaultVal = defaultVal
	return nil
}
func (p *Parameter) SetMaximum(maximum *float64) error {
	if err := p.Trix.RunHooks("maximum", p.maximum, maximum); err != nil {
		return err
	}
	p.maximum = maximum
	return nil
}
func (p *Parameter) SetExclusiveMaximum(exclusiveMaximum bool) error {
	if err := p.Trix.RunHooks("exclusiveMaximum", p.exclusiveMaximum, exclusiveMaximum); err != nil {
		return err
	}
	p.exclusiveMaximum = exclusiveMaximum
	return nil
}
func (p *Parameter) SetMinimum(minimum *float64) error {
	if err := p.Trix.RunHooks("minimum", p.minimum, minimum); err != nil {
		return err
	}
	p.minimum = minimum
	return nil
}
func (p *Parameter) SetExclusiveMinimum(exclusiveMinimum bool) error {
	if err := p.Trix.RunHooks("exclusiveMinimum", p.exclusiveMinimum, exclusiveMinimum); err != nil {
		return err
	}
	p.exclusiveMinimum = exclusiveMinimum
	return nil
}
func (p *Parameter) SetMaxLength(maxLength *uint64) error {
	if err := p.Trix.RunHooks("maxLength", p.maxLength, maxLength); err != nil {
		return err
	}
	p.maxLength = maxLength
	return nil
}
func (p *Parameter) SetMinLength(minLength *uint64) error {
	if err := p.Trix.RunHooks("minLength", p.minLength, minLength); err != nil {
		return err
	}
	p.minLength = minLength
	return nil
}
func (p *Parameter) SetPattern(pattern string) error {
	if err := p.Trix.RunHooks("pattern", p.pattern, pattern); err != nil {
		return err
	}
	p.pattern = pattern
	return nil
}
func (p *Parameter) SetMaxItems(maxItems *uint64) error {
	if err := p.Trix.RunHooks("maxItems", p.maxItems, maxItems); err != nil {
		return err
	}
	p.maxItems = maxItems
	return nil
}
func (p *Parameter) SetMinItems(minItems *uint64) error {
	if err := p.Trix.RunHooks("minItems", p.minItems, minItems); err != nil {
		return err
	}
	p.minItems = minItems
	return nil
}
func (p *Parameter) SetUniqueItems(uniqueItems bool) error {
	if err := p.Trix.RunHooks("uniqueItems", p.uniqueItems, uniqueItems); err != nil {
		return err
	}
	p.uniqueItems = uniqueItems
	return nil
}
func (p *Parameter) SetEnum(enum []interface{}) error {
	if err := p.Trix.RunHooks("enum", p.enum, enum); err != nil {
		return err
	}
	p.enum = enum
	return nil
}
func (p *Parameter) SetMultipleOf(multipleOf *float64) error {
	if err := p.Trix.RunHooks("multipleOf", p.multipleOf, multipleOf); err != nil {
		return err
	}
	p.multipleOf = multipleOf
	return nil
}

// ParameterFields groups all parameter properties for the constructor.
type ParameterFields struct {
	Name             string
	In               string
	Description      string
	Required         bool
	AllowEmptyValue  bool
	Schema           *shared.Ref[Schema]
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
