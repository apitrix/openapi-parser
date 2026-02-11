package openapi30

import (
	"encoding/json"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestLicense_MarshalJSON_AllFields(t *testing.T) {
	// Arrange
	l := NewLicense("MIT", "https://opensource.org/licenses/MIT")

	// Act
	got, err := json.Marshal(l)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"name":"MIT","url":"https://opensource.org/licenses/MIT"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestLicense_MarshalJSON_OmitsEmpty(t *testing.T) {
	// Arrange
	l := NewLicense("MIT", "")

	// Act
	got, err := json.Marshal(l)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"name":"MIT"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestLicense_MarshalYAML(t *testing.T) {
	// Arrange
	l := NewLicense("MIT", "https://opensource.org/licenses/MIT")

	// Act
	node, err := l.MarshalYAML()

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	yamlNode := node.(*yaml.Node)
	if yamlNode.Kind != yaml.MappingNode {
		t.Fatalf("expected MappingNode, got %d", yamlNode.Kind)
	}
	if len(yamlNode.Content) != 4 {
		t.Errorf("expected 4 content nodes (2 pairs), got %d", len(yamlNode.Content))
	}
}
