package openapi31

import (
	"encoding/json"
	"testing"
)

func TestPathItem_MarshalJSON_WithOperations(t *testing.T) {
	pi := NewPathItem()
	getOp := NewOperation()
	getOp.SetProperty("summary", "Get pet")
	getOp.SetProperty("operationId", "getPet")
	postOp := NewOperation()
	postOp.SetProperty("summary", "Create pet")
	pi.SetProperty("summary", "Pet operations")
	pi.SetProperty("get", getOp)
	pi.SetProperty("post", postOp)
	got, err := json.Marshal(pi)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"summary", "get", "post"} {
		if _, ok := result[key]; !ok {
			t.Errorf("expected %q key", key)
		}
	}
	// Unused HTTP methods should be omitted
	for _, key := range []string{"put", "delete", "options", "head", "patch", "trace"} {
		if _, ok := result[key]; ok {
			t.Errorf("expected %q to be omitted", key)
		}
	}
}

func TestPathItem_MarshalJSON_WithRef(t *testing.T) {
	pi := NewPathItem()
	pi.SetProperty("$ref", "#/components/pathItems/Shared")
	got, err := json.Marshal(pi)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"$ref":"#/components/pathItems/Shared"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestPathItem_MarshalJSON_Empty(t *testing.T) {
	pi := NewPathItem()
	got, err := json.Marshal(pi)
	if err != nil {
		t.Fatal(err)
	}
	want := `{}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
