package openapi20

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

func TestSecurityScheme_SetType_WithHook_Passes(t *testing.T) {
	s := NewSecurityScheme("old", "", "", "", "", "", "", nil)
	s.Trix.OnSet("type", func(field string, oldVal, newVal interface{}) error {
		return nil
	})
	err := s.SetType("new")
	if err != nil {
		t.Fatalf("SetType with passing hook should succeed, got %v", err)
	}
	if s.Type() != "new" {
		t.Errorf("Type() = %q, want %q", s.Type(), "new")
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

func TestSecurityScheme_SetFlow_WithoutHook(t *testing.T) {
	s := NewSecurityScheme("", "", "", "", "old", "", "", nil)
	err := s.SetFlow("new")
	if err != nil {
		t.Fatalf("SetFlow without hook should succeed, got %v", err)
	}
	if s.Flow() != "new" {
		t.Errorf("Flow() = %q, want %q", s.Flow(), "new")
	}
}

func TestSecurityScheme_SetAuthorizationURL_WithoutHook(t *testing.T) {
	s := NewSecurityScheme("", "", "", "", "", "http://old.com", "", nil)
	err := s.SetAuthorizationURL("http://new.com")
	if err != nil {
		t.Fatalf("SetAuthorizationURL without hook should succeed, got %v", err)
	}
	if s.AuthorizationURL() != "http://new.com" {
		t.Errorf("AuthorizationURL() = %q, want %q", s.AuthorizationURL(), "http://new.com")
	}
}

func TestSecurityScheme_SetTokenURL_WithoutHook(t *testing.T) {
	s := NewSecurityScheme("", "", "", "", "", "", "http://old.com", nil)
	err := s.SetTokenURL("http://new.com")
	if err != nil {
		t.Fatalf("SetTokenURL without hook should succeed, got %v", err)
	}
	if s.TokenURL() != "http://new.com" {
		t.Errorf("TokenURL() = %q, want %q", s.TokenURL(), "http://new.com")
	}
}

func TestSecurityScheme_SetScopes_WithoutHook(t *testing.T) {
	s := NewSecurityScheme("", "", "", "", "", "", "", nil)
	scopes := map[string]string{"read": "read access"}
	err := s.SetScopes(scopes)
	if err != nil {
		t.Fatalf("SetScopes without hook should succeed, got %v", err)
	}
	if s.Scopes()["read"] != "read access" {
		t.Errorf("Scopes() = %v, want %v", s.Scopes(), scopes)
	}
}

func TestSecurityScheme_SetScopes_WithHook_Rejects(t *testing.T) {
	s := NewSecurityScheme("", "", "", "", "", "", "", nil)
	rejectErr := errors.New("rejected")
	s.Trix.OnSet("scopes", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	scopes := map[string]string{"read": "read access"}
	err := s.SetScopes(scopes)
	if err != rejectErr {
		t.Errorf("SetScopes with rejecting hook should return error, got %v", err)
	}
	if s.Scopes() != nil {
		t.Errorf("Scopes should be unchanged after rejection")
	}
}
