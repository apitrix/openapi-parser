package openapi30

import (
	"errors"
	"testing"
)

func TestHeader_SetDescription_WithoutHook(t *testing.T) {
	h := NewHeader("old", false, false, false, "", nil, false, nil, nil, nil, nil)
	err := h.SetDescription("new")
	if err != nil {
		t.Fatalf("SetDescription without hook should succeed, got %v", err)
	}
	if h.Description() != "new" {
		t.Errorf("Description() = %q, want %q", h.Description(), "new")
	}
}

func TestHeader_SetDescription_WithHook_Rejects(t *testing.T) {
	h := NewHeader("old", false, false, false, "", nil, false, nil, nil, nil, nil)
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

func TestHeader_SetRequired_WithoutHook(t *testing.T) {
	h := NewHeader("", false, false, false, "", nil, false, nil, nil, nil, nil)
	err := h.SetRequired(true)
	if err != nil {
		t.Fatalf("SetRequired without hook should succeed, got %v", err)
	}
	if !h.Required() {
		t.Errorf("Required() = false, want true")
	}
}

func TestHeader_SetSchema_WithoutHook(t *testing.T) {
	h := NewHeader("", false, false, false, "", nil, false, nil, nil, nil, nil)
	ref := NewRefSchema("#/components/schemas/string")
	err := h.SetSchema(ref)
	if err != nil {
		t.Fatalf("SetSchema without hook should succeed, got %v", err)
	}
	if h.Schema() != ref {
		t.Errorf("Schema() = %v, want %v", h.Schema(), ref)
	}
}

func TestHeader_SetStyle_WithoutHook(t *testing.T) {
	h := NewHeader("", false, false, false, "old", nil, false, nil, nil, nil, nil)
	err := h.SetStyle("new")
	if err != nil {
		t.Fatalf("SetStyle without hook should succeed, got %v", err)
	}
	if h.Style() != "new" {
		t.Errorf("Style() = %q, want %q", h.Style(), "new")
	}
}
