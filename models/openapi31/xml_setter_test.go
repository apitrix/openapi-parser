package openapi31

import (
	"errors"
	"testing"
)

func TestXML_SetName_WithoutHook(t *testing.T) {
	x := NewXML("old", "", "", false, false)
	err := x.SetName("new")
	if err != nil {
		t.Fatalf("SetName without hook should succeed, got %v", err)
	}
	if x.Name() != "new" {
		t.Errorf("Name() = %q, want %q", x.Name(), "new")
	}
}

func TestXML_SetName_WithHook_Rejects(t *testing.T) {
	x := NewXML("old", "", "", false, false)
	rejectErr := errors.New("rejected")
	x.Trix.OnSet("name", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	err := x.SetName("new")
	if err != rejectErr {
		t.Errorf("SetName with rejecting hook should return error, got %v", err)
	}
	if x.Name() != "old" {
		t.Errorf("Name should be unchanged after rejection, got %q", x.Name())
	}
}

func TestXML_SetNamespace_WithoutHook(t *testing.T) {
	x := NewXML("", "old", "", false, false)
	err := x.SetNamespace("new")
	if err != nil {
		t.Fatalf("SetNamespace without hook should succeed, got %v", err)
	}
	if x.Namespace() != "new" {
		t.Errorf("Namespace() = %q, want %q", x.Namespace(), "new")
	}
}

func TestXML_SetPrefix_WithoutHook(t *testing.T) {
	x := NewXML("", "", "old", false, false)
	err := x.SetPrefix("new")
	if err != nil {
		t.Fatalf("SetPrefix without hook should succeed, got %v", err)
	}
	if x.Prefix() != "new" {
		t.Errorf("Prefix() = %q, want %q", x.Prefix(), "new")
	}
}

func TestXML_SetAttribute_WithoutHook(t *testing.T) {
	x := NewXML("", "", "", false, false)
	err := x.SetAttribute(true)
	if err != nil {
		t.Fatalf("SetAttribute without hook should succeed, got %v", err)
	}
	if !x.Attribute() {
		t.Errorf("Attribute() = false, want true")
	}
}

func TestXML_SetWrapped_WithoutHook(t *testing.T) {
	x := NewXML("", "", "", false, false)
	err := x.SetWrapped(true)
	if err != nil {
		t.Fatalf("SetWrapped without hook should succeed, got %v", err)
	}
	if !x.Wrapped() {
		t.Errorf("Wrapped() = false, want true")
	}
}
