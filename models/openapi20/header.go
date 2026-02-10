package openapi20

// Header represents a Header Object in a response.
// https://swagger.io/specification/v2/#header-object
type Header struct {
	Node // embedded - provides VendorExtensions and Trix

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
