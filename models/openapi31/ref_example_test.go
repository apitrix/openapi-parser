package openapi31

import (
	"encoding/json"
	"testing"
)

func TestExampleRef_MarshalJSON_Ref(t *testing.T) {
	ref := NewExampleRef("#/components/examples/Pet")
	got, err := json.Marshal(ref)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"$ref":"#/components/examples/Pet"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestExampleRef_MarshalJSON_InlineValue(t *testing.T) {
	ex := NewExample("A pet", "", "Rex", "")
	ref := &ExampleRef{Value: ex}
	got, err := json.Marshal(ref)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if _, ok := result["summary"]; !ok {
		t.Error("expected 'summary' key from inline example")
	}
}

func TestExampleRef_MarshalJSON_NilValue(t *testing.T) {
	ref := &ExampleRef{}
	got, err := json.Marshal(ref)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != `null` {
		t.Errorf("got %s, want null", got)
	}
}
