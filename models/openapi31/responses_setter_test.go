package openapi31

import (
	"errors"
	"testing"
)

func TestResponses_SetDefault_WithoutHook(t *testing.T) {
	r := NewResponses(nil, nil)
	ref := NewRefResponse("#/components/responses/default")
	err := r.SetDefault(ref)
	if err != nil {
		t.Fatalf("SetDefault without hook should succeed, got %v", err)
	}
	if r.Default() != ref {
		t.Errorf("Default() = %v, want %v", r.Default(), ref)
	}
}

func TestResponses_SetDefault_WithHook_Rejects(t *testing.T) {
	r := NewResponses(nil, nil)
	rejectErr := errors.New("rejected")
	r.Trix.OnSet("default", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	ref := NewRefResponse("#/components/responses/default")
	err := r.SetDefault(ref)
	if err != rejectErr {
		t.Errorf("SetDefault with rejecting hook should return error, got %v", err)
	}
	if r.Default() != nil {
		t.Errorf("Default should be unchanged after rejection")
	}
}

func TestResponses_SetCodes_WithoutHook(t *testing.T) {
	r := NewResponses(nil, nil)
	codes := map[string]*RefResponse{"200": NewRefResponse("#/components/responses/ok")}
	err := r.SetCodes(codes)
	if err != nil {
		t.Fatalf("SetCodes without hook should succeed, got %v", err)
	}
	if r.Codes()["200"] != codes["200"] {
		t.Errorf("Codes() = %v, want %v", r.Codes(), codes)
	}
}
