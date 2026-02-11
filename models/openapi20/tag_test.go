package openapi20

import (
	"encoding/json"
	"testing"
)

func TestTag_MarshalJSON_AllFields(t *testing.T) {
	// Arrange
	ed := NewExternalDocs("Docs", "https://example.com")
	tag := NewTag("pets", "Pet operations", ed)

	// Act
	got, err := json.Marshal(tag)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := `{"name":"pets","description":"Pet operations","externalDocs":{"description":"Docs","url":"https://example.com"}}`
	if string(got) != expected {
		t.Errorf("got %s, want %s", got, expected)
	}
}

func TestTag_MarshalJSON_NameOnly(t *testing.T) {
	// Arrange
	tag := NewTag("pets", "", nil)

	// Act
	got, err := json.Marshal(tag)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := `{"name":"pets"}`
	if string(got) != expected {
		t.Errorf("got %s, want %s", got, expected)
	}
}
