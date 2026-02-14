package openapi30

import (
	"errors"
	"testing"
)

func TestServer_SetURL_WithoutHook(t *testing.T) {
	s := NewServer("http://old.com", "", nil)
	err := s.SetURL("http://new.com")
	if err != nil {
		t.Fatalf("SetURL without hook should succeed, got %v", err)
	}
	if s.URL() != "http://new.com" {
		t.Errorf("URL() = %q, want %q", s.URL(), "http://new.com")
	}
}

func TestServer_SetURL_WithHook_Rejects(t *testing.T) {
	s := NewServer("http://old.com", "", nil)
	rejectErr := errors.New("rejected")
	s.Trix.OnSet("url", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	err := s.SetURL("http://new.com")
	if err != rejectErr {
		t.Errorf("SetURL with rejecting hook should return error, got %v", err)
	}
	if s.URL() != "http://old.com" {
		t.Errorf("URL should be unchanged after rejection, got %q", s.URL())
	}
}

func TestServer_SetDescription_WithoutHook(t *testing.T) {
	s := NewServer("", "old", nil)
	err := s.SetDescription("new")
	if err != nil {
		t.Fatalf("SetDescription without hook should succeed, got %v", err)
	}
	if s.Description() != "new" {
		t.Errorf("Description() = %q, want %q", s.Description(), "new")
	}
}

func TestServer_SetVariables_WithoutHook(t *testing.T) {
	s := NewServer("", "", nil)
	vars := map[string]*ServerVariable{"basePath": NewServerVariable("/v1", "API base path", nil)}
	err := s.SetVariables(vars)
	if err != nil {
		t.Fatalf("SetVariables without hook should succeed, got %v", err)
	}
	if s.Variables()["basePath"] != vars["basePath"] {
		t.Errorf("Variables() = %v, want %v", s.Variables(), vars)
	}
}
