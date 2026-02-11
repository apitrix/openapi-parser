package openapi20

import (
	"encoding/json"
	"testing"
)

func TestSwagger_MarshalJSON_Minimal(t *testing.T) {
	// Arrange
	doc := &Swagger{}
	doc.SetProperty("swagger", "2.0")
	doc.SetProperty("info", NewInfo("Pet Store", "", "", "1.0", nil, nil))
	doc.SetProperty("paths", NewPaths(nil))

	// Act
	got, err := json.Marshal(doc)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var m map[string]interface{}
	if err := json.Unmarshal(got, &m); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if m["swagger"] != "2.0" {
		t.Errorf("swagger = %v, want 2.0", m["swagger"])
	}
	if m["info"] == nil {
		t.Error("info should not be nil")
	}
}

func TestSwagger_MarshalJSON_WithHostAndBasePath(t *testing.T) {
	// Arrange
	doc := &Swagger{}
	doc.SetProperty("swagger", "2.0")
	doc.SetProperty("info", NewInfo("API", "", "", "1.0", nil, nil))
	doc.SetProperty("host", "api.example.com")
	doc.SetProperty("basePath", "/v1")
	doc.SetProperty("schemes", []string{"https"})

	// Act
	got, err := json.Marshal(doc)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var m map[string]interface{}
	if err := json.Unmarshal(got, &m); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if m["host"] != "api.example.com" {
		t.Errorf("host = %v, want api.example.com", m["host"])
	}
	if m["basePath"] != "/v1" {
		t.Errorf("basePath = %v, want /v1", m["basePath"])
	}
}

func TestSwagger_MarshalJSON_WithDefinitions(t *testing.T) {
	// Arrange
	doc := &Swagger{}
	doc.SetProperty("swagger", "2.0")
	doc.SetProperty("info", NewInfo("API", "", "", "1.0", nil, nil))
	doc.SetProperty("definitions", map[string]*SchemaRef{
		"Pet": {Value: NewSchema(SchemaFields{Type: "object"})},
	})

	// Act
	got, err := json.Marshal(doc)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var m map[string]interface{}
	if err := json.Unmarshal(got, &m); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	defs, ok := m["definitions"].(map[string]interface{})
	if !ok {
		t.Fatal("definitions should be an object")
	}
	if defs["Pet"] == nil {
		t.Error("Pet definition should be present")
	}
}

func TestSwagger_MarshalJSON_WithExtensions(t *testing.T) {
	// Arrange
	doc := &Swagger{}
	doc.SetProperty("swagger", "2.0")
	doc.SetProperty("info", NewInfo("API", "", "", "1.0", nil, nil))
	doc.VendorExtensions = map[string]interface{}{"x-generator": "openapi-parser"}

	// Act
	got, err := json.Marshal(doc)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var m map[string]interface{}
	if err := json.Unmarshal(got, &m); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if m["x-generator"] != "openapi-parser" {
		t.Errorf("x-generator = %v, want openapi-parser", m["x-generator"])
	}
}
