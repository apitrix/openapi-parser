package openapi31

import (
	"encoding/json"
	"testing"
)

func TestResponse_MarshalJSON_AllFields(t *testing.T) {
	mt := NewMediaType(nil, "response body", nil, nil)
	r := NewResponse("Success", nil, map[string]*MediaType{"application/json": mt}, nil)
	got, err := json.Marshal(r)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"description", "content"} {
		if _, ok := result[key]; !ok {
			t.Errorf("expected %q key", key)
		}
	}
}

func TestResponse_MarshalJSON_DescriptionOnly(t *testing.T) {
	r := NewResponse("Not found", nil, nil, nil)
	got, err := json.Marshal(r)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"description":"Not found"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestResponse_MarshalJSON_Empty(t *testing.T) {
	r := NewResponse("", nil, nil, nil)
	got, err := json.Marshal(r)
	if err != nil {
		t.Fatal(err)
	}
	want := `{}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
