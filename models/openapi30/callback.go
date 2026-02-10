package openapi30

// Callback is a map of possible out-of-band callbacks related to the parent operation.
// https://spec.openapis.org/oas/v3.0.3#callback-object
type Callback struct {
	Node // embedded - provides VendorExtensions and Trix

	// Paths maps runtime expressions to PathItem objects.
	// The key is a runtime expression that identifies the URL to use for the callback request.
	Paths map[string]*PathItem `json:"-" yaml:"-"`
}

// NewCallback creates a new Callback instance.
func NewCallback() *Callback {
	return &Callback{}
}
