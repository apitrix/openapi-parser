package openapi31

import (
	"encoding/json"
	"testing"
)

func TestEncoding_MarshalJSON_AllFields(t *testing.T) {
	explode := true
	e := NewEncoding("application/json", "form", nil, &explode, true)
	got, err := json.Marshal(e)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"contentType":"application/json","style":"form","explode":true,"allowReserved":true}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestEncoding_MarshalJSON_ContentTypeOnly(t *testing.T) {
	e := NewEncoding("application/xml", "", nil, nil, false)
	got, err := json.Marshal(e)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"contentType":"application/xml"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestEncoding_MarshalJSON_Empty(t *testing.T) {
	e := NewEncoding("", "", nil, nil, false)
	got, err := json.Marshal(e)
	if err != nil {
		t.Fatal(err)
	}
	want := `{}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
