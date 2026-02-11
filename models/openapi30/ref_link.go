package openapi30

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

// LinkRef represents a reference to a Link or an inline Link.
type LinkRef struct {
	Node            // embedded - provides VendorExtensions and Trix
	Ref      string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Value    *Link  `json:"-" yaml:"-"`
	Circular bool   `json:"-" yaml:"-"` // true if circular reference detected
}

// NewLinkRef creates a new LinkRef instance.
func NewLinkRef(ref string) *LinkRef {
	return &LinkRef{Ref: ref}
}

func (r *LinkRef) MarshalJSON() ([]byte, error) {
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
	if r.Value != nil {
		return r.Value.MarshalYAML()
	}
	return nil, nil
}

var _ yaml.Marshaler = (*LinkRef)(nil)
