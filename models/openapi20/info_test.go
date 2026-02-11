package openapi20

import (
	"encoding/json"
	"testing"
)

func TestInfo_MarshalJSON_AllFields(t *testing.T) {
	// Arrange
	c := NewContact("John", "", "")
	l := NewLicense("MIT", "")
	info := NewInfo("Pet Store", "A pet store API", "https://example.com/tos", "1.0.0", c, l)

	// Act
	got, err := json.Marshal(info)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var m map[string]interface{}
	if err := json.Unmarshal(got, &m); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if m["title"] != "Pet Store" {
		t.Errorf("title = %v, want Pet Store", m["title"])
	}
	if m["version"] != "1.0.0" {
		t.Errorf("version = %v, want 1.0.0", m["version"])
	}
	if m["termsOfService"] != "https://example.com/tos" {
		t.Errorf("termsOfService = %v, want https://example.com/tos", m["termsOfService"])
	}
}

func TestInfo_MarshalJSON_Minimal(t *testing.T) {
	// Arrange
	info := NewInfo("API", "", "", "1.0", nil, nil)

	// Act
	got, err := json.Marshal(info)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := `{"title":"API","version":"1.0"}`
	if string(got) != expected {
		t.Errorf("got %s, want %s", got, expected)
	}
}

func TestInfo_MarshalJSON_WithExtensions(t *testing.T) {
	// Arrange
	info := NewInfo("API", "", "", "1.0", nil, nil)
	info.VendorExtensions = map[string]interface{}{"x-logo": "logo.png"}

	// Act
	got, err := json.Marshal(info)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := `{"title":"API","version":"1.0","x-logo":"logo.png"}`
	if string(got) != expected {
		t.Errorf("got %s, want %s", got, expected)
	}
}
