package openapi30

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

// ExampleRef represents a reference to an Example or an inline Example.
type ExampleRef struct {
	Node            // embedded - provides VendorExtensions and Trix
	Ref      string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	value    *Example
	circular bool
	done     chan struct{}
	err      error
}

// NewExampleRef creates a new ExampleRef instance.
func NewExampleRef(ref string) *ExampleRef {
	return &ExampleRef{Ref: ref}
}

func (r *ExampleRef) Value() *Example {
	if r.done != nil {
		<-r.done
	}
	return r.value
}

func (r *ExampleRef) Circular() bool {
	if r.done != nil {
		<-r.done
	}
	return r.circular
}

func (r *ExampleRef) ResolveErr() error {
	if r.done != nil {
		<-r.done
	}
	return r.err
}

func (r *ExampleRef) RawValue() *Example      { return r.value }
func (r *ExampleRef) RawCircular() bool       { return r.circular }
func (r *ExampleRef) SetValue(v *Example)     { r.value = v }
func (r *ExampleRef) SetCircular(c bool)      { r.circular = c }
func (r *ExampleRef) SetResolveErr(err error) { r.err = err }
func (r *ExampleRef) InitDone()               { r.done = make(chan struct{}) }
func (r *ExampleRef) MarkDone() {
	if r.done != nil {
		close(r.done)
	}
}
func (r *ExampleRef) Done() <-chan struct{} { return r.done }

func (r *ExampleRef) MarshalJSON() ([]byte, error) {
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

func (r *ExampleRef) MarshalYAML() (interface{}, error) {
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

var _ yaml.Marshaler = (*ExampleRef)(nil)
