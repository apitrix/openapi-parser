package openapi30

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

// SecuritySchemeRef represents a reference to a SecurityScheme or an inline SecurityScheme.
type SecuritySchemeRef struct {
	Node                     // embedded - provides VendorExtensions and Trix
	Ref      string          `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Value    *SecurityScheme `json:"-" yaml:"-"`
	Circular bool            `json:"-" yaml:"-"` // true if circular reference detected
}

// NewSecuritySchemeRef creates a new SecuritySchemeRef instance.
func NewSecuritySchemeRef(ref string) *SecuritySchemeRef {
	return &SecuritySchemeRef{Ref: ref}
}

func (r *SecuritySchemeRef) MarshalJSON() ([]byte, error) {
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
	if r.Value != nil {
		return r.Value.MarshalYAML()
	}
	return nil, nil
}

var _ yaml.Marshaler = (*SecuritySchemeRef)(nil)
