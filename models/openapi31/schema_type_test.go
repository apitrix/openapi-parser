package openapi31

import (
	"encoding/json"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestSchemaType_MarshalJSON_Single(t *testing.T) {
	st := SchemaType{Single: "string"}
	got, err := json.Marshal(st)
	if err != nil {
		t.Fatal(err)
	}
	want := `"string"`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestSchemaType_MarshalJSON_Array(t *testing.T) {
	st := SchemaType{Array: []string{"string", "null"}}
	got, err := json.Marshal(st)
	if err != nil {
		t.Fatal(err)
	}
	want := `["string","null"]`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestSchemaType_MarshalJSON_Empty(t *testing.T) {
	st := SchemaType{}
	got, err := json.Marshal(st)
	if err != nil {
		t.Fatal(err)
	}
	want := `null`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestSchemaType_MarshalJSON_ArrayTakesPriority(t *testing.T) {
	st := SchemaType{Single: "string", Array: []string{"integer", "null"}}
	got, err := json.Marshal(st)
	if err != nil {
		t.Fatal(err)
	}
	want := `["integer","null"]`
	if string(got) != want {
		t.Errorf("got %s, want %s (Array should take priority over Single)", got, want)
	}
}

func TestSchemaType_MarshalYAML_Single(t *testing.T) {
	st := SchemaType{Single: "object"}
	node, err := st.MarshalYAML()
	if err != nil {
		t.Fatal(err)
	}
	yamlNode := node.(*yaml.Node)
	if yamlNode.Kind != yaml.ScalarNode {
		t.Fatalf("expected ScalarNode, got %d", yamlNode.Kind)
	}
	if yamlNode.Value != "object" {
		t.Errorf("got %s, want object", yamlNode.Value)
	}
}

func TestSchemaType_MarshalYAML_Array(t *testing.T) {
	st := SchemaType{Array: []string{"string", "null"}}
	node, err := st.MarshalYAML()
	if err != nil {
		t.Fatal(err)
	}
	yamlNode := node.(*yaml.Node)
	if yamlNode.Kind != yaml.SequenceNode {
		t.Fatalf("expected SequenceNode, got %d", yamlNode.Kind)
	}
	if len(yamlNode.Content) != 2 {
		t.Errorf("expected 2 items, got %d", len(yamlNode.Content))
	}
}
