package openapi30

import (
	"encoding/json"
	"testing"
)

func TestExampleRef_MarshalJSON_Ref(t *testing.T) {
	// Arrange
	ref := NewExampleRef("#/components/examples/Dog")

	// Act
	got, err := json.Marshal(ref)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"$ref":"#/components/examples/Dog"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestExampleRef_MarshalJSON_InlineValue(t *testing.T) {
	// Arrange
	ex := NewExample("Dog", "", nil, "")
	ref := &ExampleRef{Value: ex}

	// Act
	got, err := json.Marshal(ref)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"summary":"Dog"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
