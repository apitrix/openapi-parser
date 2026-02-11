package openapi31

import (
	"encoding/json"
	"testing"
)

func TestResponses_MarshalJSON_WithDefault(t *testing.T) {
	defResp := &ResponseRef{Value: NewResponse("Default error", nil, nil, nil)}
	resp := NewResponses(defResp, nil)
	got, err := json.Marshal(resp)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if _, ok := result["default"]; !ok {
		t.Error("expected 'default' key")
	}
}

func TestResponses_MarshalJSON_WithCodes(t *testing.T) {
	code200 := &ResponseRef{Value: NewResponse("Success", nil, nil, nil)}
	code404 := &ResponseRef{Value: NewResponse("Not found", nil, nil, nil)}
	resp := NewResponses(nil, map[string]*ResponseRef{
		"200": code200,
		"404": code404,
	})
	got, err := json.Marshal(resp)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"200", "404"} {
		if _, ok := result[key]; !ok {
			t.Errorf("expected %q key", key)
		}
	}
}

func TestResponses_MarshalJSON_Empty(t *testing.T) {
	resp := NewResponses(nil, nil)
	got, err := json.Marshal(resp)
	if err != nil {
		t.Fatal(err)
	}
	want := `{}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
