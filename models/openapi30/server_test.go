package openapi30

import (
	"encoding/json"
	"testing"
)

func TestServer_MarshalJSON_AllFields(t *testing.T) {
	// Arrange
	vars := map[string]*ServerVariable{
		"port": NewServerVariable("8080", "Server port", nil),
	}
	s := NewServer("https://api.example.com", "Production", vars)

	// Act
	got, err := json.Marshal(s)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"url", "description", "variables"} {
		if _, ok := result[key]; !ok {
			t.Errorf("expected %q key", key)
		}
	}
}

func TestServer_MarshalJSON_URLOnly(t *testing.T) {
	// Arrange
	s := NewServer("https://api.example.com", "", nil)

	// Act
	got, err := json.Marshal(s)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"url":"https://api.example.com"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
