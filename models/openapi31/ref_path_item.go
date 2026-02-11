package openapi31

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

// PathItemRef represents a reference to a PathItem or an inline PathItem.
type PathItemRef struct {
	Node                  // embedded - provides VendorExtensions and Trix
	Ref         string    `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Summary     string    `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string    `json:"description,omitempty" yaml:"description,omitempty"`
	Value       *PathItem `json:"-" yaml:"-"`
	Circular    bool      `json:"-" yaml:"-"` // true if circular reference detected
}

// NewPathItemRef creates a new PathItemRef instance.
func NewPathItemRef(ref string) *PathItemRef {
	return &PathItemRef{Ref: ref}
}

func (r *PathItemRef) MarshalJSON() ([]byte, error) {
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

func (r *PathItemRef) MarshalYAML() (interface{}, error) {
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

var _ yaml.Marshaler = (*PathItemRef)(nil)
