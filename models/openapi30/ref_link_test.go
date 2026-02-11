package openapi30

import (
	"encoding/json"
	"testing"
)

func TestLinkRef_MarshalJSON_Ref(t *testing.T) {
	// Arrange
	ref := NewLinkRef("#/components/links/GetUser")

	// Act
	got, err := json.Marshal(ref)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"$ref":"#/components/links/GetUser"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestLinkRef_MarshalJSON_InlineValue(t *testing.T) {
	// Arrange
	l := NewLink("", "getUser", nil, nil, "Get a user", nil)
	ref := &LinkRef{Value: l}

	// Act
	got, err := json.Marshal(ref)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"operationId":"getUser","description":"Get a user"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
