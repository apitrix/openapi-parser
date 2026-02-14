package openapi31

import (
	"errors"
	"testing"

	"openapi-parser/models/shared"
)

func TestMediaType_SetSchema_WithoutHook(t *testing.T) {
	m := NewMediaType(nil, nil, nil, nil)
	ref := shared.NewRefWithMeta[Schema]("#/components/schemas/Pet")
	err := m.SetSchema(ref)
	if err != nil {
		t.Fatalf("SetSchema without hook should succeed, got %v", err)
	}
	if m.Schema() != ref {
		t.Errorf("Schema() = %v, want %v", m.Schema(), ref)
	}
}

func TestMediaType_SetSchema_WithHook_Rejects(t *testing.T) {
	m := NewMediaType(nil, nil, nil, nil)
	rejectErr := errors.New("rejected")
	m.Trix.OnSet("schema", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	ref := shared.NewRefWithMeta[Schema]("#/components/schemas/Pet")
	err := m.SetSchema(ref)
	if err != rejectErr {
		t.Errorf("SetSchema with rejecting hook should return error, got %v", err)
	}
	if m.Schema() != nil {
		t.Errorf("Schema should be unchanged after rejection")
	}
}

func TestMediaType_SetExample_WithoutHook(t *testing.T) {
	m := NewMediaType(nil, nil, nil, nil)
	err := m.SetExample(map[string]string{"name": "Puma"})
	if err != nil {
		t.Fatalf("SetExample without hook should succeed, got %v", err)
	}
	if m.Example() == nil {
		t.Errorf("Example() = nil, want non-nil")
	}
}

func TestMediaType_SetExamples_WithoutHook(t *testing.T) {
	m := NewMediaType(nil, nil, nil, nil)
	examples := map[string]*shared.RefWithMeta[Example]{"ex1": shared.NewRefWithMeta[Example]("#/components/examples/ex1")}
	err := m.SetExamples(examples)
	if err != nil {
		t.Fatalf("SetExamples without hook should succeed, got %v", err)
	}
	if m.Examples()["ex1"] != examples["ex1"] {
		t.Errorf("Examples() = %v, want %v", m.Examples(), examples)
	}
}

func TestMediaType_SetEncoding_WithoutHook(t *testing.T) {
	m := NewMediaType(nil, nil, nil, nil)
	enc := map[string]*Encoding{"application/json": NewEncoding("application/json", "", nil, nil, false)}
	err := m.SetEncoding(enc)
	if err != nil {
		t.Fatalf("SetEncoding without hook should succeed, got %v", err)
	}
	if m.Encoding()["application/json"] != enc["application/json"] {
		t.Errorf("Encoding() = %v, want %v", m.Encoding(), enc)
	}
}
