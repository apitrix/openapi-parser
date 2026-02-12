package openapi31

import (
	"encoding/json"
	"testing"
)

func TestComponents_MarshalJSON_WithPathItems(t *testing.T) {
	// pathItems is 3.1-specific
	comp := NewComponents()
	comp.SetProperty("pathItems", map[string]*PathItemRef{
		"MyPath": {Ref: "#/components/pathItems/MyPath"},
	})
	got, err := json.Marshal(comp)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if _, ok := result["pathItems"]; !ok {
		t.Error("expected 'pathItems' key (3.1-specific)")
	}
}

func TestComponents_MarshalJSON_Empty(t *testing.T) {
	comp := NewComponents()
	got, err := json.Marshal(comp)
	if err != nil {
		t.Fatal(err)
	}
	want := `{}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestComponents_MarshalJSON_WithSchemas(t *testing.T) {
	comp := NewComponents()
	petRef := &SchemaRef{}
	petRef.SetValue(NewSchema(SchemaFields{Type: SchemaType{Single: "object"}}))
	comp.SetProperty("schemas", map[string]*SchemaRef{
		"Pet": petRef,
	})
	got, err := json.Marshal(comp)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if _, ok := result["schemas"]; !ok {
		t.Error("expected 'schemas' key")
	}
}
