package openapi31

import (
	"errors"
	"testing"
)

func TestExample_SetSummary_WithoutHook(t *testing.T) {
	e := NewExample("old", "", nil, "")
	err := e.SetSummary("new")
	if err != nil {
		t.Fatalf("SetSummary without hook should succeed, got %v", err)
	}
	if e.Summary() != "new" {
		t.Errorf("Summary() = %q, want %q", e.Summary(), "new")
	}
}

func TestExample_SetSummary_WithHook_Rejects(t *testing.T) {
	e := NewExample("old", "", nil, "")
	rejectErr := errors.New("rejected")
	e.Trix.OnSet("summary", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	err := e.SetSummary("new")
	if err != rejectErr {
		t.Errorf("SetSummary with rejecting hook should return error, got %v", err)
	}
	if e.Summary() != "old" {
		t.Errorf("Summary should be unchanged after rejection, got %q", e.Summary())
	}
}

func TestExample_SetDescription_WithoutHook(t *testing.T) {
	e := NewExample("", "old", nil, "")
	err := e.SetDescription("new")
	if err != nil {
		t.Fatalf("SetDescription without hook should succeed, got %v", err)
	}
	if e.Description() != "new" {
		t.Errorf("Description() = %q, want %q", e.Description(), "new")
	}
}

func TestExample_SetValue_WithoutHook(t *testing.T) {
	e := NewExample("", "", nil, "")
	err := e.SetValue(map[string]string{"name": "Puma"})
	if err != nil {
		t.Fatalf("SetValue without hook should succeed, got %v", err)
	}
	if e.Value() == nil {
		t.Errorf("Value() = nil, want non-nil")
	}
}

func TestExample_SetExternalValue_WithoutHook(t *testing.T) {
	e := NewExample("", "", nil, "")
	err := e.SetExternalValue("http://example.com/sample.json")
	if err != nil {
		t.Fatalf("SetExternalValue without hook should succeed, got %v", err)
	}
	if e.ExternalValue() != "http://example.com/sample.json" {
		t.Errorf("ExternalValue() = %q, want http://example.com/sample.json", e.ExternalValue())
	}
}
