package openapi31

import (
	"errors"
	"testing"
)

func TestServerVariable_SetDefault_WithoutHook(t *testing.T) {
	sv := NewServerVariable(nil, "old", "")
	err := sv.SetDefault("new")
	if err != nil {
		t.Fatalf("SetDefault without hook should succeed, got %v", err)
	}
	if sv.Default() != "new" {
		t.Errorf("Default() = %q, want %q", sv.Default(), "new")
	}
}

func TestServerVariable_SetDefault_WithHook_Rejects(t *testing.T) {
	sv := NewServerVariable(nil, "old", "")
	rejectErr := errors.New("rejected")
	sv.Trix.OnSet("default", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	err := sv.SetDefault("new")
	if err != rejectErr {
		t.Errorf("SetDefault with rejecting hook should return error, got %v", err)
	}
	if sv.Default() != "old" {
		t.Errorf("Default should be unchanged after rejection, got %q", sv.Default())
	}
}

func TestServerVariable_SetDescription_WithoutHook(t *testing.T) {
	sv := NewServerVariable(nil, "", "old")
	err := sv.SetDescription("new")
	if err != nil {
		t.Fatalf("SetDescription without hook should succeed, got %v", err)
	}
	if sv.Description() != "new" {
		t.Errorf("Description() = %q, want %q", sv.Description(), "new")
	}
}

func TestServerVariable_SetEnum_WithoutHook(t *testing.T) {
	sv := NewServerVariable(nil, "", "")
	err := sv.SetEnum([]string{"v1", "v2"})
	if err != nil {
		t.Fatalf("SetEnum without hook should succeed, got %v", err)
	}
	if len(sv.Enum()) != 2 || sv.Enum()[0] != "v1" {
		t.Errorf("Enum() = %v, want [v1 v2]", sv.Enum())
	}
}
