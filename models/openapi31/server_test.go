package openapi31

import (
	"encoding/json"
	"testing"
)

func TestServer_MarshalJSON_AllFields(t *testing.T) {
	sv := NewServerVariable([]string{"v1", "v2"}, "v1", "API version")
	s := NewServer("https://api.example.com", "Production",
		map[string]*ServerVariable{"version": sv})
	got, err := json.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"url", "description", "variables"} {
		if _, ok := result[key]; !ok {
			t.Errorf("expected %q key", key)
		}
	}
}

func TestServer_MarshalJSON_URLOnly(t *testing.T) {
	s := NewServer("https://api.example.com", "", nil)
	got, err := json.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"url":"https://api.example.com"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
