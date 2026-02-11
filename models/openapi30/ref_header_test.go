package openapi30

import (
	"encoding/json"
	"testing"
)

func TestHeaderRef_MarshalJSON_Ref(t *testing.T) {
	// Arrange
	ref := NewHeaderRef("#/components/headers/X-Rate-Limit")

	// Act
	got, err := json.Marshal(ref)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"$ref":"#/components/headers/X-Rate-Limit"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestHeaderRef_MarshalJSON_InlineValue(t *testing.T) {
	// Arrange
	h := NewHeader("Rate limit", false, false, false, "", nil, false, nil, nil, nil, nil)
	ref := &HeaderRef{Value: h}

	// Act
	got, err := json.Marshal(ref)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"description":"Rate limit"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
