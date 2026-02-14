package openapi20

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// Items describes the type of items in an array parameter.
// https://swagger.io/specification/v2/#items-object
type Items struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	itemType         string
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

func (it *Items) Type() string             { return it.itemType }
func (it *Items) Format() string           { return it.format }
func (it *Items) Items() *Items            { return it.items }
func (it *Items) CollectionFormat() string { return it.collectionFormat }
func (it *Items) Default() interface{}     { return it.defaultVal }
func (it *Items) Maximum() *float64        { return it.maximum }
func (it *Items) ExclusiveMaximum() bool   { return it.exclusiveMaximum }
func (it *Items) Minimum() *float64        { return it.minimum }
func (it *Items) ExclusiveMinimum() bool   { return it.exclusiveMinimum }
func (it *Items) MaxLength() *uint64       { return it.maxLength }
func (it *Items) MinLength() *uint64       { return it.minLength }
func (it *Items) Pattern() string          { return it.pattern }
func (it *Items) MaxItems() *uint64        { return it.maxItems }
func (it *Items) MinItems() *uint64        { return it.minItems }
func (it *Items) UniqueItems() bool        { return it.uniqueItems }
func (it *Items) Enum() []interface{}      { return it.enum }
func (it *Items) MultipleOf() *float64     { return it.multipleOf }

func (it *Items) SetType(itemType string) error {
	if err := it.Trix.RunHooks("type", it.itemType, itemType); err != nil {
		return err
	}
	it.itemType = itemType
	return nil
}
func (it *Items) SetFormat(format string) error {
	if err := it.Trix.RunHooks("format", it.format, format); err != nil {
		return err
	}
	it.format = format
	return nil
}
func (it *Items) SetItems(items *Items) error {
	if err := it.Trix.RunHooks("items", it.items, items); err != nil {
		return err
	}
	it.items = items
	return nil
}
func (it *Items) SetCollectionFormat(collectionFormat string) error {
	if err := it.Trix.RunHooks("collectionFormat", it.collectionFormat, collectionFormat); err != nil {
		return err
	}
	it.collectionFormat = collectionFormat
	return nil
}
func (it *Items) SetDefault(defaultVal interface{}) error {
	if err := it.Trix.RunHooks("default", it.defaultVal, defaultVal); err != nil {
		return err
	}
	it.defaultVal = defaultVal
	return nil
}
func (it *Items) SetMaximum(maximum *float64) error {
	if err := it.Trix.RunHooks("maximum", it.maximum, maximum); err != nil {
		return err
	}
	it.maximum = maximum
	return nil
}
func (it *Items) SetExclusiveMaximum(exclusiveMaximum bool) error {
	if err := it.Trix.RunHooks("exclusiveMaximum", it.exclusiveMaximum, exclusiveMaximum); err != nil {
		return err
	}
	it.exclusiveMaximum = exclusiveMaximum
	return nil
}
func (it *Items) SetMinimum(minimum *float64) error {
	if err := it.Trix.RunHooks("minimum", it.minimum, minimum); err != nil {
		return err
	}
	it.minimum = minimum
	return nil
}
func (it *Items) SetExclusiveMinimum(exclusiveMinimum bool) error {
	if err := it.Trix.RunHooks("exclusiveMinimum", it.exclusiveMinimum, exclusiveMinimum); err != nil {
		return err
	}
	it.exclusiveMinimum = exclusiveMinimum
	return nil
}
func (it *Items) SetMaxLength(maxLength *uint64) error {
	if err := it.Trix.RunHooks("maxLength", it.maxLength, maxLength); err != nil {
		return err
	}
	it.maxLength = maxLength
	return nil
}
func (it *Items) SetMinLength(minLength *uint64) error {
	if err := it.Trix.RunHooks("minLength", it.minLength, minLength); err != nil {
		return err
	}
	it.minLength = minLength
	return nil
}
func (it *Items) SetPattern(pattern string) error {
	if err := it.Trix.RunHooks("pattern", it.pattern, pattern); err != nil {
		return err
	}
	it.pattern = pattern
	return nil
}
func (it *Items) SetMaxItems(maxItems *uint64) error {
	if err := it.Trix.RunHooks("maxItems", it.maxItems, maxItems); err != nil {
		return err
	}
	it.maxItems = maxItems
	return nil
}
func (it *Items) SetMinItems(minItems *uint64) error {
	if err := it.Trix.RunHooks("minItems", it.minItems, minItems); err != nil {
		return err
	}
	it.minItems = minItems
	return nil
}
func (it *Items) SetUniqueItems(uniqueItems bool) error {
	if err := it.Trix.RunHooks("uniqueItems", it.uniqueItems, uniqueItems); err != nil {
		return err
	}
	it.uniqueItems = uniqueItems
	return nil
}
func (it *Items) SetEnum(enum []interface{}) error {
	if err := it.Trix.RunHooks("enum", it.enum, enum); err != nil {
		return err
	}
	it.enum = enum
	return nil
}
func (it *Items) SetMultipleOf(multipleOf *float64) error {
	if err := it.Trix.RunHooks("multipleOf", it.multipleOf, multipleOf); err != nil {
		return err
	}
	it.multipleOf = multipleOf
	return nil
}

// ItemsFields groups all items properties for the constructor.
type ItemsFields struct {
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

// NewItems creates a new Items instance.
func NewItems(f ItemsFields) *Items {
	return &Items{
		itemType: f.Type, format: f.Format, items: f.Items,
		collectionFormat: f.CollectionFormat, defaultVal: f.Default,
		maximum: f.Maximum, exclusiveMaximum: f.ExclusiveMaximum,
		minimum: f.Minimum, exclusiveMinimum: f.ExclusiveMinimum,
		maxLength: f.MaxLength, minLength: f.MinLength, pattern: f.Pattern,
		maxItems: f.MaxItems, minItems: f.MinItems, uniqueItems: f.UniqueItems,
		enum: f.Enum, multipleOf: f.MultipleOf,
	}
}

func (it *Items) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "type", Value: it.itemType},
		{Key: "format", Value: it.format},
		{Key: "items", Value: it.items},
		{Key: "collectionFormat", Value: it.collectionFormat},
		{Key: "default", Value: it.defaultVal},
		{Key: "maximum", Value: it.maximum},
		{Key: "exclusiveMaximum", Value: it.exclusiveMaximum},
		{Key: "minimum", Value: it.minimum},
		{Key: "exclusiveMinimum", Value: it.exclusiveMinimum},
		{Key: "maxLength", Value: it.maxLength},
		{Key: "minLength", Value: it.minLength},
		{Key: "pattern", Value: it.pattern},
		{Key: "maxItems", Value: it.maxItems},
		{Key: "minItems", Value: it.minItems},
		{Key: "uniqueItems", Value: it.uniqueItems},
		{Key: "enum", Value: it.enum},
		{Key: "multipleOf", Value: it.multipleOf},
	}
	return shared.AppendExtensions(fields, it.VendorExtensions)
}

func (it *Items) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(it.marshalFields())
}

func (it *Items) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(it.marshalFields())
}

var _ yaml.Marshaler = (*Items)(nil)
