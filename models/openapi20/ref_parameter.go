package openapi20

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

// ParameterRef represents a reference to a Parameter or an inline Parameter.
type ParameterRef struct {
	Node                // embedded - provides VendorExtensions and Trix
	Ref      string     `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Value    *Parameter `json:"-" yaml:"-"`
	Circular bool       `json:"-" yaml:"-"` // true if circular reference detected
}

// NewParameterRef creates a new ParameterRef instance.
func NewParameterRef(ref string) *ParameterRef {
	return &ParameterRef{Ref: ref}
}

func (r *ParameterRef) MarshalJSON() ([]byte, error) {
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

func (r *ParameterRef) MarshalYAML() (interface{}, error) {
	if r.Ref != "" {
		return &yaml.Node{
			Kind: yaml.MappingNode,
			Tag:  "!!map",
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

var _ yaml.Marshaler = (*ParameterRef)(nil)
