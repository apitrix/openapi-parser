package openapi20

import (
	"errors"
	"testing"
)

func TestResponse_SetDescription_WithoutHook(t *testing.T) {
	r := NewResponse("old", nil, nil, nil)
	err := r.SetDescription("new")
	if err != nil {
		t.Fatalf("SetDescription without hook should succeed, got %v", err)
	}
	if r.Description() != "new" {
		t.Errorf("Description() = %q, want %q", r.Description(), "new")
	}
}

func TestResponse_SetDescription_WithHook_Rejects(t *testing.T) {
	r := NewResponse("old", nil, nil, nil)
	rejectErr := errors.New("rejected")
	r.Trix.OnSet("description", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	err := r.SetDescription("new")
	if err != rejectErr {
		t.Errorf("SetDescription with rejecting hook should return error, got %v", err)
	}
	if r.Description() != "old" {
		t.Errorf("Description should be unchanged after rejection, got %q", r.Description())
	}
}

func TestResponse_SetDescription_WithHook_Passes(t *testing.T) {
	r := NewResponse("old", nil, nil, nil)
	r.Trix.OnSet("description", func(field string, oldVal, newVal interface{}) error {
		return nil
	})
	err := r.SetDescription("new")
	if err != nil {
		t.Fatalf("SetDescription with passing hook should succeed, got %v", err)
	}
	if r.Description() != "new" {
		t.Errorf("Description() = %q, want %q", r.Description(), "new")
	}
}

func TestResponse_SetSchema_WithoutHook(t *testing.T) {
	r := NewResponse("", nil, nil, nil)
	ref := NewRefSchema("#/definitions/Pet")
	err := r.SetSchema(ref)
	if err != nil {
		t.Fatalf("SetSchema without hook should succeed, got %v", err)
	}
	if r.Schema() != ref {
		t.Errorf("Schema() = %v, want %v", r.Schema(), ref)
	}
}

func TestResponse_SetHeaders_WithoutHook(t *testing.T) {
	r := NewResponse("", nil, nil, nil)
	headers := map[string]*Header{"X-Rate-Limit": NewHeader(HeaderFields{Description: "rate limit"})}
	err := r.SetHeaders(headers)
	if err != nil {
		t.Fatalf("SetHeaders without hook should succeed, got %v", err)
	}
	if r.Headers()["X-Rate-Limit"] != headers["X-Rate-Limit"] {
		t.Errorf("Headers() = %v, want %v", r.Headers(), headers)
	}
}

func TestResponse_SetExamples_WithoutHook(t *testing.T) {
	r := NewResponse("", nil, nil, nil)
	examples := map[string]interface{}{"application/json": map[string]string{"name": "Puma"}}
	err := r.SetExamples(examples)
	if err != nil {
		t.Fatalf("SetExamples without hook should succeed, got %v", err)
	}
	if r.Examples()["application/json"] == nil {
		t.Errorf("Examples() = %v, want %v", r.Examples(), examples)
	}
}
