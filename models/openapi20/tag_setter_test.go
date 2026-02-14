package openapi20

import (
	"errors"
	"testing"
)

func TestTag_SetName_WithoutHook(t *testing.T) {
	tag := NewTag("old", "", nil)
	err := tag.SetName("new")
	if err != nil {
		t.Fatalf("SetName without hook should succeed, got %v", err)
	}
	if tag.Name() != "new" {
		t.Errorf("Name() = %q, want %q", tag.Name(), "new")
	}
}

func TestTag_SetName_WithHook_Rejects(t *testing.T) {
	tag := NewTag("old", "", nil)
	rejectErr := errors.New("rejected")
	tag.Trix.OnSet("name", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	err := tag.SetName("new")
	if err != rejectErr {
		t.Errorf("SetName with rejecting hook should return error, got %v", err)
	}
	if tag.Name() != "old" {
		t.Errorf("Name should be unchanged after rejection, got %q", tag.Name())
	}
}

func TestTag_SetName_WithHook_Passes(t *testing.T) {
	tag := NewTag("old", "", nil)
	tag.Trix.OnSet("name", func(field string, oldVal, newVal interface{}) error {
		return nil
	})
	err := tag.SetName("new")
	if err != nil {
		t.Fatalf("SetName with passing hook should succeed, got %v", err)
	}
	if tag.Name() != "new" {
		t.Errorf("Name() = %q, want %q", tag.Name(), "new")
	}
}

func TestTag_SetDescription_WithoutHook(t *testing.T) {
	tag := NewTag("x", "old", nil)
	err := tag.SetDescription("new")
	if err != nil {
		t.Fatalf("SetDescription without hook should succeed, got %v", err)
	}
	if tag.Description() != "new" {
		t.Errorf("Description() = %q, want %q", tag.Description(), "new")
	}
}

func TestTag_SetExternalDocs_WithoutHook(t *testing.T) {
	tag := NewTag("x", "", nil)
	ed := NewExternalDocs("desc", "http://x.com")
	err := tag.SetExternalDocs(ed)
	if err != nil {
		t.Fatalf("SetExternalDocs without hook should succeed, got %v", err)
	}
	if tag.ExternalDocs() != ed {
		t.Errorf("ExternalDocs() = %v, want %v", tag.ExternalDocs(), ed)
	}
}

func TestTag_SetExternalDocs_WithHook_Rejects(t *testing.T) {
	tag := NewTag("x", "", nil)
	rejectErr := errors.New("rejected")
	tag.Trix.OnSet("externalDocs", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	ed := NewExternalDocs("desc", "http://x.com")
	err := tag.SetExternalDocs(ed)
	if err != rejectErr {
		t.Errorf("SetExternalDocs with rejecting hook should return error, got %v", err)
	}
	if tag.ExternalDocs() != nil {
		t.Errorf("ExternalDocs should be unchanged after rejection")
	}
}
