package openapi20

import (
	"errors"
	"testing"
)

func TestHeader_SetDescription_WithoutHook(t *testing.T) {
	h := NewHeader(HeaderFields{Description: "old"})
	err := h.SetDescription("new")
	if err != nil {
		t.Fatalf("SetDescription without hook should succeed, got %v", err)
	}
	if h.Description() != "new" {
		t.Errorf("Description() = %q, want %q", h.Description(), "new")
	}
}

func TestHeader_SetDescription_WithHook_Rejects(t *testing.T) {
	h := NewHeader(HeaderFields{Description: "old"})
	rejectErr := errors.New("rejected")
	h.Trix.OnSet("description", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	err := h.SetDescription("new")
	if err != rejectErr {
		t.Errorf("SetDescription with rejecting hook should return error, got %v", err)
	}
	if h.Description() != "old" {
		t.Errorf("Description should be unchanged after rejection, got %q", h.Description())
	}
}

func TestHeader_SetDescription_WithHook_Passes(t *testing.T) {
	h := NewHeader(HeaderFields{Description: "old"})
	h.Trix.OnSet("description", func(field string, oldVal, newVal interface{}) error {
		return nil
	})
	err := h.SetDescription("new")
	if err != nil {
		t.Fatalf("SetDescription with passing hook should succeed, got %v", err)
	}
	if h.Description() != "new" {
		t.Errorf("Description() = %q, want %q", h.Description(), "new")
	}
}

func TestHeader_SetType_WithoutHook(t *testing.T) {
	h := NewHeader(HeaderFields{Type: "old"})
	err := h.SetType("new")
	if err != nil {
		t.Fatalf("SetType without hook should succeed, got %v", err)
	}
	if h.Type() != "new" {
		t.Errorf("Type() = %q, want %q", h.Type(), "new")
	}
}

func TestHeader_SetFormat_WithoutHook(t *testing.T) {
	h := NewHeader(HeaderFields{Format: "old"})
	err := h.SetFormat("new")
	if err != nil {
		t.Fatalf("SetFormat without hook should succeed, got %v", err)
	}
	if h.Format() != "new" {
		t.Errorf("Format() = %q, want %q", h.Format(), "new")
	}
}

func TestHeader_SetPattern_WithoutHook(t *testing.T) {
	h := NewHeader(HeaderFields{Pattern: "old"})
	err := h.SetPattern("new")
	if err != nil {
		t.Fatalf("SetPattern without hook should succeed, got %v", err)
	}
	if h.Pattern() != "new" {
		t.Errorf("Pattern() = %q, want %q", h.Pattern(), "new")
	}
}
