package openapi30

import (
	"errors"
	"testing"
)

func TestOAuthFlow_SetAuthorizationURL_WithoutHook(t *testing.T) {
	f := NewOAuthFlow("http://old.com", "", "", nil)
	err := f.SetAuthorizationURL("http://new.com")
	if err != nil {
		t.Fatalf("SetAuthorizationURL without hook should succeed, got %v", err)
	}
	if f.AuthorizationURL() != "http://new.com" {
		t.Errorf("AuthorizationURL() = %q, want http://new.com", f.AuthorizationURL())
	}
}

func TestOAuthFlow_SetAuthorizationURL_WithHook_Rejects(t *testing.T) {
	f := NewOAuthFlow("http://old.com", "", "", nil)
	rejectErr := errors.New("rejected")
	f.Trix.OnSet("authorizationUrl", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	err := f.SetAuthorizationURL("http://new.com")
	if err != rejectErr {
		t.Errorf("SetAuthorizationURL with rejecting hook should return error, got %v", err)
	}
	if f.AuthorizationURL() != "http://old.com" {
		t.Errorf("AuthorizationURL should be unchanged after rejection, got %q", f.AuthorizationURL())
	}
}

func TestOAuthFlow_SetTokenURL_WithoutHook(t *testing.T) {
	f := NewOAuthFlow("", "http://old.com", "", nil)
	err := f.SetTokenURL("http://new.com")
	if err != nil {
		t.Fatalf("SetTokenURL without hook should succeed, got %v", err)
	}
	if f.TokenURL() != "http://new.com" {
		t.Errorf("TokenURL() = %q, want http://new.com", f.TokenURL())
	}
}

func TestOAuthFlow_SetRefreshURL_WithoutHook(t *testing.T) {
	f := NewOAuthFlow("", "", "http://old.com", nil)
	err := f.SetRefreshURL("http://new.com")
	if err != nil {
		t.Fatalf("SetRefreshURL without hook should succeed, got %v", err)
	}
	if f.RefreshURL() != "http://new.com" {
		t.Errorf("RefreshURL() = %q, want http://new.com", f.RefreshURL())
	}
}

func TestOAuthFlow_SetScopes_WithoutHook(t *testing.T) {
	f := NewOAuthFlow("", "", "", nil)
	scopes := map[string]string{"read": "read access"}
	err := f.SetScopes(scopes)
	if err != nil {
		t.Fatalf("SetScopes without hook should succeed, got %v", err)
	}
	if f.Scopes()["read"] != "read access" {
		t.Errorf("Scopes() = %v, want %v", f.Scopes(), scopes)
	}
}
