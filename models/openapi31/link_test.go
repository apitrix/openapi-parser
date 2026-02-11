package openapi31

import (
	"encoding/json"
	"testing"
)

func TestLink_MarshalJSON_OperationId(t *testing.T) {
	l := NewLink("", "getUser", "Get a user", map[string]interface{}{"userId": "$response.body#/id"}, nil, nil)
	got, err := json.Marshal(l)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if _, ok := result["operationId"]; !ok {
		t.Error("expected 'operationId' key")
	}
	if _, ok := result["description"]; !ok {
		t.Error("expected 'description' key")
	}
	if _, ok := result["parameters"]; !ok {
		t.Error("expected 'parameters' key")
	}
}

func TestLink_MarshalJSON_OperationRef(t *testing.T) {
	l := NewLink("#/paths/~1users~1{id}/get", "", "", nil, nil, nil)
	got, err := json.Marshal(l)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"operationRef":"#/paths/~1users~1{id}/get"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestLink_MarshalJSON_Empty(t *testing.T) {
	l := NewLink("", "", "", nil, nil, nil)
	got, err := json.Marshal(l)
	if err != nil {
		t.Fatal(err)
	}
	want := `{}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
