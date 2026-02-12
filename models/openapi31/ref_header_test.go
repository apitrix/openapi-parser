package openapi31

import (
	"encoding/json"
	"testing"
)

func TestHeaderRef_MarshalJSON_Ref(t *testing.T) {
	ref := NewHeaderRef("#/components/headers/X-Rate-Limit")
	got, err := json.Marshal(ref)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"$ref":"#/components/headers/X-Rate-Limit"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestHeaderRef_MarshalJSON_InlineValue(t *testing.T) {
	h := NewHeader(HeaderFields{Description: "Rate limit"})
	ref := &HeaderRef{}
	ref.SetValue(h)
	got, err := json.Marshal(ref)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if _, ok := result["description"]; !ok {
		t.Error("expected 'description' key from inline header")
	}
}

func TestHeaderRef_MarshalJSON_NilValue(t *testing.T) {
	ref := &HeaderRef{}
	got, err := json.Marshal(ref)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != `null` {
		t.Errorf("got %s, want null", got)
	}
}
