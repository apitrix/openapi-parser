package openapi31

import (
	"errors"
	"testing"
)

func TestDiscriminator_SetPropertyName_WithoutHook(t *testing.T) {
	d := NewDiscriminator("old", nil)
	err := d.SetPropertyName("new")
	if err != nil {
		t.Fatalf("SetPropertyName without hook should succeed, got %v", err)
	}
	if d.PropertyName() != "new" {
		t.Errorf("PropertyName() = %q, want %q", d.PropertyName(), "new")
	}
}

func TestDiscriminator_SetPropertyName_WithHook_Rejects(t *testing.T) {
	d := NewDiscriminator("old", nil)
	rejectErr := errors.New("rejected")
	d.Trix.OnSet("propertyName", func(field string, oldVal, newVal interface{}) error {
		return rejectErr
	})
	err := d.SetPropertyName("new")
	if err != rejectErr {
		t.Errorf("SetPropertyName with rejecting hook should return error, got %v", err)
	}
	if d.PropertyName() != "old" {
		t.Errorf("PropertyName should be unchanged after rejection, got %q", d.PropertyName())
	}
}

func TestDiscriminator_SetMapping_WithoutHook(t *testing.T) {
	d := NewDiscriminator("", nil)
	mapping := map[string]string{"dog": "#/components/schemas/Dog"}
	err := d.SetMapping(mapping)
	if err != nil {
		t.Fatalf("SetMapping without hook should succeed, got %v", err)
	}
	if d.Mapping()["dog"] != mapping["dog"] {
		t.Errorf("Mapping() = %v, want %v", d.Mapping(), mapping)
	}
}
