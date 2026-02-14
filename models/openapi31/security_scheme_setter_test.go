package openapi31

import (
	"errors"
	"testing"
)

func TestSecurityScheme_SetType_WithoutHook(t *testing.T) {
	s := NewSecurityScheme("old", "", "", "", "", "", "", nil)
	err := s.SetType("new")
	if err != nil {
		t.Fatalf("SetType without hook should succeed, got %v", err)
	}
	if s.Type() != "new" {
		t.Errorf("Type() = %q, want %q", s.Type(), "new")
	}
}

func TestSecurityScheme_SetType_WithHook_Rejects(t *testing.T) {
	s := NewSecurityScheme("old", "", "", "", "", "", "", nil)
	rejectErr := errors.New("rejected")
	s.Trix.OnSet("type", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	err := s.SetType("new")
	if err != rejectErr {
		t.Errorf("SetType with rejecting hook should return error, got %v", err)
	}
	if s.Type() != "old" {
		t.Errorf("Type should be unchanged after rejection, got %q", s.Type())
	}
}

func TestSecurityScheme_SetDescription_WithoutHook(t *testing.T) {
	s := NewSecurityScheme("", "old", "", "", "", "", "", nil)
	err := s.SetDescription("new")
	if err != nil {
		t.Fatalf("SetDescription without hook should succeed, got %v", err)
	}
	if s.Description() != "new" {
		t.Errorf("Description() = %q, want %q", s.Description(), "new")
	}
}

func TestSecurityScheme_SetName_WithoutHook(t *testing.T) {
	s := NewSecurityScheme("", "", "old", "", "", "", "", nil)
	err := s.SetName("new")
	if err != nil {
		t.Fatalf("SetName without hook should succeed, got %v", err)
	}
	if s.Name() != "new" {
		t.Errorf("Name() = %q, want %q", s.Name(), "new")
	}
}

func TestSecurityScheme_SetIn_WithoutHook(t *testing.T) {
	s := NewSecurityScheme("", "", "", "old", "", "", "", nil)
	err := s.SetIn("new")
	if err != nil {
		t.Fatalf("SetIn without hook should succeed, got %v", err)
	}
	if s.In() != "new" {
		t.Errorf("In() = %q, want %q", s.In(), "new")
	}
}

func TestSecurityScheme_SetScheme_WithoutHook(t *testing.T) {
	s := NewSecurityScheme("", "", "", "", "old", "", "", nil)
	err := s.SetScheme("new")
	if err != nil {
		t.Fatalf("SetScheme without hook should succeed, got %v", err)
	}
	if s.Scheme() != "new" {
		t.Errorf("Scheme() = %q, want %q", s.Scheme(), "new")
	}
}

func TestSecurityScheme_SetBearerFormat_WithoutHook(t *testing.T) {
	s := NewSecurityScheme("", "", "", "", "", "old", "", nil)
	err := s.SetBearerFormat("new")
	if err != nil {
		t.Fatalf("SetBearerFormat without hook should succeed, got %v", err)
	}
	if s.BearerFormat() != "new" {
		t.Errorf("BearerFormat() = %q, want %q", s.BearerFormat(), "new")
	}
}

func TestSecurityScheme_SetOpenIDConnectURL_WithoutHook(t *testing.T) {
	s := NewSecurityScheme("", "", "", "", "", "", "http://old.com", nil)
	err := s.SetOpenIDConnectURL("http://new.com")
	if err != nil {
		t.Fatalf("SetOpenIDConnectURL without hook should succeed, got %v", err)
	}
	if s.OpenIDConnectURL() != "http://new.com" {
		t.Errorf("OpenIDConnectURL() = %q, want http://new.com", s.OpenIDConnectURL())
	}
}
