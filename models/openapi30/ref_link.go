package openapi30

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

// LinkRef represents a reference to a Link or an inline Link.
type LinkRef struct {
	Node            // embedded - provides VendorExtensions and Trix
	Ref      string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	value    *Link
	circular bool
	done     chan struct{}
	err      error
}

// NewLinkRef creates a new LinkRef instance.
func NewLinkRef(ref string) *LinkRef {
	return &LinkRef{Ref: ref}
}

func (r *LinkRef) Value() *Link {
	if r.done != nil {
		<-r.done
	}
	return r.value
}

func (r *LinkRef) Circular() bool {
	if r.done != nil {
		<-r.done
	}
	return r.circular
}

func (r *LinkRef) ResolveErr() error {
	if r.done != nil {
		<-r.done
	}
	return r.err
}

func (r *LinkRef) RawValue() *Link         { return r.value }
func (r *LinkRef) RawCircular() bool       { return r.circular }
func (r *LinkRef) SetValue(v *Link)        { r.value = v }
func (r *LinkRef) SetCircular(c bool)      { r.circular = c }
func (r *LinkRef) SetResolveErr(err error) { r.err = err }
func (r *LinkRef) InitDone()               { r.done = make(chan struct{}) }
func (r *LinkRef) MarkDone() {
	if r.done != nil {
		close(r.done)
	}
}
func (r *LinkRef) Done() <-chan struct{} { return r.done }

func (r *LinkRef) MarshalJSON() ([]byte, error) {
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

func (r *LinkRef) MarshalYAML() (interface{}, error) {
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

var _ yaml.Marshaler = (*LinkRef)(nil)
