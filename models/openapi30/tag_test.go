package openapi30

import (
	"encoding/json"
	"testing"
)

func TestTag_MarshalJSON_AllFields(t *testing.T) {
	// Arrange
	ed := NewExternalDocumentation("https://docs.example.com", "Tag docs")
	tag := NewTag("pets", "Pet operations", ed)

	// Act
	got, err := json.Marshal(tag)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if _, ok := result["name"]; !ok {
		t.Error("expected 'name' key")
	}
	if _, ok := result["description"]; !ok {
		t.Error("expected 'description' key")
	}
	if _, ok := result["externalDocs"]; !ok {
		t.Error("expected 'externalDocs' key")
	}
}

func TestTag_MarshalJSON_NameOnly(t *testing.T) {
	// Arrange
	tag := NewTag("pets", "", nil)

	// Act
	got, err := json.Marshal(tag)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"name":"pets"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
