package openapi20

import (
	"openapi-parser/models/shared"
	"sort"

	"gopkg.in/yaml.v3"
)

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

func (r *Responses) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "default", Value: r.defaultResp},
	}
	keys := make([]string, 0, len(r.codes))
	for k := range r.codes {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fields = append(fields, shared.Field{Key: k, Value: r.codes[k]})
	}
	return shared.AppendExtensions(fields, r.VendorExtensions)
}

func (r *Responses) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(r.marshalFields())
}

func (r *Responses) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(r.marshalFields())
}

var _ yaml.Marshaler = (*Responses)(nil)
