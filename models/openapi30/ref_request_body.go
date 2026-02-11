package openapi30

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

// RequestBodyRef represents a reference to a RequestBody or an inline RequestBody.
type RequestBodyRef struct {
	Node                  // embedded - provides VendorExtensions and Trix
	Ref      string       `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Value    *RequestBody `json:"-" yaml:"-"`
	Circular bool         `json:"-" yaml:"-"` // true if circular reference detected
}

// NewRequestBodyRef creates a new RequestBodyRef instance.
func NewRequestBodyRef(ref string) *RequestBodyRef {
	return &RequestBodyRef{Ref: ref}
}

func (r *RequestBodyRef) MarshalJSON() ([]byte, error) {
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

func (r *RequestBodyRef) MarshalYAML() (interface{}, error) {
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

var _ yaml.Marshaler = (*RequestBodyRef)(nil)
