package openapi30

import (
	"encoding/json"
	"testing"
)

func TestInfo_MarshalJSON_AllFields(t *testing.T) {
	// Arrange
	c := NewContact("John", "", "")
	l := NewLicense("MIT", "")
	info := NewInfo("Pet Store", "A pet store API", "https://tos.example.com", "1.0.0", c, l)

	// Act
	got, err := json.Marshal(info)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"title", "description", "termsOfService", "contact", "license", "version"} {
		if _, ok := result[key]; !ok {
			t.Errorf("expected %q key", key)
		}
	}
}

func TestInfo_MarshalJSON_MinimalFields(t *testing.T) {
	// Arrange
	info := NewInfo("Pet Store", "", "", "1.0.0", nil, nil)

	// Act
	got, err := json.Marshal(info)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"title":"Pet Store","version":"1.0.0"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestInfo_MarshalJSON_WithExtensions(t *testing.T) {
	// Arrange
	info := NewInfo("API", "", "", "1.0", nil, nil)
	info.VendorExtensions = map[string]interface{}{"x-logo": "https://logo.png"}

	// Act
	got, err := json.Marshal(info)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"title":"API","version":"1.0","x-logo":"https://logo.png"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
