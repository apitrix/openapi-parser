package openapi30

import (
	"encoding/json"
	"testing"
)

func TestPaths_MarshalJSON_SortedKeys(t *testing.T) {
	// Arrange
	paths := NewPaths(map[string]*PathItem{
		"/pets":       NewPathItem("", "", "", nil, nil, nil, nil, nil, nil, nil, nil, nil, nil),
		"/animals":    NewPathItem("", "", "", nil, nil, nil, nil, nil, nil, nil, nil, nil, nil),
		"/users/{id}": NewPathItem("", "", "", nil, nil, nil, nil, nil, nil, nil, nil, nil, nil),
	})

	// Act
	got, err := json.Marshal(paths)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	// Keys should be sorted alphabetically
	want := `{"/animals":{},"/pets":{},"/users/{id}":{}}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestPaths_MarshalJSON_Empty(t *testing.T) {
	// Arrange
	paths := NewPaths(nil)

	// Act
	got, err := json.Marshal(paths)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestPaths_MarshalJSON_WithExtensions(t *testing.T) {
	// Arrange
	paths := NewPaths(map[string]*PathItem{
		"/pets": NewPathItem("", "", "", nil, nil, nil, nil, nil, nil, nil, nil, nil, nil),
	})
	paths.VendorExtensions = map[string]interface{}{"x-custom": true}

	// Act
	got, err := json.Marshal(paths)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"/pets":{},"x-custom":true}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
