package openapi30

import (
	"encoding/json"
	"testing"
)

func TestSchemaRef_MarshalJSON_Ref(t *testing.T) {
	// Arrange
	ref := NewSchemaRef("#/components/schemas/Pet")

	// Act
	got, err := json.Marshal(ref)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"$ref":"#/components/schemas/Pet"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestSchemaRef_MarshalJSON_InlineValue(t *testing.T) {
	// Arrange
	schema := NewSchema(SchemaFields{Type: "string", Description: "A name"})
	ref := &SchemaRef{Value: schema}

	// Act
	got, err := json.Marshal(ref)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if _, ok := result["type"]; !ok {
		t.Error("expected 'type' key from inline schema")
	}
	if _, ok := result["$ref"]; ok {
		t.Error("unexpected '$ref' key for inline schema")
	}
}

func TestSchemaRef_MarshalJSON_NilValue(t *testing.T) {
	// Arrange
	ref := &SchemaRef{}

	// Act
	got, err := json.Marshal(ref)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `null`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
