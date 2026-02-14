package openapi31

import (
	"openapi-parser/models/shared"
	"sort"

	"gopkg.in/yaml.v3"
)

// Responses is a container for expected responses of an operation.
// https://spec.openapis.org/oas/v3.1.0#responses-object
type Responses struct {
	ElementBase // embedded - provides VendorExtensions and Trix

	defaultResp *shared.RefWithMeta[Response]
	codes       map[string]*shared.RefWithMeta[Response]
}

func (r *Responses) Default() *shared.RefWithMeta[Response]          { return r.defaultResp }
func (r *Responses) Codes() map[string]*shared.RefWithMeta[Response] { return r.codes }

func (r *Responses) SetDefault(defaultResp *shared.RefWithMeta[Response]) error {
	if err := r.Trix.RunHooks("default", r.defaultResp, defaultResp); err != nil {
		return err
	}
	r.defaultResp = defaultResp
	return nil
}
func (r *Responses) SetCodes(codes map[string]*shared.RefWithMeta[Response]) error {
	if err := r.Trix.RunHooks("codes", r.codes, codes); err != nil {
		return err
	}
	r.codes = codes
	return nil
}

// NewResponses creates a new Responses instance.
func NewResponses(defaultResp *shared.RefWithMeta[Response], codes map[string]*shared.RefWithMeta[Response]) *Responses {
	return &Responses{defaultResp: defaultResp, codes: codes}
}

func (r *Responses) marshalFields() []shared.Field {
	fields := []shared.Field{
		{Key: "default", Value: r.defaultResp},
	}
	if len(r.codes) > 0 {
		keys := make([]string, 0, len(r.codes))
		for k := range r.codes {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			fields = append(fields, shared.Field{Key: k, Value: r.codes[k]})
		}
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
