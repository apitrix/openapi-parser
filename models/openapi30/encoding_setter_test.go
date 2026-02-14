package openapi30

import (
	"errors"
	"testing"
)

func TestEncoding_SetContentType_WithoutHook(t *testing.T) {
	e := NewEncoding("old", nil, "", nil, false)
	err := e.SetContentType("new")
	if err != nil {
		t.Fatalf("SetContentType without hook should succeed, got %v", err)
	}
	if e.ContentType() != "new" {
		t.Errorf("ContentType() = %q, want %q", e.ContentType(), "new")
	}
}

func TestEncoding_SetContentType_WithHook_Rejects(t *testing.T) {
	e := NewEncoding("old", nil, "", nil, false)
	rejectErr := errors.New("rejected")
	e.Trix.OnSet("contentType", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	err := e.SetContentType("new")
	if err != rejectErr {
		t.Errorf("SetContentType with rejecting hook should return error, got %v", err)
	}
	if e.ContentType() != "old" {
		t.Errorf("ContentType should be unchanged after rejection, got %q", e.ContentType())
	}
}

func TestEncoding_SetHeaders_WithoutHook(t *testing.T) {
	e := NewEncoding("", nil, "", nil, false)
	headers := map[string]*RefHeader{"X-Custom": NewRefHeader("#/components/headers/X-Custom")}
	err := e.SetHeaders(headers)
	if err != nil {
		t.Fatalf("SetHeaders without hook should succeed, got %v", err)
	}
	if e.Headers()["X-Custom"] != headers["X-Custom"] {
		t.Errorf("Headers() = %v, want %v", e.Headers(), headers)
	}
}

func TestEncoding_SetStyle_WithoutHook(t *testing.T) {
	e := NewEncoding("", nil, "old", nil, false)
	err := e.SetStyle("new")
	if err != nil {
		t.Fatalf("SetStyle without hook should succeed, got %v", err)
	}
	if e.Style() != "new" {
		t.Errorf("Style() = %q, want %q", e.Style(), "new")
	}
}

func TestEncoding_SetAllowReserved_WithoutHook(t *testing.T) {
	e := NewEncoding("", nil, "", nil, false)
	err := e.SetAllowReserved(true)
	if err != nil {
		t.Fatalf("SetAllowReserved without hook should succeed, got %v", err)
	}
	if !e.AllowReserved() {
		t.Errorf("AllowReserved() = false, want true")
	}
}
