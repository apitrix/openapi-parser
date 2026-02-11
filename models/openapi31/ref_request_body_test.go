package openapi31

import (
	"encoding/json"
	"testing"
)

func TestRequestBodyRef_MarshalJSON_Ref(t *testing.T) {
	ref := NewRequestBodyRef("#/components/requestBodies/PetBody")
	got, err := json.Marshal(ref)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"$ref":"#/components/requestBodies/PetBody"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestRequestBodyRef_MarshalJSON_InlineValue(t *testing.T) {
	rb := NewRequestBody("A pet body", nil, true)
	ref := &RequestBodyRef{Value: rb}
	got, err := json.Marshal(ref)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if _, ok := result["description"]; !ok {
		t.Error("expected 'description' key from inline request body")
	}
}

func TestRequestBodyRef_MarshalJSON_NilValue(t *testing.T) {
	ref := &RequestBodyRef{}
	got, err := json.Marshal(ref)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != `null` {
		t.Errorf("got %s, want null", got)
	}
}
