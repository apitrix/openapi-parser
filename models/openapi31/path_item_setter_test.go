package openapi31

import (
	"errors"
	"testing"

	"openapi-parser/models/shared"
)

func TestPathItem_SetRef_WithoutHook(t *testing.T) {
	pi := NewPathItem()
	pi.ref = "old"
	err := pi.SetRef("new")
	if err != nil {
		t.Fatalf("SetRef without hook should succeed, got %v", err)
	}
	if pi.Ref() != "new" {
		t.Errorf("Ref() = %q, want %q", pi.Ref(), "new")
	}
}

func TestPathItem_SetRef_WithHook_Rejects(t *testing.T) {
	pi := NewPathItem()
	_ = pi.SetRef("old")
	rejectErr := errors.New("rejected")
	pi.Trix.OnSet("$ref", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	err := pi.SetRef("new")
	if err != rejectErr {
		t.Errorf("SetRef with rejecting hook should return error, got %v", err)
	}
	if pi.Ref() != "old" {
		t.Errorf("Ref should be unchanged after rejection, got %q", pi.Ref())
	}
}

func TestPathItem_SetSummary_WithoutHook(t *testing.T) {
	pi := NewPathItem()
	err := pi.SetSummary("new")
	if err != nil {
		t.Fatalf("SetSummary without hook should succeed, got %v", err)
	}
	if pi.Summary() != "new" {
		t.Errorf("Summary() = %q, want %q", pi.Summary(), "new")
	}
}

func TestPathItem_SetDescription_WithoutHook(t *testing.T) {
	pi := NewPathItem()
	err := pi.SetDescription("new")
	if err != nil {
		t.Fatalf("SetDescription without hook should succeed, got %v", err)
	}
	if pi.Description() != "new" {
		t.Errorf("Description() = %q, want %q", pi.Description(), "new")
	}
}

func TestPathItem_SetGet_WithoutHook(t *testing.T) {
	pi := NewPathItem()
	op := NewOperation()
	_ = op.SetSummary("get pets")
	err := pi.SetGet(op)
	if err != nil {
		t.Fatalf("SetGet without hook should succeed, got %v", err)
	}
	if pi.Get() != op {
		t.Errorf("Get() = %v, want %v", pi.Get(), op)
	}
}

func TestPathItem_SetParameters_WithoutHook(t *testing.T) {
	pi := NewPathItem()
	params := []*shared.RefWithMeta[Parameter]{shared.NewRefWithMeta[Parameter]("#/components/parameters/1")}
	err := pi.SetParameters(params)
	if err != nil {
		t.Fatalf("SetParameters without hook should succeed, got %v", err)
	}
	if len(pi.Parameters()) != 1 {
		t.Errorf("Parameters() len = %d, want 1", len(pi.Parameters()))
	}
}
