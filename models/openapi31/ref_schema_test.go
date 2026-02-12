package openapi31

import (
	"encoding/json"
	"testing"
)

func TestSchemaRef_MarshalJSON_Ref(t *testing.T) {
	ref := NewSchemaRef("#/components/schemas/Pet")
	got, err := json.Marshal(ref)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"$ref":"#/components/schemas/Pet"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestSchemaRef_MarshalJSON_RefWithSummaryDescription(t *testing.T) {
	ref := NewSchemaRef("#/components/schemas/Pet")
	ref.Summary = "A pet"
	ref.Description = "A pet in the store"
	got, err := json.Marshal(ref)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]string
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if result["$ref"] != "#/components/schemas/Pet" {
		t.Error("expected $ref")
	}
	if result["summary"] != "A pet" {
		t.Error("expected summary")
	}
	if result["description"] != "A pet in the store" {
		t.Error("expected description")
	}
}

func TestSchemaRef_MarshalJSON_InlineValue(t *testing.T) {
	schema := NewSchema(SchemaFields{Type: SchemaType{Single: "string"}, Description: "A name"})
	ref := &SchemaRef{}
	ref.SetValue(schema)
	got, err := json.Marshal(ref)
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
	ref := &SchemaRef{}
	got, err := json.Marshal(ref)
	if err != nil {
		t.Fatal(err)
	}
	want := `null`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
