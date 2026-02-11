package openapi30

import (
	"encoding/json"
	"testing"
)

func TestExternalDocumentation_MarshalJSON_AllFields(t *testing.T) {
	// Arrange
	ed := NewExternalDocumentation("https://docs.example.com", "More info")

	// Act
	got, err := json.Marshal(ed)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"description":"More info","url":"https://docs.example.com"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestExternalDocumentation_MarshalJSON_OmitsEmpty(t *testing.T) {
	// Arrange
	ed := NewExternalDocumentation("https://docs.example.com", "")

	// Act
	got, err := json.Marshal(ed)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"url":"https://docs.example.com"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
