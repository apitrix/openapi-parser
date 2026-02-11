package openapi20

import (
	"encoding/json"
	"testing"
)

func TestItems_MarshalJSON_Simple(t *testing.T) {
	// Arrange
	it := NewItems(ItemsFields{Type: "string"})

	// Act
	got, err := json.Marshal(it)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := `{"type":"string"}`
	if string(got) != expected {
		t.Errorf("got %s, want %s", got, expected)
	}
}

func TestItems_MarshalJSON_Nested(t *testing.T) {
	// Arrange
	inner := NewItems(ItemsFields{Type: "integer"})
	it := NewItems(ItemsFields{Type: "array", Items: inner})

	// Act
	got, err := json.Marshal(it)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := `{"type":"array","items":{"type":"integer"}}`
	if string(got) != expected {
		t.Errorf("got %s, want %s", got, expected)
	}
}

func TestItems_MarshalJSON_WithConstraints(t *testing.T) {
	// Arrange
	max := float64(100)
	it := NewItems(ItemsFields{
		Type:    "integer",
		Maximum: &max,
	})

	// Act
	got, err := json.Marshal(it)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var m map[string]interface{}
	if err := json.Unmarshal(got, &m); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if m["maximum"] != float64(100) {
		t.Errorf("maximum = %v, want 100", m["maximum"])
	}
}
