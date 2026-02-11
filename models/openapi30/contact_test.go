package openapi30

import (
	"encoding/json"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestContact_MarshalJSON_AllFields(t *testing.T) {
	// Arrange
	c := NewContact("John", "https://example.com", "john@example.com")

	// Act
	got, err := json.Marshal(c)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"name":"John","url":"https://example.com","email":"john@example.com"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestContact_MarshalJSON_OmitsEmpty(t *testing.T) {
	// Arrange
	c := NewContact("John", "", "")

	// Act
	got, err := json.Marshal(c)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"name":"John"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestContact_MarshalJSON_WithExtensions(t *testing.T) {
	// Arrange
	c := NewContact("John", "", "")
	c.VendorExtensions = map[string]interface{}{"x-custom": "val"}

	// Act
	got, err := json.Marshal(c)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"name":"John","x-custom":"val"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestContact_MarshalYAML(t *testing.T) {
	// Arrange
	c := NewContact("John", "https://example.com", "")

	// Act
	node, err := c.MarshalYAML()

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	yamlNode := node.(*yaml.Node)
	if yamlNode.Kind != yaml.MappingNode {
		t.Fatalf("expected MappingNode, got %d", yamlNode.Kind)
	}
	// Should have 2 pairs (name + url), email omitted
	if len(yamlNode.Content) != 4 {
		t.Errorf("expected 4 content nodes (2 pairs), got %d", len(yamlNode.Content))
	}
}
