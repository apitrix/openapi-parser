package openapi31

import (
	"errors"
	"testing"
)

func TestLicense_SetName_WithoutHook(t *testing.T) {
	l := NewLicense("old", "", "http://x.com")
	err := l.SetName("new")
	if err != nil {
		t.Fatalf("SetName without hook should succeed, got %v", err)
	}
	if l.Name() != "new" {
		t.Errorf("Name() = %q, want %q", l.Name(), "new")
	}
}

func TestLicense_SetName_WithHook_Rejects(t *testing.T) {
	l := NewLicense("old", "", "http://x.com")
	rejectErr := errors.New("rejected")
	l.Trix.OnSet("name", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	err := l.SetName("new")
	if err != rejectErr {
		t.Errorf("SetName with rejecting hook should return error, got %v", err)
	}
	if l.Name() != "old" {
		t.Errorf("Name should be unchanged after rejection, got %q", l.Name())
	}
}

func TestLicense_SetIdentifier_WithoutHook(t *testing.T) {
	l := NewLicense("MIT", "old", "http://x.com")
	err := l.SetIdentifier("new")
	if err != nil {
		t.Fatalf("SetIdentifier without hook should succeed, got %v", err)
	}
	if l.Identifier() != "new" {
		t.Errorf("Identifier() = %q, want %q", l.Identifier(), "new")
	}
}

func TestLicense_SetURL_WithoutHook(t *testing.T) {
	l := NewLicense("MIT", "", "http://old.com")
	err := l.SetURL("http://new.com")
	if err != nil {
		t.Fatalf("SetURL without hook should succeed, got %v", err)
	}
	if l.URL() != "http://new.com" {
		t.Errorf("URL() = %q, want %q", l.URL(), "http://new.com")
	}
}
