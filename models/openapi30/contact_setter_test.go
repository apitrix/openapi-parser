package openapi30

import (
	"errors"
	"testing"
)

func TestContact_SetName_WithoutHook(t *testing.T) {
	c := NewContact("old", "http://x.com", "old@x.com")
	err := c.SetName("new")
	if err != nil {
		t.Fatalf("SetName without hook should succeed, got %v", err)
	}
	if c.Name() != "new" {
		t.Errorf("Name() = %q, want %q", c.Name(), "new")
	}
}

func TestContact_SetName_WithHook_Rejects(t *testing.T) {
	c := NewContact("old", "http://x.com", "old@x.com")
	rejectErr := errors.New("rejected")
	c.Trix.OnSet("name", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	err := c.SetName("new")
	if err != rejectErr {
		t.Errorf("SetName with rejecting hook should return error, got %v", err)
	}
	if c.Name() != "old" {
		t.Errorf("Name should be unchanged after rejection, got %q", c.Name())
	}
}

func TestContact_SetName_WithHook_Passes(t *testing.T) {
	c := NewContact("old", "http://x.com", "old@x.com")
	c.Trix.OnSet("name", func(field string, oldVal, newVal interface{}) error {
		return nil
	})
	err := c.SetName("new")
	if err != nil {
		t.Fatalf("SetName with passing hook should succeed, got %v", err)
	}
	if c.Name() != "new" {
		t.Errorf("Name() = %q, want %q", c.Name(), "new")
	}
}

func TestContact_SetURL_WithoutHook(t *testing.T) {
	c := NewContact("x", "http://old.com", "x@x.com")
	err := c.SetURL("http://new.com")
	if err != nil {
		t.Fatalf("SetURL without hook should succeed, got %v", err)
	}
	if c.URL() != "http://new.com" {
		t.Errorf("URL() = %q, want %q", c.URL(), "http://new.com")
	}
}

func TestContact_SetURL_WithHook_Rejects(t *testing.T) {
	c := NewContact("x", "http://old.com", "x@x.com")
	rejectErr := errors.New("rejected")
	c.Trix.OnSet("url", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	err := c.SetURL("http://new.com")
	if err != rejectErr {
		t.Errorf("SetURL with rejecting hook should return error, got %v", err)
	}
	if c.URL() != "http://old.com" {
		t.Errorf("URL should be unchanged after rejection, got %q", c.URL())
	}
}

func TestContact_SetEmail_WithoutHook(t *testing.T) {
	c := NewContact("x", "http://x.com", "old@x.com")
	err := c.SetEmail("new@x.com")
	if err != nil {
		t.Fatalf("SetEmail without hook should succeed, got %v", err)
	}
	if c.Email() != "new@x.com" {
		t.Errorf("Email() = %q, want %q", c.Email(), "new@x.com")
	}
}
