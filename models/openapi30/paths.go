package openapi30

// Paths holds the relative paths to individual endpoints.
// https://spec.openapis.org/oas/v3.0.3#paths-object
type Paths struct {
	Node // embedded - provides VendorExtensions and Trix

	// items maps paths to their definitions.
	items map[string]*PathItem
}

func (p *Paths) Items() map[string]*PathItem { return p.items }

// NewPaths creates a new Paths instance with the given items map.
func NewPaths(items map[string]*PathItem) *Paths {
	return &Paths{items: items}
}
