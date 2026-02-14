package shared

import (
	"errors"
	"testing"
)

func TestTrix_RunHooks_NoHooks(t *testing.T) {
	trix := &Trix{}
	err := trix.RunHooks("name", "old", "new")
	if err != nil {
		t.Errorf("RunHooks with no hooks should return nil, got %v", err)
	}
}

func TestTrix_RunHooks_SingleHook_Passes(t *testing.T) {
	trix := &Trix{}
	trix.OnSet("name", func(field string, oldVal, newVal interface{}) error {
		return nil
	})
	err := trix.RunHooks("name", "old", "new")
	if err != nil {
		t.Errorf("RunHooks with passing hook should return nil, got %v", err)
	}
}

func TestTrix_RunHooks_SingleHook_Rejects(t *testing.T) {
	trix := &Trix{}
	rejectErr := errors.New("rejected")
	trix.OnSet("name", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	err := trix.RunHooks("name", "old", "new")
	if err != rejectErr {
		t.Errorf("RunHooks with rejecting hook should return error, got %v", err)
	}
}

func TestTrix_RunHooks_MultipleHooks_AllPass(t *testing.T) {
	trix := &Trix{}
	callCount := 0
	trix.OnSet("tags", func(field string, oldVal, newVal interface{}) error {
		callCount++
		return nil
	})
	trix.OnSet("tags", func(field string, oldVal, newVal interface{}) error {
		callCount++
		return nil
	})
	err := trix.RunHooks("tags", []string{}, []string{"a"})
	if err != nil {
		t.Errorf("RunHooks with multiple passing hooks should return nil, got %v", err)
	}
	if callCount != 2 {
		t.Errorf("expected 2 hook calls, got %d", callCount)
	}
}

func TestTrix_RunHooks_MultipleHooks_SecondRejects(t *testing.T) {
	trix := &Trix{}
	callCount := 0
	rejectErr := errors.New("rejected")
	trix.OnSet("tags", func(field string, oldVal, newVal interface{}) error {
		callCount++
		return nil
	})
	trix.OnSet("tags", func(field string, oldVal, newVal interface{}) error {
		callCount++
		return rejectErr
	})
	err := trix.RunHooks("tags", []string{}, []string{"a"})
	if err != rejectErr {
		t.Errorf("RunHooks should return first error, got %v", err)
	}
	if callCount != 2 {
		t.Errorf("expected 2 hook calls before rejection, got %d", callCount)
	}
}

func TestTrix_RunHooks_FieldSpecific(t *testing.T) {
	trix := &Trix{}
	trix.OnSet("name", func(field string, oldVal, newVal interface{}) error {
		return errors.New("name rejected")
	})
	// RunHooks for different field should not invoke name's hook
	err := trix.RunHooks("email", "old@x.com", "new@x.com")
	if err != nil {
		t.Errorf("RunHooks for unregistered field should return nil, got %v", err)
	}
}
