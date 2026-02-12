package openapi30

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

// RequestBodyRef represents a reference to a RequestBody or an inline RequestBody.
type RequestBodyRef struct {
	Node            // embedded - provides VendorExtensions and Trix
	Ref      string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	value    *RequestBody
	circular bool
	done     chan struct{}
	err      error
}

// NewRequestBodyRef creates a new RequestBodyRef instance.
func NewRequestBodyRef(ref string) *RequestBodyRef {
	return &RequestBodyRef{Ref: ref}
}

func (r *RequestBodyRef) Value() *RequestBody {
	if r.done != nil {
		<-r.done
	}
	return r.value
}

func (r *RequestBodyRef) Circular() bool {
	if r.done != nil {
		<-r.done
	}
	return r.circular
}

func (r *RequestBodyRef) ResolveErr() error {
	if r.done != nil {
		<-r.done
	}
	return r.err
}

func (r *RequestBodyRef) RawValue() *RequestBody  { return r.value }
func (r *RequestBodyRef) RawCircular() bool       { return r.circular }
func (r *RequestBodyRef) SetValue(v *RequestBody) { r.value = v }
func (r *RequestBodyRef) SetCircular(c bool)      { r.circular = c }
func (r *RequestBodyRef) SetResolveErr(err error) { r.err = err }
func (r *RequestBodyRef) InitDone()               { r.done = make(chan struct{}) }
func (r *RequestBodyRef) MarkDone() {
	if r.done != nil {
		close(r.done)
	}
}
func (r *RequestBodyRef) Done() <-chan struct{} { return r.done }

func (r *RequestBodyRef) MarshalJSON() ([]byte, error) {
	if r.Ref != "" {
		return json.Marshal(struct {
			Ref string `json:"$ref"`
		}{Ref: r.Ref})
	}
	if r.value != nil {
		return r.value.MarshalJSON()
	}
	return []byte("null"), nil
}

func (r *RequestBodyRef) MarshalYAML() (interface{}, error) {
	if r.Ref != "" {
		return &yaml.Node{
			Kind: yaml.MappingNode, Tag: "!!map",
			Content: []*yaml.Node{
				{Kind: yaml.ScalarNode, Tag: "!!str", Value: "$ref"},
				{Kind: yaml.ScalarNode, Tag: "!!str", Value: r.Ref},
			},
		}, nil
	}
	if r.value != nil {
		return r.value.MarshalYAML()
	}
	return nil, nil
}

var _ yaml.Marshaler = (*RequestBodyRef)(nil)
