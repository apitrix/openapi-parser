package openapi30

import (
	"encoding/json"
	"testing"
)

func TestResponseRef_MarshalJSON_Ref(t *testing.T) {
	// Arrange
	ref := NewResponseRef("#/components/responses/NotFound")

	// Act
	got, err := json.Marshal(ref)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"$ref":"#/components/responses/NotFound"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestResponseRef_MarshalJSON_InlineValue(t *testing.T) {
	// Arrange
	resp := NewResponse("Success", nil, nil, nil)
	ref := &ResponseRef{}
	ref.SetValue(resp)

	// Act
	got, err := json.Marshal(ref)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"description":"Success"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
