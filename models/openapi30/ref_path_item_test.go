package openapi30

import (
	"encoding/json"
	"testing"
)

func TestPathItemRef_MarshalJSON_Ref(t *testing.T) {
	// Arrange
	ref := NewPathItemRef("#/components/pathItems/Shared")

	// Act
	got, err := json.Marshal(ref)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"$ref":"#/components/pathItems/Shared"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestPathItemRef_MarshalJSON_InlineValue(t *testing.T) {
	// Arrange
	pi := NewPathItem("", "Shared path item", "", nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	ref := &PathItemRef{Value: pi}

	// Act
	got, err := json.Marshal(ref)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"summary":"Shared path item"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
