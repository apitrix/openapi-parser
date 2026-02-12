package openapi31

import (
	"encoding/json"
	"testing"
)

func TestMediaType_MarshalJSON_WithSchema(t *testing.T) {
	schema := &SchemaRef{}
	schema.SetValue(NewSchema(SchemaFields{Type: SchemaType{Single: "object"}}))
	mt := NewMediaType(schema, nil, nil, nil)
	got, err := json.Marshal(mt)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if _, ok := result["schema"]; !ok {
		t.Error("expected 'schema' key")
	}
}

func TestMediaType_MarshalJSON_WithExample(t *testing.T) {
	mt := NewMediaType(nil, "sample value", nil, nil)
	got, err := json.Marshal(mt)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if _, ok := result["example"]; !ok {
		t.Error("expected 'example' key")
	}
}

func TestMediaType_MarshalJSON_Empty(t *testing.T) {
	mt := NewMediaType(nil, nil, nil, nil)
	got, err := json.Marshal(mt)
	if err != nil {
		t.Fatal(err)
	}
	want := `{}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
