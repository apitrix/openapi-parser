package openapi31

import (
	"encoding/json"
	"testing"
)

func TestParameter_MarshalJSON_AllFields(t *testing.T) {
	schemaRef := &SchemaRef{}
	schemaRef.SetValue(NewSchema(SchemaFields{Type: SchemaType{Single: "integer"}}))
	p := NewParameter(ParameterFields{
		Name:        "limit",
		In:          "query",
		Description: "Max items to return",
		Required:    true,
		Schema:      schemaRef,
	})
	got, err := json.Marshal(p)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"name", "in", "description", "required", "schema"} {
		if _, ok := result[key]; !ok {
			t.Errorf("expected %q key", key)
		}
	}
}

func TestParameter_MarshalJSON_Minimal(t *testing.T) {
	p := NewParameter(ParameterFields{Name: "id", In: "path"})
	got, err := json.Marshal(p)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"name":"id","in":"path"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestParameter_MarshalJSON_Empty(t *testing.T) {
	p := NewParameter(ParameterFields{})
	got, err := json.Marshal(p)
	if err != nil {
		t.Fatal(err)
	}
	want := `{}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
