package openapi20

import (
	"encoding/json"
	"testing"
)

func TestExternalDocs_MarshalJSON_AllFields(t *testing.T) {
	// Arrange
	ed := NewExternalDocs("More info", "https://example.com/docs")

	// Act
	got, err := json.Marshal(ed)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := `{"description":"More info","url":"https://example.com/docs"}`
	if string(got) != expected {
		t.Errorf("got %s, want %s", got, expected)
	}
}

func TestExternalDocs_MarshalJSON_URLOnly(t *testing.T) {
	// Arrange
	ed := NewExternalDocs("", "https://example.com")

	// Act
	got, err := json.Marshal(ed)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := `{"url":"https://example.com"}`
	if string(got) != expected {
		t.Errorf("got %s, want %s", got, expected)
	}
}
