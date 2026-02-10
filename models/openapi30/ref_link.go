package openapi30

// LinkRef represents a reference to a Link or an inline Link.
type LinkRef struct {
	Node            // embedded - provides VendorExtensions and Trix
	Ref      string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Value    *Link  `json:"-" yaml:"-"`
	Circular bool   `json:"-" yaml:"-"` // true if circular reference detected
}

// NewLinkRef creates a new LinkRef instance.
func NewLinkRef(ref string) *LinkRef {
	return &LinkRef{Ref: ref}
}
