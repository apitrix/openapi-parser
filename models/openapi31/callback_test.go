package openapi31

import (
	"encoding/json"
	"testing"
)

func TestCallback_MarshalJSON_WithPaths(t *testing.T) {
	pi := NewPathItem()
	pi.SetProperty("summary", "Webhook")
	cb := NewCallback(map[string]*PathItem{
		"{$request.body#/callbackUrl}": pi,
	})
	got, err := json.Marshal(cb)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if _, ok := result["{$request.body#/callbackUrl}"]; !ok {
		t.Error("expected callback expression key")
	}
}

func TestCallback_MarshalJSON_Empty(t *testing.T) {
	cb := NewCallback(nil)
	got, err := json.Marshal(cb)
	if err != nil {
		t.Fatal(err)
	}
	want := `{}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestCallback_MarshalJSON_WithExtensions(t *testing.T) {
	cb := NewCallback(nil)
	cb.VendorExtensions = map[string]interface{}{"x-custom": "val"}
	got, err := json.Marshal(cb)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"x-custom":"val"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
