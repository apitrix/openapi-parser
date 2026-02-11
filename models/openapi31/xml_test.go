package openapi31

import (
	"encoding/json"
	"testing"
)

func TestXML_MarshalJSON_AllFields(t *testing.T) {
	x := NewXML("animal", "urn:example", "ex", true, true)
	got, err := json.Marshal(x)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"name":"animal","namespace":"urn:example","prefix":"ex","attribute":true,"wrapped":true}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestXML_MarshalJSON_NameOnly(t *testing.T) {
	x := NewXML("animal", "", "", false, false)
	got, err := json.Marshal(x)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"name":"animal"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestXML_MarshalJSON_Empty(t *testing.T) {
	x := NewXML("", "", "", false, false)
	got, err := json.Marshal(x)
	if err != nil {
		t.Fatal(err)
	}
	want := `{}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
