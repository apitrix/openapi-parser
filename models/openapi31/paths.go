package openapi31

// Paths holds the relative paths to individual endpoints.
// https://spec.openapis.org/oas/v3.1.0#paths-object
type Paths struct {
	Node // embedded - provides VendorExtensions and Trix

	// Items maps paths to their definitions.
	Items map[string]*PathItem `json:"-" yaml:"-"`
}

// NewPaths creates a new Paths instance with the given items map.
func NewPaths(items map[string]*PathItem) *Paths {
	return &Paths{Items: items}
}
