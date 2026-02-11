package openapi30

import (
	"encoding/json"
	"testing"
)

func TestExample_MarshalJSON_AllFields(t *testing.T) {
	// Arrange
	e := NewExample("A dog", "An example of a dog", map[string]interface{}{"name": "Fido"}, "")

	// Act
	got, err := json.Marshal(e)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if _, ok := result["summary"]; !ok {
		t.Error("expected 'summary' key")
	}
	if _, ok := result["description"]; !ok {
		t.Error("expected 'description' key")
	}
	if _, ok := result["value"]; !ok {
		t.Error("expected 'value' key")
	}
	if _, ok := result["externalValue"]; ok {
		t.Error("unexpected 'externalValue' key (should be omitted)")
	}
}

func TestExample_MarshalJSON_ExternalValue(t *testing.T) {
	// Arrange
	e := NewExample("", "", nil, "https://example.com/dog.json")

	// Act
	got, err := json.Marshal(e)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"externalValue":"https://example.com/dog.json"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
