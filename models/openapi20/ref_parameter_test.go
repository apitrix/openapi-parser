package openapi20

import (
	"encoding/json"
	"testing"
)

func TestParameterRef_MarshalJSON_Ref(t *testing.T) {
	// Arrange
	ref := NewParameterRef("#/parameters/PageSize")

	// Act
	got, err := json.Marshal(ref)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := `{"$ref":"#/parameters/PageSize"}`
	if string(got) != expected {
		t.Errorf("got %s, want %s", got, expected)
	}
}

func TestParameterRef_MarshalJSON_InlineValue(t *testing.T) {
	// Arrange
	p := NewParameter(ParameterFields{Name: "limit", In: "query", Type: "integer"})
	ref := &ParameterRef{}
	ref.SetValue(p)

	// Act
	got, err := json.Marshal(ref)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := `{"name":"limit","in":"query","type":"integer"}`
	if string(got) != expected {
		t.Errorf("got %s, want %s", got, expected)
	}
}
