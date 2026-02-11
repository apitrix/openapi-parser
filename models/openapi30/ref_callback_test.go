package openapi30

import (
	"encoding/json"
	"testing"
)

func TestCallbackRef_MarshalJSON_Ref(t *testing.T) {
	// Arrange
	ref := NewCallbackRef("#/components/callbacks/onPetCreate")

	// Act
	got, err := json.Marshal(ref)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"$ref":"#/components/callbacks/onPetCreate"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestCallbackRef_MarshalJSON_InlineValue(t *testing.T) {
	// Arrange
	pi := NewPathItem("", "Webhook", "", nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	cb := NewCallback(map[string]*PathItem{"{$request.body#/callbackUrl}": pi})
	ref := &CallbackRef{Value: cb}

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
	if _, ok := result["{$request.body#/callbackUrl}"]; !ok {
		t.Error("expected callback expression key")
	}
}
