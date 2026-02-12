package openapi20

import (
	"encoding/json"
	"testing"
)

func TestSchemaRef_MarshalJSON_Ref(t *testing.T) {
	// Arrange
	ref := NewSchemaRef("#/definitions/Pet")

	// Act
	got, err := json.Marshal(ref)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := `{"$ref":"#/definitions/Pet"}`
	if string(got) != expected {
		t.Errorf("got %s, want %s", got, expected)
	}
}

func TestSchemaRef_MarshalJSON_InlineValue(t *testing.T) {
	// Arrange
	ref := &SchemaRef{}
	ref.SetValue(NewSchema(SchemaFields{Type: "string"}))

	// Act
	got, err := json.Marshal(ref)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := `{"type":"string"}`
	if string(got) != expected {
		t.Errorf("got %s, want %s", got, expected)
	}
}

func TestSchemaRef_MarshalJSON_Nil(t *testing.T) {
	// Arrange
	ref := &SchemaRef{}

	// Act
	got, err := json.Marshal(ref)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(got) != "null" {
		t.Errorf("got %s, want null", got)
	}
}
