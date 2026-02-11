package openapi31

import (
	"encoding/json"
	"testing"
)

func TestCallbackRef_MarshalJSON_Ref(t *testing.T) {
	ref := NewCallbackRef("#/components/callbacks/onPetCreate")
	got, err := json.Marshal(ref)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"$ref":"#/components/callbacks/onPetCreate"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestCallbackRef_MarshalJSON_RefWithSummary(t *testing.T) {
	ref := NewCallbackRef("#/components/callbacks/onPetCreate")
	ref.Summary = "Pet callback"
	got, err := json.Marshal(ref)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]string
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if result["$ref"] != "#/components/callbacks/onPetCreate" {
		t.Error("expected $ref")
	}
	if result["summary"] != "Pet callback" {
		t.Error("expected summary")
	}
}

func TestCallbackRef_MarshalJSON_InlineValue(t *testing.T) {
	pi := NewPathItem()
	pi.SetProperty("summary", "Webhook")
	cb := NewCallback(map[string]*PathItem{"{$request.body#/callbackUrl}": pi})
	ref := &CallbackRef{Value: cb}
	got, err := json.Marshal(ref)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if _, ok := result["{$request.body#/callbackUrl}"]; !ok {
		t.Error("expected callback expression key")
	}
}

func TestCallbackRef_MarshalJSON_NilValue(t *testing.T) {
	ref := &CallbackRef{}
	got, err := json.Marshal(ref)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != `null` {
		t.Errorf("got %s, want null", got)
	}
}
