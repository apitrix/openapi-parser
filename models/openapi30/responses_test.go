package openapi30

import (
	"encoding/json"
	"testing"
)

func TestResponses_MarshalJSON_DefaultFirst(t *testing.T) {
	// Arrange
	defaultResp := &RefResponse{}
	defaultResp.SetValue(NewResponse("Unexpected error", nil, nil, nil))
	ok200 := &RefResponse{}
	ok200.SetValue(NewResponse("OK", nil, nil, nil))
	nf404 := &RefResponse{}
	nf404.SetValue(NewResponse("Not Found", nil, nil, nil))
	codes := map[string]*RefResponse{
		"200": ok200,
		"404": nf404,
	}
	r := NewResponses(defaultResp, codes)

	// Act
	got, err := json.Marshal(r)

	// Assert
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
	if _, ok := result["200"]; !ok {
		t.Error("expected '200' key")
	}
	if _, ok := result["404"]; !ok {
		t.Error("expected '404' key")
	}
}

func TestResponses_MarshalJSON_NoDefault(t *testing.T) {
	// Arrange
	ok200 := &RefResponse{}
	ok200.SetValue(NewResponse("OK", nil, nil, nil))
	codes := map[string]*RefResponse{
		"200": ok200,
	}
	r := NewResponses(nil, codes)

	// Act
	got, err := json.Marshal(r)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if _, ok := result["default"]; ok {
		t.Error("unexpected 'default' key")
	}
	if _, ok := result["200"]; !ok {
		t.Error("expected '200' key")
	}
}

func TestResponses_MarshalJSON_Empty(t *testing.T) {
	// Arrange
	r := NewResponses(nil, nil)

	// Act
	got, err := json.Marshal(r)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
