package openapi31

import (
	"encoding/json"
	"testing"
)

func TestTag_MarshalJSON_AllFields(t *testing.T) {
	ed := NewExternalDocumentation("https://docs.example.com", "Tag docs")
	tag := NewTag("pets", "Pet operations", ed)
	got, err := json.Marshal(tag)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"name", "description", "externalDocs"} {
		if _, ok := result[key]; !ok {
			t.Errorf("expected %q key", key)
		}
	}
}

func TestTag_MarshalJSON_NameOnly(t *testing.T) {
	tag := NewTag("pets", "", nil)
	got, err := json.Marshal(tag)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"name":"pets"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
