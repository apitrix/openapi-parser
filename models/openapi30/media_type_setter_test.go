package openapi30

import (
	"errors"
	"testing"
)

func TestMediaType_SetSchema_WithoutHook(t *testing.T) {
	mt := NewMediaType(nil, nil, nil, nil)
	ref := NewRefSchema("#/components/schemas/Pet")
	err := mt.SetSchema(ref)
	if err != nil {
		t.Fatalf("SetSchema without hook should succeed, got %v", err)
	}
	if mt.Schema() != ref {
		t.Errorf("Schema() = %v, want %v", mt.Schema(), ref)
	}
}

func TestMediaType_SetSchema_WithHook_Rejects(t *testing.T) {
	mt := NewMediaType(nil, nil, nil, nil)
	rejectErr := errors.New("rejected")
	mt.Trix.OnSet("schema", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	ref := NewRefSchema("#/components/schemas/Pet")
	err := mt.SetSchema(ref)
	if err != rejectErr {
		t.Errorf("SetSchema with rejecting hook should return error, got %v", err)
	}
	if mt.Schema() != nil {
		t.Errorf("Schema should be unchanged after rejection")
	}
}

func TestMediaType_SetExample_WithoutHook(t *testing.T) {
	mt := NewMediaType(nil, nil, nil, nil)
	err := mt.SetExample(map[string]string{"name": "Puma"})
	if err != nil {
		t.Fatalf("SetExample without hook should succeed, got %v", err)
	}
	if mt.Example() == nil {
		t.Errorf("Example() = nil, want non-nil")
	}
}

func TestMediaType_SetExamples_WithoutHook(t *testing.T) {
	mt := NewMediaType(nil, nil, nil, nil)
	examples := map[string]*RefExample{"ex1": NewRefExample("#/components/examples/ex1")}
	err := mt.SetExamples(examples)
	if err != nil {
		t.Fatalf("SetExamples without hook should succeed, got %v", err)
	}
	if mt.Examples()["ex1"] != examples["ex1"] {
		t.Errorf("Examples() = %v, want %v", mt.Examples(), examples)
	}
}

func TestMediaType_SetEncoding_WithoutHook(t *testing.T) {
	mt := NewMediaType(nil, nil, nil, nil)
	enc := map[string]*Encoding{"application/json": NewEncoding("application/json", nil, "", nil, false)}
	err := mt.SetEncoding(enc)
	if err != nil {
		t.Fatalf("SetEncoding without hook should succeed, got %v", err)
	}
	if mt.Encoding()["application/json"] != enc["application/json"] {
		t.Errorf("Encoding() = %v, want %v", mt.Encoding(), enc)
	}
}
