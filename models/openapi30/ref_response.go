package openapi30

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

// ResponseRef represents a reference to a Response or an inline Response.
type ResponseRef struct {
	Node            // embedded - provides VendorExtensions and Trix
	Ref      string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	value    *Response
	circular bool
	done     chan struct{}
	err      error
}

// NewResponseRef creates a new ResponseRef instance.
func NewResponseRef(ref string) *ResponseRef {
	return &ResponseRef{Ref: ref}
}

func (r *ResponseRef) Value() *Response {
	if r.done != nil {
		<-r.done
	}
	return r.value
}

func (r *ResponseRef) Circular() bool {
	if r.done != nil {
		<-r.done
	}
	return r.circular
}

func (r *ResponseRef) ResolveErr() error {
	if r.done != nil {
		<-r.done
	}
	return r.err
}

func (r *ResponseRef) RawValue() *Response     { return r.value }
func (r *ResponseRef) RawCircular() bool       { return r.circular }
func (r *ResponseRef) SetValue(v *Response)    { r.value = v }
func (r *ResponseRef) SetCircular(c bool)      { r.circular = c }
func (r *ResponseRef) SetResolveErr(err error) { r.err = err }
func (r *ResponseRef) InitDone()               { r.done = make(chan struct{}) }
func (r *ResponseRef) MarkDone() {
	if r.done != nil {
		close(r.done)
	}
}
func (r *ResponseRef) Done() <-chan struct{} { return r.done }

func (r *ResponseRef) MarshalJSON() ([]byte, error) {
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

func (r *ResponseRef) MarshalYAML() (interface{}, error) {
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

var _ yaml.Marshaler = (*ResponseRef)(nil)
