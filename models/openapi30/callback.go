package openapi30

// Callback is a map of possible out-of-band callbacks related to the parent operation.
// https://spec.openapis.org/oas/v3.0.3#callback-object
type Callback struct {
	Node // embedded - provides VendorExtensions and Trix

	// paths maps runtime expressions to PathItem objects.
	paths map[string]*PathItem
}

func (c *Callback) Paths() map[string]*PathItem { return c.paths }

// NewCallback creates a new Callback instance.
func NewCallback(paths map[string]*PathItem) *Callback {
	return &Callback{paths: paths}
}
