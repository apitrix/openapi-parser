package openapi31

import (
	"encoding/json"
	"testing"
)

func TestExternalDocumentation_MarshalJSON_AllFields(t *testing.T) {
	ed := NewExternalDocumentation("More docs", "https://docs.example.com")
	got, err := json.Marshal(ed)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"description":"More docs","url":"https://docs.example.com"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestExternalDocumentation_MarshalJSON_URLOnly(t *testing.T) {
	ed := NewExternalDocumentation("", "https://docs.example.com")
	got, err := json.Marshal(ed)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"url":"https://docs.example.com"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
