package openapi30

// Responses is a container for expected responses of an operation.
// https://spec.openapis.org/oas/v3.0.3#responses-object
type Responses struct {
	Node // embedded - provides VendorExtensions and Trix

	defaultResp *ResponseRef
	codes       map[string]*ResponseRef
}

func (r *Responses) Default() *ResponseRef          { return r.defaultResp }
func (r *Responses) Codes() map[string]*ResponseRef { return r.codes }

// NewResponses creates a new Responses instance.
func NewResponses(defaultResp *ResponseRef, codes map[string]*ResponseRef) *Responses {
	return &Responses{defaultResp: defaultResp, codes: codes}
}
