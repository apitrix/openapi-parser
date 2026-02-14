package openapi30

import (
	"errors"
	"testing"

	"openapi-parser/models/shared"
)

func TestOperation_SetTags_WithoutHook(t *testing.T) {
	op := NewOperation([]string{"old"}, "", "", nil, "", nil, nil, nil, nil, false, nil, nil)
	err := op.SetTags([]string{"new"})
	if err != nil {
		t.Fatalf("SetTags without hook should succeed, got %v", err)
	}
	if len(op.Tags()) != 1 || op.Tags()[0] != "new" {
		t.Errorf("Tags() = %v, want [new]", op.Tags())
	}
}

func TestOperation_SetTags_WithHook_Rejects(t *testing.T) {
	op := NewOperation([]string{"old"}, "", "", nil, "", nil, nil, nil, nil, false, nil, nil)
	rejectErr := errors.New("rejected")
	op.Trix.OnSet("tags", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	err := op.SetTags([]string{"new"})
	if err != rejectErr {
		t.Errorf("SetTags with rejecting hook should return error, got %v", err)
	}
	if len(op.Tags()) != 1 || op.Tags()[0] != "old" {
		t.Errorf("Tags should be unchanged after rejection, got %v", op.Tags())
	}
}

func TestOperation_SetSummary_WithoutHook(t *testing.T) {
	op := NewOperation(nil, "old", "", nil, "", nil, nil, nil, nil, false, nil, nil)
	err := op.SetSummary("new")
	if err != nil {
		t.Fatalf("SetSummary without hook should succeed, got %v", err)
	}
	if op.Summary() != "new" {
		t.Errorf("Summary() = %q, want %q", op.Summary(), "new")
	}
}

func TestOperation_SetOperationID_WithoutHook(t *testing.T) {
	op := NewOperation(nil, "", "", nil, "old", nil, nil, nil, nil, false, nil, nil)
	err := op.SetOperationID("new")
	if err != nil {
		t.Fatalf("SetOperationID without hook should succeed, got %v", err)
	}
	if op.OperationID() != "new" {
		t.Errorf("OperationID() = %q, want %q", op.OperationID(), "new")
	}
}

func TestOperation_SetParameters_WithoutHook(t *testing.T) {
	op := NewOperation(nil, "", "", nil, "", nil, nil, nil, nil, false, nil, nil)
	params := []*shared.Ref[Parameter]{shared.NewRef[Parameter]("#/components/parameters/1")}
	err := op.SetParameters(params)
	if err != nil {
		t.Fatalf("SetParameters without hook should succeed, got %v", err)
	}
	if len(op.Parameters()) != 1 {
		t.Errorf("Parameters() len = %d, want 1", len(op.Parameters()))
	}
}

func TestOperation_SetDeprecated_WithoutHook(t *testing.T) {
	op := NewOperation(nil, "", "", nil, "", nil, nil, nil, nil, false, nil, nil)
	err := op.SetDeprecated(true)
	if err != nil {
		t.Fatalf("SetDeprecated without hook should succeed, got %v", err)
	}
	if !op.Deprecated() {
		t.Errorf("Deprecated() = false, want true")
	}
}
