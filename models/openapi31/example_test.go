package openapi31

import (
	"encoding/json"
	"testing"
)

func TestExample_MarshalJSON_AllFields(t *testing.T) {
	e := NewExample("A pet example", "A dog named Rex", "Rex", "")
	got, err := json.Marshal(e)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"summary", "description", "value"} {
		if _, ok := result[key]; !ok {
			t.Errorf("expected %q key", key)
		}
	}
}

func TestExample_MarshalJSON_ExternalValue(t *testing.T) {
	e := NewExample("", "", nil, "https://example.com/sample.json")
	got, err := json.Marshal(e)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"externalValue":"https://example.com/sample.json"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestExample_MarshalJSON_Empty(t *testing.T) {
	e := NewExample("", "", nil, "")
	got, err := json.Marshal(e)
	if err != nil {
		t.Fatal(err)
	}
	want := `{}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
