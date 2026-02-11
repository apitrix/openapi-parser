package openapi31

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

// SchemaType represents a JSON Schema type field that can be either a single
// string or an array of strings (JSON Schema Draft 2020-12).
type SchemaType struct {
	// Single is set when the type is a single string value (e.g. "string").
	Single string
	// Array is set when the type is an array of strings (e.g. ["string", "null"]).
	Array []string
}

// IsEmpty returns true if no type was specified.
func (t SchemaType) IsEmpty() bool {
	return t.Single == "" && len(t.Array) == 0
}

// Values returns all type values as a slice, whether specified as single or array.
func (t SchemaType) Values() []string {
	if len(t.Array) > 0 {
		return t.Array
	}
	if t.Single != "" {
		return []string{t.Single}
	}
	return nil
}

func (t SchemaType) MarshalJSON() ([]byte, error) {
	if len(t.Array) > 0 {
		return json.Marshal(t.Array)
	}
	if t.Single != "" {
		return json.Marshal(t.Single)
	}
	return []byte("null"), nil
}

func (t SchemaType) MarshalYAML() (interface{}, error) {
	if len(t.Array) > 0 {
		node := &yaml.Node{Kind: yaml.SequenceNode, Tag: "!!seq"}
		for _, v := range t.Array {
			node.Content = append(node.Content, &yaml.Node{
				Kind: yaml.ScalarNode, Tag: "!!str", Value: v,
			})
		}
		return node, nil
	}
	if t.Single != "" {
		return &yaml.Node{Kind: yaml.ScalarNode, Tag: "!!str", Value: t.Single}, nil
	}
	return nil, nil
}

var _ yaml.Marshaler = SchemaType{}
