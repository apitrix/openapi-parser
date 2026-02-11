package openapi30

import (
	"encoding/json"
	"testing"
)

func TestRequestBody_MarshalJSON_AllFields(t *testing.T) {
	// Arrange
	content := map[string]*MediaType{
		"application/json": NewMediaType(nil, nil, nil, nil),
	}
	rb := NewRequestBody("A new pet", content, true)

	// Act
	got, err := json.Marshal(rb)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"description", "content", "required"} {
		if _, ok := result[key]; !ok {
			t.Errorf("expected %q key", key)
		}
	}
}

func TestRequestBody_MarshalJSON_OmitsFalseRequired(t *testing.T) {
	// Arrange
	rb := NewRequestBody("A body", nil, false)

	// Act
	got, err := json.Marshal(rb)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"description":"A body"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
