package openapi20

// Items describes the type of items in an array parameter.
// https://swagger.io/specification/v2/#items-object
type Items struct {
	Node // embedded - provides VendorExtensions and Trix

	Type             string        `json:"type" yaml:"type"`
	Format           string        `json:"format,omitempty" yaml:"format,omitempty"`
	Items            *Items        `json:"items,omitempty" yaml:"items,omitempty"`
	CollectionFormat string        `json:"collectionFormat,omitempty" yaml:"collectionFormat,omitempty"`
	Default          interface{}   `json:"default,omitempty" yaml:"default,omitempty"`
	Maximum          *float64      `json:"maximum,omitempty" yaml:"maximum,omitempty"`
	ExclusiveMaximum bool          `json:"exclusiveMaximum,omitempty" yaml:"exclusiveMaximum,omitempty"`
	Minimum          *float64      `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	ExclusiveMinimum bool          `json:"exclusiveMinimum,omitempty" yaml:"exclusiveMinimum,omitempty"`
	MaxLength        *uint64       `json:"maxLength,omitempty" yaml:"maxLength,omitempty"`
	MinLength        *uint64       `json:"minLength,omitempty" yaml:"minLength,omitempty"`
	Pattern          string        `json:"pattern,omitempty" yaml:"pattern,omitempty"`
	MaxItems         *uint64       `json:"maxItems,omitempty" yaml:"maxItems,omitempty"`
	MinItems         *uint64       `json:"minItems,omitempty" yaml:"minItems,omitempty"`
	UniqueItems      bool          `json:"uniqueItems,omitempty" yaml:"uniqueItems,omitempty"`
	Enum             []interface{} `json:"enum,omitempty" yaml:"enum,omitempty"`
	MultipleOf       *float64      `json:"multipleOf,omitempty" yaml:"multipleOf,omitempty"`
}

// NewItems creates a new Items instance.
func NewItems(itemType string) *Items {
	return &Items{Type: itemType}
}
