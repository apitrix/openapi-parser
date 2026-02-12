package openapi30

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

// PathItemRef represents a reference to a PathItem or an inline PathItem.
type PathItemRef struct {
	Node            // embedded - provides VendorExtensions and Trix
	Ref      string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	value    *PathItem
	circular bool
	done     chan struct{}
	err      error
}

// NewPathItemRef creates a new PathItemRef instance.
func NewPathItemRef(ref string) *PathItemRef {
	return &PathItemRef{Ref: ref}
}

func (r *PathItemRef) Value() *PathItem {
	if r.done != nil {
		<-r.done
	}
	return r.value
}

func (r *PathItemRef) Circular() bool {
	if r.done != nil {
		<-r.done
	}
	return r.circular
}

func (r *PathItemRef) ResolveErr() error {
	if r.done != nil {
		<-r.done
	}
	return r.err
}

func (r *PathItemRef) RawValue() *PathItem     { return r.value }
func (r *PathItemRef) RawCircular() bool       { return r.circular }
func (r *PathItemRef) SetValue(v *PathItem)    { r.value = v }
func (r *PathItemRef) SetCircular(c bool)      { r.circular = c }
func (r *PathItemRef) SetResolveErr(err error) { r.err = err }
func (r *PathItemRef) InitDone()               { r.done = make(chan struct{}) }
func (r *PathItemRef) MarkDone() {
	if r.done != nil {
		close(r.done)
	}
}
func (r *PathItemRef) Done() <-chan struct{} { return r.done }

func (r *PathItemRef) MarshalJSON() ([]byte, error) {
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

func (r *PathItemRef) MarshalYAML() (interface{}, error) {
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

var _ yaml.Marshaler = (*PathItemRef)(nil)
