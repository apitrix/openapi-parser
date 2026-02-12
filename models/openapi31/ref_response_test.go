package openapi31

import (
	"encoding/json"
	"testing"
)

func TestResponseRef_MarshalJSON_Ref(t *testing.T) {
	ref := NewResponseRef("#/components/responses/NotFound")
	got, err := json.Marshal(ref)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"$ref":"#/components/responses/NotFound"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestResponseRef_MarshalJSON_InlineValue(t *testing.T) {
	r := NewResponse("Not found", nil, nil, nil)
	ref := &ResponseRef{}
	ref.SetValue(r)
	got, err := json.Marshal(ref)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if _, ok := result["description"]; !ok {
		t.Error("expected 'description' key from inline response")
	}
}

func TestResponseRef_MarshalJSON_NilValue(t *testing.T) {
	ref := &ResponseRef{}
	got, err := json.Marshal(ref)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != `null` {
		t.Errorf("got %s, want null", got)
	}
}
