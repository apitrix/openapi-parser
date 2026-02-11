package openapi30

import (
	"encoding/json"
	"testing"
)

func TestCallback_MarshalJSON_WithPaths(t *testing.T) {
	// Arrange
	pi := NewPathItem("", "", "", nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	cb := NewCallback(map[string]*PathItem{
		"{$request.body#/callbackUrl}": pi,
	})

	// Act
	got, err := json.Marshal(cb)

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

func TestCallback_MarshalJSON_Empty(t *testing.T) {
	// Arrange
	cb := NewCallback(nil)

	// Act
	got, err := json.Marshal(cb)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
