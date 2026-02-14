package openapi31

import (
	"errors"
	"testing"
)

func TestPaths_SetItems_WithoutHook(t *testing.T) {
	p := NewPaths(nil)
	items := map[string]*PathItem{"/pets": NewPathItem()}
	err := p.SetItems(items)
	if err != nil {
		t.Fatalf("SetItems without hook should succeed, got %v", err)
	}
	if p.Items()["/pets"] != items["/pets"] {
		t.Errorf("Items() = %v, want %v", p.Items(), items)
	}
}

func TestPaths_SetItems_WithHook_Rejects(t *testing.T) {
	p := NewPaths(nil)
	rejectErr := errors.New("rejected")
	p.Trix.OnSet("items", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	items := map[string]*PathItem{"/pets": NewPathItem()}
	err := p.SetItems(items)
	if err != rejectErr {
		t.Errorf("SetItems with rejecting hook should return error, got %v", err)
	}
	if p.Items() != nil {
		t.Errorf("Items should be unchanged after rejection")
	}
}
