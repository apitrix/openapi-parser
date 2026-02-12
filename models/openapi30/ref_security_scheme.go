package openapi30

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

// SecuritySchemeRef represents a reference to a SecurityScheme or an inline SecurityScheme.
type SecuritySchemeRef struct {
	Node            // embedded - provides VendorExtensions and Trix
	Ref      string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	value    *SecurityScheme
	circular bool
	done     chan struct{}
	err      error
}

// NewSecuritySchemeRef creates a new SecuritySchemeRef instance.
func NewSecuritySchemeRef(ref string) *SecuritySchemeRef {
	return &SecuritySchemeRef{Ref: ref}
}

func (r *SecuritySchemeRef) Value() *SecurityScheme {
	if r.done != nil {
		<-r.done
	}
	return r.value
}

func (r *SecuritySchemeRef) Circular() bool {
	if r.done != nil {
		<-r.done
	}
	return r.circular
}

func (r *SecuritySchemeRef) ResolveErr() error {
	if r.done != nil {
		<-r.done
	}
	return r.err
}

func (r *SecuritySchemeRef) RawValue() *SecurityScheme  { return r.value }
func (r *SecuritySchemeRef) RawCircular() bool          { return r.circular }
func (r *SecuritySchemeRef) SetValue(v *SecurityScheme) { r.value = v }
func (r *SecuritySchemeRef) SetCircular(c bool)         { r.circular = c }
func (r *SecuritySchemeRef) SetResolveErr(err error)    { r.err = err }
func (r *SecuritySchemeRef) InitDone()                  { r.done = make(chan struct{}) }
func (r *SecuritySchemeRef) MarkDone() {
	if r.done != nil {
		close(r.done)
	}
}
func (r *SecuritySchemeRef) Done() <-chan struct{} { return r.done }

func (r *SecuritySchemeRef) MarshalJSON() ([]byte, error) {
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

func (r *SecuritySchemeRef) MarshalYAML() (interface{}, error) {
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

var _ yaml.Marshaler = (*SecuritySchemeRef)(nil)
