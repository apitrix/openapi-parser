package openapi30

import (
	"encoding/json"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestOpenAPI_MarshalJSON_Full(t *testing.T) {
	// Arrange
	info := NewInfo("Pet Store", "A sample API", "", "1.0.0", nil, nil)
	server := NewServer("https://api.example.com", "", nil)
	paths := NewPaths(map[string]*PathItem{
		"/pets": NewPathItem("", "", "", nil, nil, nil, nil, nil, nil, nil, nil, nil, nil),
	})
	doc := NewOpenAPI("3.0.3", info)
	doc.SetProperty("servers", []*Server{server})
	doc.SetProperty("paths", paths)

	// Act
	got, err := json.Marshal(doc)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"openapi", "info", "servers", "paths"} {
		if _, ok := result[key]; !ok {
			t.Errorf("expected %q key", key)
		}
	}
	// Omitted nil fields
	for _, key := range []string{"components", "security", "tags", "externalDocs"} {
		if _, ok := result[key]; ok {
			t.Errorf("expected %q to be omitted", key)
		}
	}
}

func TestOpenAPI_MarshalJSON_Minimal(t *testing.T) {
	// Arrange
	info := NewInfo("API", "", "", "1.0", nil, nil)
	doc := NewOpenAPI("3.0.0", info)

	// Act
	got, err := json.Marshal(doc)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if _, ok := result["openapi"]; !ok {
		t.Error("expected 'openapi' key")
	}
	if _, ok := result["info"]; !ok {
		t.Error("expected 'info' key")
	}
}

func TestOpenAPI_MarshalYAML_RoundTrip(t *testing.T) {
	// Arrange
	info := NewInfo("Pet Store", "", "", "1.0.0", nil, nil)
	doc := NewOpenAPI("3.0.3", info)

	// Act
	node, err := doc.MarshalYAML()
	if err != nil {
		t.Fatal(err)
	}
	yamlDoc := &yaml.Node{Kind: yaml.DocumentNode, Content: []*yaml.Node{node.(*yaml.Node)}}
	out, err := yaml.Marshal(yamlDoc)
	if err != nil {
		t.Fatal(err)
	}

	// Assert
	var result map[string]interface{}
	if err := yaml.Unmarshal(out, &result); err != nil {
		t.Fatal(err)
	}
	if result["openapi"] != "3.0.3" {
		t.Errorf("openapi = %v, want %q", result["openapi"], "3.0.3")
	}
	info2, ok := result["info"].(map[string]interface{})
	if !ok {
		t.Fatal("info should be a map")
	}
	if info2["title"] != "Pet Store" {
		t.Errorf("info.title = %v, want %q", info2["title"], "Pet Store")
	}
}

func TestOpenAPI_MarshalJSON_WithExtensions(t *testing.T) {
	// Arrange
	info := NewInfo("API", "", "", "1.0", nil, nil)
	doc := NewOpenAPI("3.0.0", info)
	doc.VendorExtensions = map[string]interface{}{"x-generator": "openapi-parser"}

	// Act
	got, err := json.Marshal(doc)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if _, ok := result["x-generator"]; !ok {
		t.Error("expected 'x-generator' extension key")
	}
}
