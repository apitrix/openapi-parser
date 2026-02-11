package openapi31

import (
	"encoding/json"
	"testing"
)

func TestOperation_MarshalJSON_AllFields(t *testing.T) {
	op := NewOperation()
	op.SetProperty("tags", []string{"pets"})
	op.SetProperty("summary", "List pets")
	op.SetProperty("operationId", "listPets")
	op.SetProperty("deprecated", true)
	got, err := json.Marshal(op)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"tags", "summary", "operationId", "deprecated"} {
		if _, ok := result[key]; !ok {
			t.Errorf("expected %q key", key)
		}
	}
}

func TestOperation_MarshalJSON_Empty(t *testing.T) {
	op := NewOperation()
	got, err := json.Marshal(op)
	if err != nil {
		t.Fatal(err)
	}
	want := `{}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
