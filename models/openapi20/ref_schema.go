package openapi20

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

// SchemaRef represents a reference to a Schema or an inline Schema.
type SchemaRef struct {
	Node             // embedded - provides VendorExtensions and Trix
	Ref      string  `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Value    *Schema `json:"-" yaml:"-"`
	Circular bool    `json:"-" yaml:"-"` // true if circular reference detected
}

// NewSchemaRef creates a new SchemaRef instance.
func NewSchemaRef(ref string) *SchemaRef {
	return &SchemaRef{Ref: ref}
}

func (r *SchemaRef) MarshalJSON() ([]byte, error) {
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

func (r *SchemaRef) MarshalYAML() (interface{}, error) {
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

var _ yaml.Marshaler = (*SchemaRef)(nil)
