package openapi30

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

// CallbackRef represents a reference to a Callback or an inline Callback.
type CallbackRef struct {
	Node            // embedded - provides VendorExtensions and Trix
	Ref      string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	value    *Callback
	circular bool
	done     chan struct{}
	err      error
}

// NewCallbackRef creates a new CallbackRef instance.
func NewCallbackRef(ref string) *CallbackRef {
	return &CallbackRef{Ref: ref}
}

func (r *CallbackRef) Value() *Callback {
	if r.done != nil {
		<-r.done
	}
	return r.value
}

func (r *CallbackRef) Circular() bool {
	if r.done != nil {
		<-r.done
	}
	return r.circular
}

func (r *CallbackRef) ResolveErr() error {
	if r.done != nil {
		<-r.done
	}
	return r.err
}

func (r *CallbackRef) RawValue() *Callback     { return r.value }
func (r *CallbackRef) RawCircular() bool       { return r.circular }
func (r *CallbackRef) SetValue(v *Callback)    { r.value = v }
func (r *CallbackRef) SetCircular(c bool)      { r.circular = c }
func (r *CallbackRef) SetResolveErr(err error) { r.err = err }
func (r *CallbackRef) InitDone()               { r.done = make(chan struct{}) }
func (r *CallbackRef) MarkDone() {
	if r.done != nil {
		close(r.done)
	}
}
func (r *CallbackRef) Done() <-chan struct{} { return r.done }

func (r *CallbackRef) MarshalJSON() ([]byte, error) {
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

func (r *CallbackRef) MarshalYAML() (interface{}, error) {
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

var _ yaml.Marshaler = (*CallbackRef)(nil)
