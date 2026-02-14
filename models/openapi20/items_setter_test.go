package openapi20

import (
	"errors"
	"testing"
)

func TestItems_SetType_WithoutHook(t *testing.T) {
	it := NewItems(ItemsFields{Type: "old"})
	err := it.SetType("new")
	if err != nil {
		t.Fatalf("SetType without hook should succeed, got %v", err)
	}
	if it.Type() != "new" {
		t.Errorf("Type() = %q, want %q", it.Type(), "new")
	}
}

func TestItems_SetType_WithHook_Rejects(t *testing.T) {
	it := NewItems(ItemsFields{Type: "old"})
	rejectErr := errors.New("rejected")
	it.Trix.OnSet("type", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	err := it.SetType("new")
	if err != rejectErr {
		t.Errorf("SetType with rejecting hook should return error, got %v", err)
	}
	if it.Type() != "old" {
		t.Errorf("Type should be unchanged after rejection, got %q", it.Type())
	}
}

func TestItems_SetType_WithHook_Passes(t *testing.T) {
	it := NewItems(ItemsFields{Type: "old"})
	it.Trix.OnSet("type", func(field string, oldVal, newVal interface{}) error {
		return nil
	})
	err := it.SetType("new")
	if err != nil {
		t.Fatalf("SetType with passing hook should succeed, got %v", err)
	}
	if it.Type() != "new" {
		t.Errorf("Type() = %q, want %q", it.Type(), "new")
	}
}

func TestItems_SetFormat_WithoutHook(t *testing.T) {
	it := NewItems(ItemsFields{Format: "old"})
	err := it.SetFormat("new")
	if err != nil {
		t.Fatalf("SetFormat without hook should succeed, got %v", err)
	}
	if it.Format() != "new" {
		t.Errorf("Format() = %q, want %q", it.Format(), "new")
	}
}

func TestItems_SetCollectionFormat_WithoutHook(t *testing.T) {
	it := NewItems(ItemsFields{CollectionFormat: "old"})
	err := it.SetCollectionFormat("new")
	if err != nil {
		t.Fatalf("SetCollectionFormat without hook should succeed, got %v", err)
	}
	if it.CollectionFormat() != "new" {
		t.Errorf("CollectionFormat() = %q, want %q", it.CollectionFormat(), "new")
	}
}

func TestItems_SetPattern_WithoutHook(t *testing.T) {
	it := NewItems(ItemsFields{Pattern: "old"})
	err := it.SetPattern("new")
	if err != nil {
		t.Fatalf("SetPattern without hook should succeed, got %v", err)
	}
	if it.Pattern() != "new" {
		t.Errorf("Pattern() = %q, want %q", it.Pattern(), "new")
	}
}

func TestItems_SetUniqueItems_WithoutHook(t *testing.T) {
	it := NewItems(ItemsFields{})
	err := it.SetUniqueItems(true)
	if err != nil {
		t.Fatalf("SetUniqueItems without hook should succeed, got %v", err)
	}
	if !it.UniqueItems() {
		t.Errorf("UniqueItems() = false, want true")
	}
}
