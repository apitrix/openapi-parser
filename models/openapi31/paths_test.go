package openapi31

import (
	"encoding/json"
	"testing"
)

func TestPaths_MarshalJSON_WithItems(t *testing.T) {
	pi := NewPathItem()
	pi.SetProperty("summary", "Pet endpoint")
	paths := NewPaths(map[string]*PathItem{"/pets": pi})
	got, err := json.Marshal(paths)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if _, ok := result["/pets"]; !ok {
		t.Error("expected '/pets' key")
	}
}

func TestPaths_MarshalJSON_SortedKeys(t *testing.T) {
	paths := NewPaths(map[string]*PathItem{
		"/z-path": NewPathItem(),
		"/a-path": NewPathItem(),
	})
	got, err := json.Marshal(paths)
	if err != nil {
		t.Fatal(err)
	}
	// Keys should appear in sorted order
	want := `{"/a-path":{},"/z-path":{}}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestPaths_MarshalJSON_Empty(t *testing.T) {
	paths := NewPaths(nil)
	got, err := json.Marshal(paths)
	if err != nil {
		t.Fatal(err)
	}
	want := `{}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
