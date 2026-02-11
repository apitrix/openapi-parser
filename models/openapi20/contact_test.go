package openapi20

import (
	"encoding/json"
	"testing"
)

func TestContact_MarshalJSON_AllFields(t *testing.T) {
	// Arrange
	c := NewContact("John", "https://example.com", "john@example.com")

	// Act
	got, err := json.Marshal(c)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := `{"name":"John","url":"https://example.com","email":"john@example.com"}`
	if string(got) != expected {
		t.Errorf("got %s, want %s", got, expected)
	}
}

func TestContact_MarshalJSON_Empty(t *testing.T) {
	// Arrange
	c := NewContact("", "", "")

	// Act
	got, err := json.Marshal(c)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(got) != "{}" {
		t.Errorf("got %s, want {}", got)
	}
}

func TestContact_MarshalJSON_WithExtensions(t *testing.T) {
	// Arrange
	c := NewContact("John", "", "")
	c.VendorExtensions = map[string]interface{}{"x-team": "backend"}

	// Act
	got, err := json.Marshal(c)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := `{"name":"John","x-team":"backend"}`
	if string(got) != expected {
		t.Errorf("got %s, want %s", got, expected)
	}
}
