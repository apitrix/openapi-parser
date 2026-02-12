package openapi30

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

// HeaderRef represents a reference to a Header or an inline Header.
type HeaderRef struct {
	Node            // embedded - provides VendorExtensions and Trix
	Ref      string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	value    *Header
	circular bool
	done     chan struct{}
	err      error
}

// NewHeaderRef creates a new HeaderRef instance.
func NewHeaderRef(ref string) *HeaderRef {
	return &HeaderRef{Ref: ref}
}

func (r *HeaderRef) Value() *Header {
	if r.done != nil {
		<-r.done
	}
	return r.value
}

func (r *HeaderRef) Circular() bool {
	if r.done != nil {
		<-r.done
	}
	return r.circular
}

func (r *HeaderRef) ResolveErr() error {
	if r.done != nil {
		<-r.done
	}
	return r.err
}

func (r *HeaderRef) RawValue() *Header       { return r.value }
func (r *HeaderRef) RawCircular() bool       { return r.circular }
func (r *HeaderRef) SetValue(v *Header)      { r.value = v }
func (r *HeaderRef) SetCircular(c bool)      { r.circular = c }
func (r *HeaderRef) SetResolveErr(err error) { r.err = err }
func (r *HeaderRef) InitDone()               { r.done = make(chan struct{}) }
func (r *HeaderRef) MarkDone() {
	if r.done != nil {
		close(r.done)
	}
}
func (r *HeaderRef) Done() <-chan struct{} { return r.done }

func (r *HeaderRef) MarshalJSON() ([]byte, error) {
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

func (r *HeaderRef) MarshalYAML() (interface{}, error) {
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

var _ yaml.Marshaler = (*HeaderRef)(nil)
