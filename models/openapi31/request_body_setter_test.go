package openapi31

import (
	"errors"
	"testing"
)

func TestRequestBody_SetDescription_WithoutHook(t *testing.T) {
	rb := NewRequestBody("old", nil, false)
	err := rb.SetDescription("new")
	if err != nil {
		t.Fatalf("SetDescription without hook should succeed, got %v", err)
	}
	if rb.Description() != "new" {
		t.Errorf("Description() = %q, want %q", rb.Description(), "new")
	}
}

func TestRequestBody_SetDescription_WithHook_Rejects(t *testing.T) {
	rb := NewRequestBody("old", nil, false)
	rejectErr := errors.New("rejected")
	rb.Trix.OnSet("description", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	err := rb.SetDescription("new")
	if err != rejectErr {
		t.Errorf("SetDescription with rejecting hook should return error, got %v", err)
	}
	if rb.Description() != "old" {
		t.Errorf("Description should be unchanged after rejection, got %q", rb.Description())
	}
}

func TestRequestBody_SetContent_WithoutHook(t *testing.T) {
	rb := NewRequestBody("", nil, false)
	content := map[string]*MediaType{"application/json": NewMediaType(nil, nil, nil, nil)}
	err := rb.SetContent(content)
	if err != nil {
		t.Fatalf("SetContent without hook should succeed, got %v", err)
	}
	if rb.Content()["application/json"] != content["application/json"] {
		t.Errorf("Content() = %v, want %v", rb.Content(), content)
	}
}

func TestRequestBody_SetRequired_WithoutHook(t *testing.T) {
	rb := NewRequestBody("", nil, false)
	err := rb.SetRequired(true)
	if err != nil {
		t.Fatalf("SetRequired without hook should succeed, got %v", err)
	}
	if !rb.Required() {
		t.Errorf("Required() = false, want true")
	}
}
