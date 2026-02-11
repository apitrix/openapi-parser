package openapi30

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

// ExampleRef represents a reference to an Example or an inline Example.
type ExampleRef struct {
	Node              // embedded - provides VendorExtensions and Trix
	Ref      string   `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Value    *Example `json:"-" yaml:"-"`
	Circular bool     `json:"-" yaml:"-"` // true if circular reference detected
}

// NewExampleRef creates a new ExampleRef instance.
func NewExampleRef(ref string) *ExampleRef {
	return &ExampleRef{Ref: ref}
}

func (r *ExampleRef) MarshalJSON() ([]byte, error) {
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
	if r.Value != nil {
		return r.Value.MarshalYAML()
	}
	return nil, nil
}

var _ yaml.Marshaler = (*ExampleRef)(nil)
