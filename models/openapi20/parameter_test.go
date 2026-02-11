package openapi20

import (
	"encoding/json"
	"testing"
)

func TestParameter_MarshalJSON_BodyParam(t *testing.T) {
	// Arrange
	schema := &SchemaRef{Value: NewSchema(SchemaFields{Type: "object"})}
	p := NewParameter(ParameterFields{
		Name:        "body",
		In:          "body",
		Description: "Request body",
		Required:    true,
		Schema:      schema,
	})

	// Act
	got, err := json.Marshal(p)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var m map[string]interface{}
	if err := json.Unmarshal(got, &m); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if m["name"] != "body" {
		t.Errorf("name = %v, want body", m["name"])
	}
	if m["in"] != "body" {
		t.Errorf("in = %v, want body", m["in"])
	}
	if m["required"] != true {
		t.Errorf("required = %v, want true", m["required"])
	}
	if m["schema"] == nil {
		t.Error("schema should not be nil")
	}
}

func TestParameter_MarshalJSON_QueryParam(t *testing.T) {
	// Arrange
	p := NewParameter(ParameterFields{
		Name: "limit",
		In:   "query",
		Type: "integer",
	})

	// Act
	got, err := json.Marshal(p)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := `{"name":"limit","in":"query","type":"integer"}`
	if string(got) != expected {
		t.Errorf("got %s, want %s", got, expected)
	}
}

func TestParameter_MarshalJSON_WithCollectionFormat(t *testing.T) {
	// Arrange
	p := NewParameter(ParameterFields{
		Name:             "tags",
		In:               "query",
		Type:             "array",
		CollectionFormat: "csv",
		Items:            NewItems(ItemsFields{Type: "string"}),
	})

	// Act
	got, err := json.Marshal(p)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var m map[string]interface{}
	if err := json.Unmarshal(got, &m); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if m["collectionFormat"] != "csv" {
		t.Errorf("collectionFormat = %v, want csv", m["collectionFormat"])
	}
	if m["items"] == nil {
		t.Error("items should not be nil")
	}
}
