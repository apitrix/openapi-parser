package openapi31

import (
	"encoding/json"
	"testing"
)

func TestHeader_MarshalJSON_AllFields(t *testing.T) {
	schemaRef := &SchemaRef{}
	schemaRef.SetValue(NewSchema(SchemaFields{Type: SchemaType{Single: "integer"}}))
	h := NewHeader(HeaderFields{
		Description: "Rate limit",
		Required:    true,
		Schema:      schemaRef,
	})
	got, err := json.Marshal(h)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"description", "required", "schema"} {
		if _, ok := result[key]; !ok {
			t.Errorf("expected %q key", key)
		}
	}
}

func TestHeader_MarshalJSON_Empty(t *testing.T) {
	h := NewHeader(HeaderFields{})
	got, err := json.Marshal(h)
	if err != nil {
		t.Fatal(err)
	}
	want := `{}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
