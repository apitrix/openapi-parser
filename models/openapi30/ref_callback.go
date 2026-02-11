package openapi30

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

// CallbackRef represents a reference to a Callback or an inline Callback.
type CallbackRef struct {
	Node               // embedded - provides VendorExtensions and Trix
	Ref      string    `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Value    *Callback `json:"-" yaml:"-"`
	Circular bool      `json:"-" yaml:"-"` // true if circular reference detected
}

// NewCallbackRef creates a new CallbackRef instance.
func NewCallbackRef(ref string) *CallbackRef {
	return &CallbackRef{Ref: ref}
}

func (r *CallbackRef) MarshalJSON() ([]byte, error) {
	if r.Ref != "" {
		return json.Marshal(struct {
			Ref string `json:"$ref"`
		}{Ref: r.Ref})
	}
	if r.Value != nil {
		return r.Value.MarshalJSON()
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
	if r.Value != nil {
		return r.Value.MarshalYAML()
	}
	return nil, nil
}

var _ yaml.Marshaler = (*CallbackRef)(nil)
