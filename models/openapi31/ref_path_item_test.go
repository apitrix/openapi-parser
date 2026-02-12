package openapi31

import (
	"encoding/json"
	"testing"
)

func TestPathItemRef_MarshalJSON_Ref(t *testing.T) {
	ref := NewPathItemRef("#/components/pathItems/SharedPath")
	got, err := json.Marshal(ref)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"$ref":"#/components/pathItems/SharedPath"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestPathItemRef_MarshalJSON_RefWithSummaryDescription(t *testing.T) {
	ref := NewPathItemRef("#/components/pathItems/SharedPath")
	ref.Summary = "Shared path"
	ref.Description = "A shared path item"
	got, err := json.Marshal(ref)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]string
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if result["$ref"] != "#/components/pathItems/SharedPath" {
		t.Error("expected $ref")
	}
	if result["summary"] != "Shared path" {
		t.Error("expected summary")
	}
	if result["description"] != "A shared path item" {
		t.Error("expected description")
	}
}

func TestPathItemRef_MarshalJSON_InlineValue(t *testing.T) {
	pi := NewPathItem()
	pi.SetProperty("summary", "Pet endpoint")
	ref := &PathItemRef{}
	ref.SetValue(pi)
	got, err := json.Marshal(ref)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if _, ok := result["summary"]; !ok {
		t.Error("expected 'summary' key from inline path item")
	}
}

func TestPathItemRef_MarshalJSON_NilValue(t *testing.T) {
	ref := &PathItemRef{}
	got, err := json.Marshal(ref)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != `null` {
		t.Errorf("got %s, want null", got)
	}
}
