package openapi31

import (
	"encoding/json"
	"testing"
)

func TestRequestBody_MarshalJSON_AllFields(t *testing.T) {
	mt := NewMediaType(nil, "body", nil, nil)
	rb := NewRequestBody("A request body", map[string]*MediaType{"application/json": mt}, true)
	got, err := json.Marshal(rb)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"description", "content", "required"} {
		if _, ok := result[key]; !ok {
			t.Errorf("expected %q key", key)
		}
	}
}

func TestRequestBody_MarshalJSON_DescriptionOnly(t *testing.T) {
	rb := NewRequestBody("A body", nil, false)
	got, err := json.Marshal(rb)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"description":"A body"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestRequestBody_MarshalJSON_Empty(t *testing.T) {
	rb := NewRequestBody("", nil, false)
	got, err := json.Marshal(rb)
	if err != nil {
		t.Fatal(err)
	}
	want := `{}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
