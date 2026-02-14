package openapi20

import (
	"openapi-parser/models/shared"
	"sort"

	"gopkg.in/yaml.v3"
)

// Responses is a container for expected responses of an operation.
// https://swagger.io/specification/v2/#responses-object
type Responses struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	defaultResp *RefResponse
	codes       map[string]*RefResponse // HTTP status codes (e.g., "200", "404")
}

func (r *Responses) Default() *RefResponse          { return r.defaultResp }
func (r *Responses) Codes() map[string]*RefResponse { return r.codes }

func (r *Responses) SetDefault(defaultResp *RefResponse) error {
	if err := r.Trix.RunHooks("default", r.defaultResp, defaultResp); err != nil {
		return err
	}
	r.defaultResp = defaultResp
	return nil
}
func (r *Responses) SetCodes(codes map[string]*RefResponse) error {
	if err := r.Trix.RunHooks("codes", r.codes, codes); err != nil {
		return err
	}
	r.codes = codes
	return nil
}

// NewResponses creates a new Responses instance.
func NewResponses(defaultResp *RefResponse, codes map[string]*RefResponse) *Responses {
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

// MarshalFields implements shared.MarshalFieldsProvider for export.
func (r *Responses) MarshalFields() []shared.Field { return r.marshalFields() }

func (r *Responses) MarshalJSON() ([]byte, error) {
	return shared.MarshalFieldsJSON(r.marshalFields())
}

func (r *Responses) MarshalYAML() (interface{}, error) {
	return shared.MarshalFieldsYAML(r.marshalFields())
}

var _ yaml.Marshaler = (*Responses)(nil)
