package openapi31

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

// ParameterRef represents a reference to a Parameter or an inline Parameter.
type ParameterRef struct {
	Node                   // embedded - provides VendorExtensions and Trix
	Ref         string     `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Summary     string     `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string     `json:"description,omitempty" yaml:"description,omitempty"`
	Value       *Parameter `json:"-" yaml:"-"`
	Circular    bool       `json:"-" yaml:"-"` // true if circular reference detected
}

// NewParameterRef creates a new ParameterRef instance.
func NewParameterRef(ref string) *ParameterRef {
	return &ParameterRef{Ref: ref}
}

func (r *ParameterRef) MarshalJSON() ([]byte, error) {
	if r.Ref != "" {
		m := map[string]string{"$ref": r.Ref}
		if r.Summary != "" {
			m["summary"] = r.Summary
		}
		if r.Description != "" {
			m["description"] = r.Description
		}
		return json.Marshal(m)
	}
	if r.Value != nil {
		return r.Value.MarshalJSON()
	}
	return []byte("null"), nil
}

func (r *ParameterRef) MarshalYAML() (interface{}, error) {
	if r.Ref != "" {
		content := []*yaml.Node{
			{Kind: yaml.ScalarNode, Tag: "!!str", Value: "$ref"},
			{Kind: yaml.ScalarNode, Tag: "!!str", Value: r.Ref},
		}
		if r.Summary != "" {
			content = append(content,
				&yaml.Node{Kind: yaml.ScalarNode, Tag: "!!str", Value: "summary"},
				&yaml.Node{Kind: yaml.ScalarNode, Tag: "!!str", Value: r.Summary},
			)
		}
		if r.Description != "" {
			content = append(content,
				&yaml.Node{Kind: yaml.ScalarNode, Tag: "!!str", Value: "description"},
				&yaml.Node{Kind: yaml.ScalarNode, Tag: "!!str", Value: r.Description},
			)
		}
		return &yaml.Node{Kind: yaml.MappingNode, Tag: "!!map", Content: content}, nil
	}
	if r.Value != nil {
		return r.Value.MarshalYAML()
	}
	return nil, nil
}

var _ yaml.Marshaler = (*ParameterRef)(nil)
