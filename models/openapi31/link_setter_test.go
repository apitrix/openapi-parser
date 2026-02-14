package openapi31

import (
	"errors"
	"testing"
)

func TestLink_SetOperationRef_WithoutHook(t *testing.T) {
	l := NewLink("old", "", "", nil, nil, nil)
	err := l.SetOperationRef("new")
	if err != nil {
		t.Fatalf("SetOperationRef without hook should succeed, got %v", err)
	}
	if l.OperationRef() != "new" {
		t.Errorf("OperationRef() = %q, want %q", l.OperationRef(), "new")
	}
}

func TestLink_SetOperationRef_WithHook_Rejects(t *testing.T) {
	l := NewLink("old", "", "", nil, nil, nil)
	rejectErr := errors.New("rejected")
	l.Trix.OnSet("operationRef", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	err := l.SetOperationRef("new")
	if err != rejectErr {
		t.Errorf("SetOperationRef with rejecting hook should return error, got %v", err)
	}
	if l.OperationRef() != "old" {
		t.Errorf("OperationRef should be unchanged after rejection, got %q", l.OperationRef())
	}
}

func TestLink_SetOperationID_WithoutHook(t *testing.T) {
	l := NewLink("", "old", "", nil, nil, nil)
	err := l.SetOperationID("new")
	if err != nil {
		t.Fatalf("SetOperationID without hook should succeed, got %v", err)
	}
	if l.OperationID() != "new" {
		t.Errorf("OperationID() = %q, want %q", l.OperationID(), "new")
	}
}

func TestLink_SetDescription_WithoutHook(t *testing.T) {
	l := NewLink("", "", "old", nil, nil, nil)
	err := l.SetDescription("new")
	if err != nil {
		t.Fatalf("SetDescription without hook should succeed, got %v", err)
	}
	if l.Description() != "new" {
		t.Errorf("Description() = %q, want %q", l.Description(), "new")
	}
}

func TestLink_SetServer_WithoutHook(t *testing.T) {
	l := NewLink("", "", "", nil, nil, nil)
	srv := NewServer("http://api.com", "", nil)
	err := l.SetServer(srv)
	if err != nil {
		t.Fatalf("SetServer without hook should succeed, got %v", err)
	}
	if l.Server() != srv {
		t.Errorf("Server() = %v, want %v", l.Server(), srv)
	}
}
