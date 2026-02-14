package openapi31

import (
	"errors"
	"testing"
)

func TestParameter_SetName_WithoutHook(t *testing.T) {
	p := NewParameter(ParameterFields{Name: "old"})
	err := p.SetName("new")
	if err != nil {
		t.Fatalf("SetName without hook should succeed, got %v", err)
	}
	if p.Name() != "new" {
		t.Errorf("Name() = %q, want %q", p.Name(), "new")
	}
}

func TestParameter_SetName_WithHook_Rejects(t *testing.T) {
	p := NewParameter(ParameterFields{Name: "old"})
	rejectErr := errors.New("rejected")
	p.Trix.OnSet("name", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	err := p.SetName("new")
	if err != rejectErr {
		t.Errorf("SetName with rejecting hook should return error, got %v", err)
	}
	if p.Name() != "old" {
		t.Errorf("Name should be unchanged after rejection, got %q", p.Name())
	}
}

func TestParameter_SetIn_WithoutHook(t *testing.T) {
	p := NewParameter(ParameterFields{In: "old"})
	err := p.SetIn("new")
	if err != nil {
		t.Fatalf("SetIn without hook should succeed, got %v", err)
	}
	if p.In() != "new" {
		t.Errorf("In() = %q, want %q", p.In(), "new")
	}
}

func TestParameter_SetDescription_WithoutHook(t *testing.T) {
	p := NewParameter(ParameterFields{Description: "old"})
	err := p.SetDescription("new")
	if err != nil {
		t.Fatalf("SetDescription without hook should succeed, got %v", err)
	}
	if p.Description() != "new" {
		t.Errorf("Description() = %q, want %q", p.Description(), "new")
	}
}

func TestParameter_SetRequired_WithoutHook(t *testing.T) {
	p := NewParameter(ParameterFields{Required: false})
	err := p.SetRequired(true)
	if err != nil {
		t.Fatalf("SetRequired without hook should succeed, got %v", err)
	}
	if !p.Required() {
		t.Errorf("Required() = false, want true")
	}
}

func TestParameter_SetSchema_WithoutHook(t *testing.T) {
	p := NewParameter(ParameterFields{})
	ref := NewRefSchema("#/components/schemas/Pet")
	err := p.SetSchema(ref)
	if err != nil {
		t.Fatalf("SetSchema without hook should succeed, got %v", err)
	}
	if p.Schema() != ref {
		t.Errorf("Schema() = %v, want %v", p.Schema(), ref)
	}
}
