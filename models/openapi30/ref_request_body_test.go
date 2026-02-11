package openapi30

import (
	"encoding/json"
	"testing"
)

func TestRequestBodyRef_MarshalJSON_Ref(t *testing.T) {
	// Arrange
	ref := NewRequestBodyRef("#/components/requestBodies/CreatePet")

	// Act
	got, err := json.Marshal(ref)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"$ref":"#/components/requestBodies/CreatePet"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestRequestBodyRef_MarshalJSON_InlineValue(t *testing.T) {
	// Arrange
	rb := NewRequestBody("A request body", nil, true)
	ref := &RequestBodyRef{Value: rb}

	// Act
	got, err := json.Marshal(ref)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"description":"A request body","required":true}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
