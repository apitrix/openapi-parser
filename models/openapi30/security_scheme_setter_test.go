package openapi30

import (
	"errors"
	"testing"
)

func TestSecurityScheme_SetType_WithoutHook(t *testing.T) {
	ss := NewSecurityScheme("old", "", "", "", "", "", nil, "")
	err := ss.SetType("new")
	if err != nil {
		t.Fatalf("SetType without hook should succeed, got %v", err)
	}
	if ss.Type() != "new" {
		t.Errorf("Type() = %q, want %q", ss.Type(), "new")
	}
}

func TestSecurityScheme_SetType_WithHook_Rejects(t *testing.T) {
	ss := NewSecurityScheme("old", "", "", "", "", "", nil, "")
	rejectErr := errors.New("rejected")
	ss.Trix.OnSet("type", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	err := ss.SetType("new")
	if err != rejectErr {
		t.Errorf("SetType with rejecting hook should return error, got %v", err)
	}
	if ss.Type() != "old" {
		t.Errorf("Type should be unchanged after rejection, got %q", ss.Type())
	}
}

func TestSecurityScheme_SetDescription_WithoutHook(t *testing.T) {
	ss := NewSecurityScheme("", "old", "", "", "", "", nil, "")
	err := ss.SetDescription("new")
	if err != nil {
		t.Fatalf("SetDescription without hook should succeed, got %v", err)
	}
	if ss.Description() != "new" {
		t.Errorf("Description() = %q, want %q", ss.Description(), "new")
	}
}

func TestSecurityScheme_SetName_WithoutHook(t *testing.T) {
	ss := NewSecurityScheme("", "", "old", "", "", "", nil, "")
	err := ss.SetName("new")
	if err != nil {
		t.Fatalf("SetName without hook should succeed, got %v", err)
	}
	if ss.Name() != "new" {
		t.Errorf("Name() = %q, want %q", ss.Name(), "new")
	}
}

func TestSecurityScheme_SetIn_WithoutHook(t *testing.T) {
	ss := NewSecurityScheme("", "", "", "old", "", "", nil, "")
	err := ss.SetIn("new")
	if err != nil {
		t.Fatalf("SetIn without hook should succeed, got %v", err)
	}
	if ss.In() != "new" {
		t.Errorf("In() = %q, want %q", ss.In(), "new")
	}
}

func TestSecurityScheme_SetScheme_WithoutHook(t *testing.T) {
	ss := NewSecurityScheme("", "", "", "", "old", "", nil, "")
	err := ss.SetScheme("new")
	if err != nil {
		t.Fatalf("SetScheme without hook should succeed, got %v", err)
	}
	if ss.Scheme() != "new" {
		t.Errorf("Scheme() = %q, want %q", ss.Scheme(), "new")
	}
}

func TestSecurityScheme_SetBearerFormat_WithoutHook(t *testing.T) {
	ss := NewSecurityScheme("", "", "", "", "", "old", nil, "")
	err := ss.SetBearerFormat("new")
	if err != nil {
		t.Fatalf("SetBearerFormat without hook should succeed, got %v", err)
	}
	if ss.BearerFormat() != "new" {
		t.Errorf("BearerFormat() = %q, want %q", ss.BearerFormat(), "new")
	}
}

func TestSecurityScheme_SetOpenIDConnectURL_WithoutHook(t *testing.T) {
	ss := NewSecurityScheme("", "", "", "", "", "", nil, "http://old.com")
	err := ss.SetOpenIDConnectURL("http://new.com")
	if err != nil {
		t.Fatalf("SetOpenIDConnectURL without hook should succeed, got %v", err)
	}
	if ss.OpenIDConnectURL() != "http://new.com" {
		t.Errorf("OpenIDConnectURL() = %q, want http://new.com", ss.OpenIDConnectURL())
	}
}
