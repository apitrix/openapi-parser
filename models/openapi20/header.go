package openapi20

import (
	"openapi-parser/models/shared"

	"gopkg.in/yaml.v3"
)

// Header represents a Header Object in a response.
// https://swagger.io/specification/v2/#header-object
type Header struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	description      string
	headerType       string
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

func (h *Header) Description() string      { return h.description }
func (h *Header) Type() string             { return h.headerType }
func (h *Header) Format() string           { return h.format }
func (h *Header) Items() *Items            { return h.items }
func (h *Header) CollectionFormat() string { return h.collectionFormat }
func (h *Header) Default() interface{}     { return h.defaultVal }
func (h *Header) Maximum() *float64        { return h.maximum }
func (h *Header) ExclusiveMaximum() bool   { return h.exclusiveMaximum }
func (h *Header) Minimum() *float64        { return h.minimum }
func (h *Header) ExclusiveMinimum() bool   { return h.exclusiveMinimum }
func (h *Header) MaxLength() *uint64       { return h.maxLength }
func (h *Header) MinLength() *uint64       { return h.minLength }
func (h *Header) Pattern() string          { return h.pattern }
func (h *Header) MaxItems() *uint64        { return h.maxItems }
func (h *Header) MinItems() *uint64        { return h.minItems }
func (h *Header) UniqueItems() bool        { return h.uniqueItems }
func (h *Header) Enum() []interface{}      { return h.enum }
func (h *Header) MultipleOf() *float64     { return h.multipleOf }

func (h *Header) SetDescription(description string) error {
	if err := h.Trix.RunHooks("description", h.description, description); err != nil {
		return err
	}
	h.description = description
	return nil
}
func (h *Header) SetType(headerType string) error {
	if err := h.Trix.RunHooks("type", h.headerType, headerType); err != nil {
		return err
	}
	h.headerType = headerType
	return nil
}
func (h *Header) SetFormat(format string) error {
	if err := h.Trix.RunHooks("format", h.format, format); err != nil {
		return err
	}
	h.format = format
	return nil
}
func (h *Header) SetItems(items *Items) error {
	if err := h.Trix.RunHooks("items", h.items, items); err != nil {
		return err
	}
	h.items = items
	return nil
}
func (h *Header) SetCollectionFormat(collectionFormat string) error {
	if err := h.Trix.RunHooks("collectionFormat", h.collectionFormat, collectionFormat); err != nil {
		return err
	}
	h.collectionFormat = collectionFormat
	return nil
}
func (h *Header) SetDefault(defaultVal interface{}) error {
	if err := h.Trix.RunHooks("default", h.defaultVal, defaultVal); err != nil {
		return err
	}
	h.defaultVal = defaultVal
	return nil
}
func (h *Header) SetMaximum(maximum *float64) error {
	if err := h.Trix.RunHooks("maximum", h.maximum, maximum); err != nil {
		return err
	}
	h.maximum = maximum
	return nil
}
func (h *Header) SetExclusiveMaximum(exclusiveMaximum bool) error {
	if err := h.Trix.RunHooks("exclusiveMaximum", h.exclusiveMaximum, exclusiveMaximum); err != nil {
		return err
	}
	h.exclusiveMaximum = exclusiveMaximum
	return nil
}
func (h *Header) SetMinimum(minimum *float64) error {
	if err := h.Trix.RunHooks("minimum", h.minimum, minimum); err != nil {
		return err
	}
	h.minimum = minimum
	return nil
}
func (h *Header) SetExclusiveMinimum(exclusiveMinimum bool) error {
	if err := h.Trix.RunHooks("exclusiveMinimum", h.exclusiveMinimum, exclusiveMinimum); err != nil {
		return err
	}
	h.exclusiveMinimum = exclusiveMinimum
	return nil
}
func (h *Header) SetMaxLength(maxLength *uint64) error {
	if err := h.Trix.RunHooks("maxLength", h.maxLength, maxLength); err != nil {
		return err
	}
	h.maxLength = maxLength
	return nil
}
func (h *Header) SetMinLength(minLength *uint64) error {
	if err := h.Trix.RunHooks("minLength", h.minLength, minLength); err != nil {
		return err
	}
	h.minLength = minLength
	return nil
}
func (h *Header) SetPattern(pattern string) error {
	if err := h.Trix.RunHooks("pattern", h.pattern, pattern); err != nil {
		return err
	}
	h.pattern = pattern
	return nil
}
func (h *Header) SetMaxItems(maxItems *uint64) error {
	if err := h.Trix.RunHooks("maxItems", h.maxItems, maxItems); err != nil {
		return err
	}
	h.maxItems = maxItems
	return nil
}
func (h *Header) SetMinItems(minItems *uint64) error {
	if err := h.Trix.RunHooks("minItems", h.minItems, minItems); err != nil {
		return err
	}
	h.minItems = minItems
	return nil
}
func (h *Header) SetUniqueItems(uniqueItems bool) error {
	if err := h.Trix.RunHooks("uniqueItems", h.uniqueItems, uniqueItems); err != nil {
		return err
	}
	h.uniqueItems = uniqueItems
	return nil
}
func (h *Header) SetEnum(enum []interface{}) error {
	if err := h.Trix.RunHooks("enum", h.enum, enum); err != nil {
		return err
	}
	h.enum = enum
	return nil
}
func (h *Header) SetMultipleOf(multipleOf *float64) error {
	if err := h.Trix.RunHooks("multipleOf", h.multipleOf, multipleOf); err != nil {
		return err
	}
	h.multipleOf = multipleOf
	return nil
}

// HeaderFields groups all header properties for the constructor.
type HeaderFields struct {
	Description      string
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

// NewHeader creates a new Header instance.
func NewHeader(f HeaderFields) *Header {
	return &Header{
		description: f.Description, headerType: f.Type, format: f.Format,
		items: f.Items, collectionFormat: f.CollectionFormat,
		defaultVal: f.Default, maximum: f.Maximum,
		exclusiveMaximum: f.ExclusiveMaximum, minimum: f.Minimum,
		exclusiveMinimum: f.ExclusiveMinimum, maxLength: f.MaxLength,
		minLength: f.MinLength, pattern: f.Pattern, maxItems: f.MaxItems,
		minItems: f.MinItems, uniqueItems: f.UniqueItems, enum: f.Enum,
		multipleOf: f.MultipleOf,
	}
}

func (h *Header) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "description", Value: h.description},
		{Key: "type", Value: h.headerType},
		{Key: "format", Value: h.format},
		{Key: "items", Value: h.items},
		{Key: "collectionFormat", Value: h.collectionFormat},
		{Key: "default", Value: h.defaultVal},
		{Key: "maximum", Value: h.maximum},
		{Key: "exclusiveMaximum", Value: h.exclusiveMaximum},
		{Key: "minimum", Value: h.minimum},
		{Key: "exclusiveMinimum", Value: h.exclusiveMinimum},
		{Key: "maxLength", Value: h.maxLength},
		{Key: "minLength", Value: h.minLength},
		{Key: "pattern", Value: h.pattern},
		{Key: "maxItems", Value: h.maxItems},
		{Key: "minItems", Value: h.minItems},
		{Key: "uniqueItems", Value: h.uniqueItems},
		{Key: "enum", Value: h.enum},
		{Key: "multipleOf", Value: h.multipleOf},
	}
	return shared.AppendExtensions(fields, h.VendorExtensions)
}

// MarshalFields implements shared.MarshalFieldsProvider for export.
func (h *Header) MarshalFields() []shared.Field { return h.marshalFields() }

func (h *Header) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(h.marshalFields())
}

func (h *Header) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(h.marshalFields())
}

var _ yaml.Marshaler = (*Header)(nil)
