package openapi20

// Responses is a container for expected responses of an operation.
// https://swagger.io/specification/v2/#responses-object
type Responses struct {
	Node // embedded - provides VendorExtensions and Trix

	defaultResp *ResponseRef
	codes       map[string]*ResponseRef // HTTP status codes (e.g., "200", "404")
}

func (r *Responses) Default() *ResponseRef          { return r.defaultResp }
func (r *Responses) Codes() map[string]*ResponseRef { return r.codes }

// NewResponses creates a new Responses instance.
func NewResponses(defaultResp *ResponseRef, codes map[string]*ResponseRef) *Responses {
	return &Responses{defaultResp: defaultResp, codes: codes}
}
