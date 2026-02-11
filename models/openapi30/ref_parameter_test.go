package openapi30

import (
	"encoding/json"
	"testing"
)

func TestParameterRef_MarshalJSON_Ref(t *testing.T) {
	// Arrange
	ref := NewParameterRef("#/components/parameters/PageSize")

	// Act
	got, err := json.Marshal(ref)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"$ref":"#/components/parameters/PageSize"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestParameterRef_MarshalJSON_InlineValue(t *testing.T) {
	// Arrange
	p := NewParameter("limit", "query", "Max items", false, false, false, "", nil, false, nil, nil, nil, nil)
	ref := &ParameterRef{Value: p}

	// Act
	got, err := json.Marshal(ref)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if _, ok := result["name"]; !ok {
		t.Error("expected 'name' key from inline parameter")
	}
	if _, ok := result["in"]; !ok {
		t.Error("expected 'in' key from inline parameter")
	}
}
