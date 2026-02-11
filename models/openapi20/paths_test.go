package openapi20

import (
	"encoding/json"
	"testing"
)

func TestPaths_MarshalJSON_WithItems(t *testing.T) {
	// Arrange
	getOp := NewOperation(nil, "List", "", nil, "", nil, nil, nil, nil, nil, false, nil)
	pi := NewPathItem("", getOp, nil, nil, nil, nil, nil, nil, nil)
	p := NewPaths(map[string]*PathItem{
		"/pets": pi,
	})

	// Act
	got, err := json.Marshal(p)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var m map[string]interface{}
	if err := json.Unmarshal(got, &m); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if m["/pets"] == nil {
		t.Error("/pets path should be present")
	}
}

func TestPaths_MarshalJSON_Sorted(t *testing.T) {
	// Arrange
	pi := NewPathItem("", nil, nil, nil, nil, nil, nil, nil, nil)
	p := NewPaths(map[string]*PathItem{
		"/z": pi,
		"/a": pi,
	})

	// Act
	got, err := json.Marshal(p)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// /a should appear before /z in sorted output
	s := string(got)
	aIdx := len(s)
	zIdx := 0
	for i := range s {
		if i+2 < len(s) && s[i:i+2] == "/a" {
			aIdx = i
		}
		if i+2 < len(s) && s[i:i+2] == "/z" {
			zIdx = i
		}
	}
	if aIdx > zIdx {
		t.Errorf("expected /a before /z in sorted output, got %s", s)
	}
}

func TestPaths_MarshalJSON_Empty(t *testing.T) {
	// Arrange
	p := NewPaths(nil)

	// Act
	got, err := json.Marshal(p)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(got) != "{}" {
		t.Errorf("got %s, want {}", got)
	}
}
