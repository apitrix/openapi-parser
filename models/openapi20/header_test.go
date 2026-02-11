package openapi20

import (
	"encoding/json"
	"testing"
)

func TestHeader_MarshalJSON_AllFields(t *testing.T) {
	// Arrange
	h := NewHeader(HeaderFields{
		Description: "Rate limit",
		Type:        "integer",
		Format:      "int32",
	})

	// Act
	got, err := json.Marshal(h)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := `{"description":"Rate limit","type":"integer","format":"int32"}`
	if string(got) != expected {
		t.Errorf("got %s, want %s", got, expected)
	}
}

func TestHeader_MarshalJSON_TypeOnly(t *testing.T) {
	// Arrange
	h := NewHeader(HeaderFields{Type: "string"})

	// Act
	got, err := json.Marshal(h)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := `{"type":"string"}`
	if string(got) != expected {
		t.Errorf("got %s, want %s", got, expected)
	}
}
