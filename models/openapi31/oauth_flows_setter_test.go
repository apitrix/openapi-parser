package openapi31

import (
	"errors"
	"testing"
)

func TestOAuthFlows_SetImplicit_WithoutHook(t *testing.T) {
	f := NewOAuthFlows(nil, nil, nil, nil)
	flow := NewOAuthFlow("http://auth.com", "", "", nil)
	err := f.SetImplicit(flow)
	if err != nil {
		t.Fatalf("SetImplicit without hook should succeed, got %v", err)
	}
	if f.Implicit() != flow {
		t.Errorf("Implicit() = %v, want %v", f.Implicit(), flow)
	}
}

func TestOAuthFlows_SetImplicit_WithHook_Rejects(t *testing.T) {
	f := NewOAuthFlows(nil, nil, nil, nil)
	rejectErr := errors.New("rejected")
	f.Trix.OnSet("implicit", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	flow := NewOAuthFlow("http://auth.com", "", "", nil)
	err := f.SetImplicit(flow)
	if err != rejectErr {
		t.Errorf("SetImplicit with rejecting hook should return error, got %v", err)
	}
	if f.Implicit() != nil {
		t.Errorf("Implicit should be unchanged after rejection")
	}
}

func TestOAuthFlows_SetPassword_WithoutHook(t *testing.T) {
	f := NewOAuthFlows(nil, nil, nil, nil)
	flow := NewOAuthFlow("", "http://token.com", "", nil)
	err := f.SetPassword(flow)
	if err != nil {
		t.Fatalf("SetPassword without hook should succeed, got %v", err)
	}
	if f.Password() != flow {
		t.Errorf("Password() = %v, want %v", f.Password(), flow)
	}
}

func TestOAuthFlows_SetClientCredentials_WithoutHook(t *testing.T) {
	f := NewOAuthFlows(nil, nil, nil, nil)
	flow := NewOAuthFlow("", "http://token.com", "", nil)
	err := f.SetClientCredentials(flow)
	if err != nil {
		t.Fatalf("SetClientCredentials without hook should succeed, got %v", err)
	}
	if f.ClientCredentials() != flow {
		t.Errorf("ClientCredentials() = %v, want %v", f.ClientCredentials(), flow)
	}
}

func TestOAuthFlows_SetAuthorizationCode_WithoutHook(t *testing.T) {
	f := NewOAuthFlows(nil, nil, nil, nil)
	flow := NewOAuthFlow("http://auth.com", "http://token.com", "", nil)
	err := f.SetAuthorizationCode(flow)
	if err != nil {
		t.Fatalf("SetAuthorizationCode without hook should succeed, got %v", err)
	}
	if f.AuthorizationCode() != flow {
		t.Errorf("AuthorizationCode() = %v, want %v", f.AuthorizationCode(), flow)
	}
}
