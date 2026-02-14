package openapi30

import (
	"errors"
	"testing"
)

func TestOpenAPI_SetOpenAPIVersion_WithoutHook(t *testing.T) {
	o := NewOpenAPI("3.0.0", NewInfo("API", "", "", "1.0", nil, nil))
	o.openAPI = "old"
	err := o.SetOpenAPIVersion("new")
	if err != nil {
		t.Fatalf("SetOpenAPIVersion without hook should succeed, got %v", err)
	}
	if o.OpenAPIVersion() != "new" {
		t.Errorf("OpenAPIVersion() = %q, want %q", o.OpenAPIVersion(), "new")
	}
}

func TestOpenAPI_SetOpenAPIVersion_WithHook_Rejects(t *testing.T) {
	o := NewOpenAPI("3.0.0", NewInfo("API", "", "", "1.0", nil, nil))
	o.openAPI = "old"
	rejectErr := errors.New("rejected")
	o.Trix.OnSet("openapi", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	err := o.SetOpenAPIVersion("new")
	if err != rejectErr {
		t.Errorf("SetOpenAPIVersion with rejecting hook should return error, got %v", err)
	}
	if o.OpenAPIVersion() != "old" {
		t.Errorf("OpenAPIVersion should be unchanged after rejection, got %q", o.OpenAPIVersion())
	}
}

func TestOpenAPI_SetInfo_WithoutHook(t *testing.T) {
	o := NewOpenAPI("3.0.0", nil)
	info := NewInfo("API", "", "", "1.0", nil, nil)
	err := o.SetInfo(info)
	if err != nil {
		t.Fatalf("SetInfo without hook should succeed, got %v", err)
	}
	if o.Info() != info {
		t.Errorf("Info() = %v, want %v", o.Info(), info)
	}
}

func TestOpenAPI_SetPaths_WithoutHook(t *testing.T) {
	o := NewOpenAPI("3.0.0", NewInfo("API", "", "", "1.0", nil, nil))
	paths := NewPaths(map[string]*PathItem{"/pets": NewPathItem("", "", "", nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)})
	err := o.SetPaths(paths)
	if err != nil {
		t.Fatalf("SetPaths without hook should succeed, got %v", err)
	}
	if o.Paths() != paths {
		t.Errorf("Paths() = %v, want %v", o.Paths(), paths)
	}
}
