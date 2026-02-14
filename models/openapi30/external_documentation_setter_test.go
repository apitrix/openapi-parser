package openapi30

import (
	"errors"
	"testing"
)

func TestExternalDocumentation_SetDescription_WithoutHook(t *testing.T) {
	ed := NewExternalDocumentation("http://x.com", "old")
	err := ed.SetDescription("new")
	if err != nil {
		t.Fatalf("SetDescription without hook should succeed, got %v", err)
	}
	if ed.Description() != "new" {
		t.Errorf("Description() = %q, want %q", ed.Description(), "new")
	}
}

func TestExternalDocumentation_SetDescription_WithHook_Rejects(t *testing.T) {
	ed := NewExternalDocumentation("http://x.com", "old")
	rejectErr := errors.New("rejected")
	ed.Trix.OnSet("description", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	err := ed.SetDescription("new")
	if err != rejectErr {
		t.Errorf("SetDescription with rejecting hook should return error, got %v", err)
	}
	if ed.Description() != "old" {
		t.Errorf("Description should be unchanged after rejection, got %q", ed.Description())
	}
}

func TestExternalDocumentation_SetURL_WithoutHook(t *testing.T) {
	ed := NewExternalDocumentation("http://old.com", "")
	err := ed.SetURL("http://new.com")
	if err != nil {
		t.Fatalf("SetURL without hook should succeed, got %v", err)
	}
	if ed.URL() != "http://new.com" {
		t.Errorf("URL() = %q, want %q", ed.URL(), "http://new.com")
	}
}

func TestExternalDocumentation_SetURL_WithHook_Rejects(t *testing.T) {
	ed := NewExternalDocumentation("http://old.com", "")
	rejectErr := errors.New("rejected")
	ed.Trix.OnSet("url", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	err := ed.SetURL("http://new.com")
	if err != rejectErr {
		t.Errorf("SetURL with rejecting hook should return error, got %v", err)
	}
	if ed.URL() != "http://old.com" {
		t.Errorf("URL should be unchanged after rejection, got %q", ed.URL())
	}
}
