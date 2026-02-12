package openapi30

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

// ParameterRef represents a reference to a Parameter or an inline Parameter.
type ParameterRef struct {
	Node            // embedded - provides VendorExtensions and Trix
	Ref      string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	value    *Parameter
	circular bool
	done     chan struct{}
	err      error
}

// NewParameterRef creates a new ParameterRef instance.
func NewParameterRef(ref string) *ParameterRef {
	return &ParameterRef{Ref: ref}
}

func (r *ParameterRef) Value() *Parameter {
	if r.done != nil {
		<-r.done
	}
	return r.value
}

func (r *ParameterRef) Circular() bool {
	if r.done != nil {
		<-r.done
	}
	return r.circular
}

func (r *ParameterRef) ResolveErr() error {
	if r.done != nil {
		<-r.done
	}
	return r.err
}

func (r *ParameterRef) RawValue() *Parameter    { return r.value }
func (r *ParameterRef) RawCircular() bool       { return r.circular }
func (r *ParameterRef) SetValue(v *Parameter)   { r.value = v }
func (r *ParameterRef) SetCircular(c bool)      { r.circular = c }
func (r *ParameterRef) SetResolveErr(err error) { r.err = err }
func (r *ParameterRef) InitDone()               { r.done = make(chan struct{}) }
func (r *ParameterRef) MarkDone() {
	if r.done != nil {
		close(r.done)
	}
}
func (r *ParameterRef) Done() <-chan struct{} { return r.done }

func (r *ParameterRef) MarshalJSON() ([]byte, error) {
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

func (r *ParameterRef) MarshalYAML() (interface{}, error) {
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

var _ yaml.Marshaler = (*ParameterRef)(nil)
