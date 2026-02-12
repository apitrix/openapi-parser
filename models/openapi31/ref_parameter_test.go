package openapi31

import (
	"encoding/json"
	"testing"
)

func TestParameterRef_MarshalJSON_Ref(t *testing.T) {
	ref := NewParameterRef("#/components/parameters/PageSize")
	got, err := json.Marshal(ref)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"$ref":"#/components/parameters/PageSize"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestParameterRef_MarshalJSON_InlineValue(t *testing.T) {
	p := NewParameter(ParameterFields{Name: "limit", In: "query"})
	ref := &ParameterRef{}
	ref.SetValue(p)
	got, err := json.Marshal(ref)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if _, ok := result["name"]; !ok {
		t.Error("expected 'name' key from inline parameter")
	}
}

func TestParameterRef_MarshalJSON_NilValue(t *testing.T) {
	ref := &ParameterRef{}
	got, err := json.Marshal(ref)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != `null` {
		t.Errorf("got %s, want null", got)
	}
}
