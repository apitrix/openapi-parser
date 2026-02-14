package openapi20

import (
	"errors"
	"testing"
)

func TestPathItem_SetRef_WithoutHook(t *testing.T) {
	pi := NewPathItem("old", nil, nil, nil, nil, nil, nil, nil, nil)
	err := pi.SetRef("new")
	if err != nil {
		t.Fatalf("SetRef without hook should succeed, got %v", err)
	}
	if pi.Ref() != "new" {
		t.Errorf("Ref() = %q, want %q", pi.Ref(), "new")
	}
}

func TestPathItem_SetRef_WithHook_Rejects(t *testing.T) {
	pi := NewPathItem("old", nil, nil, nil, nil, nil, nil, nil, nil)
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

func TestPathItem_SetRef_WithHook_Passes(t *testing.T) {
	pi := NewPathItem("old", nil, nil, nil, nil, nil, nil, nil, nil)
	pi.Trix.OnSet("$ref", func(field string, oldVal, newVal interface{}) error {
		return nil
	})
	err := pi.SetRef("new")
	if err != nil {
		t.Fatalf("SetRef with passing hook should succeed, got %v", err)
	}
	if pi.Ref() != "new" {
		t.Errorf("Ref() = %q, want %q", pi.Ref(), "new")
	}
}

func TestPathItem_SetGet_WithoutHook(t *testing.T) {
	pi := NewPathItem("", nil, nil, nil, nil, nil, nil, nil, nil)
	op := NewOperation(nil, "get pets", "", nil, "", nil, nil, nil, nil, nil, false, nil)
	err := pi.SetGet(op)
	if err != nil {
		t.Fatalf("SetGet without hook should succeed, got %v", err)
	}
	if pi.Get() != op {
		t.Errorf("Get() = %v, want %v", pi.Get(), op)
	}
}

func TestPathItem_SetParameters_WithoutHook(t *testing.T) {
	pi := NewPathItem("", nil, nil, nil, nil, nil, nil, nil, nil)
	params := []*RefParameter{NewRefParameter("#/params/1")}
	err := pi.SetParameters(params)
	if err != nil {
		t.Fatalf("SetParameters without hook should succeed, got %v", err)
	}
	if len(pi.Parameters()) != 1 {
		t.Errorf("Parameters() len = %d, want 1", len(pi.Parameters()))
	}
}

func TestPathItem_SetParameters_WithHook_Rejects(t *testing.T) {
	pi := NewPathItem("", nil, nil, nil, nil, nil, nil, nil, nil)
	rejectErr := errors.New("rejected")
	pi.Trix.OnSet("parameters", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	params := []*RefParameter{NewRefParameter("#/params/1")}
	err := pi.SetParameters(params)
	if err != rejectErr {
		t.Errorf("SetParameters with rejecting hook should return error, got %v", err)
	}
	if len(pi.Parameters()) != 0 {
		t.Errorf("Parameters should be unchanged after rejection")
	}
}
