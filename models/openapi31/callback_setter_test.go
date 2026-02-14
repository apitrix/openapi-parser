package openapi31

import (
	"errors"
	"testing"
)

func TestCallback_SetPaths_WithoutHook(t *testing.T) {
	c := NewCallback(nil)
	paths := map[string]*PathItem{"{$request.url}": NewPathItem()}
	err := c.SetPaths(paths)
	if err != nil {
		t.Fatalf("SetPaths without hook should succeed, got %v", err)
	}
	if c.Paths()["{$request.url}"] != paths["{$request.url}"] {
		t.Errorf("Paths() = %v, want %v", c.Paths(), paths)
	}
}

func TestCallback_SetPaths_WithHook_Rejects(t *testing.T) {
	c := NewCallback(nil)
	rejectErr := errors.New("rejected")
	c.Trix.OnSet("paths", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	paths := map[string]*PathItem{"{$request.url}": NewPathItem()}
	err := c.SetPaths(paths)
	if err != rejectErr {
		t.Errorf("SetPaths with rejecting hook should return error, got %v", err)
	}
	if c.Paths() != nil {
		t.Errorf("Paths should be unchanged after rejection")
	}
}
