package openapi20

import (
	"encoding/json"
	"testing"
)

func TestLicense_MarshalJSON_AllFields(t *testing.T) {
	// Arrange
	l := NewLicense("MIT", "https://opensource.org/licenses/MIT")

	// Act
	got, err := json.Marshal(l)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := `{"name":"MIT","url":"https://opensource.org/licenses/MIT"}`
	if string(got) != expected {
		t.Errorf("got %s, want %s", got, expected)
	}
}

func TestLicense_MarshalJSON_NameOnly(t *testing.T) {
	// Arrange
	l := NewLicense("Apache-2.0", "")

	// Act
	got, err := json.Marshal(l)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := `{"name":"Apache-2.0"}`
	if string(got) != expected {
		t.Errorf("got %s, want %s", got, expected)
	}
}
