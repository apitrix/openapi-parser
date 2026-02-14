package openapi20

import (
	"errors"
	"testing"
)

func TestOperation_SetTags_WithoutHook(t *testing.T) {
	op := NewOperation(nil, "", "", nil, "", nil, nil, nil, nil, nil, false, nil)
	err := op.SetTags([]string{"pets", "store"})
	if err != nil {
		t.Fatalf("SetTags without hook should succeed, got %v", err)
	}
	if len(op.Tags()) != 2 || op.Tags()[0] != "pets" || op.Tags()[1] != "store" {
		t.Errorf("Tags() = %v, want [pets store]", op.Tags())
	}
}

func TestOperation_SetTags_WithHook_Rejects(t *testing.T) {
	op := NewOperation([]string{"old"}, "", "", nil, "", nil, nil, nil, nil, nil, false, nil)
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
