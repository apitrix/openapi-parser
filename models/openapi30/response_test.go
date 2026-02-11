package openapi30

import (
	"encoding/json"
	"testing"
)

func TestResponse_MarshalJSON_AllFields(t *testing.T) {
	// Arrange
	content := map[string]*MediaType{
		"application/json": NewMediaType(nil, nil, nil, nil),
	}
	r := NewResponse("OK", nil, content, nil)

	// Act
	got, err := json.Marshal(r)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if _, ok := result["description"]; !ok {
		t.Error("expected 'description' key")
	}
	if _, ok := result["content"]; !ok {
		t.Error("expected 'content' key")
	}
}

func TestResponse_MarshalJSON_DescriptionOnly(t *testing.T) {
	// Arrange
	r := NewResponse("Not Found", nil, nil, nil)

	// Act
	got, err := json.Marshal(r)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"description":"Not Found"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
