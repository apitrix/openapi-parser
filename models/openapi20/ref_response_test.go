package openapi20

import (
	"encoding/json"
	"testing"
)

func TestResponseRef_MarshalJSON_Ref(t *testing.T) {
	// Arrange
	ref := NewResponseRef("#/responses/NotFound")

	// Act
	got, err := json.Marshal(ref)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := `{"$ref":"#/responses/NotFound"}`
	if string(got) != expected {
		t.Errorf("got %s, want %s", got, expected)
	}
}

func TestResponseRef_MarshalJSON_InlineValue(t *testing.T) {
	// Arrange
	ref := &ResponseRef{Value: NewResponse("OK", nil, nil, nil)}

	// Act
	got, err := json.Marshal(ref)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := `{"description":"OK"}`
	if string(got) != expected {
		t.Errorf("got %s, want %s", got, expected)
	}
}
