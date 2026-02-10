package openapi20

// Responses is a container for expected responses of an operation.
// https://swagger.io/specification/v2/#responses-object
type Responses struct {
	Node // embedded - provides VendorExtensions and Trix

	Default *ResponseRef            `json:"default,omitempty" yaml:"default,omitempty"`
	Codes   map[string]*ResponseRef `json:"-" yaml:"-"` // HTTP status codes (e.g., "200", "404")
}

// NewResponses creates a new Responses instance.
func NewResponses() *Responses {
	return &Responses{}
}
