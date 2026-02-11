package openapi30

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

// PathItemRef represents a reference to a PathItem or an inline PathItem.
type PathItemRef struct {
	Node               // embedded - provides VendorExtensions and Trix
	Ref      string    `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Value    *PathItem `json:"-" yaml:"-"`
	Circular bool      `json:"-" yaml:"-"` // true if circular reference detected
}

// NewPathItemRef creates a new PathItemRef instance.
func NewPathItemRef(ref string) *PathItemRef {
	return &PathItemRef{Ref: ref}
}

func (r *PathItemRef) MarshalJSON() ([]byte, error) {
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
	if r.Value != nil {
		return r.Value.MarshalYAML()
	}
	return nil, nil
}

var _ yaml.Marshaler = (*PathItemRef)(nil)
