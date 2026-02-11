package openapi31

import (
	"encoding/json"
	"testing"
)

func TestLinkRef_MarshalJSON_Ref(t *testing.T) {
	ref := NewLinkRef("#/components/links/GetUser")
	got, err := json.Marshal(ref)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"$ref":"#/components/links/GetUser"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestLinkRef_MarshalJSON_InlineValue(t *testing.T) {
	l := NewLink("", "getUser", "Get a user by ID", nil, nil, nil)
	ref := &LinkRef{Value: l}
	got, err := json.Marshal(ref)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if _, ok := result["operationId"]; !ok {
		t.Error("expected 'operationId' key from inline link")
	}
}

func TestLinkRef_MarshalJSON_NilValue(t *testing.T) {
	ref := &LinkRef{}
	got, err := json.Marshal(ref)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != `null` {
		t.Errorf("got %s, want null", got)
	}
}
