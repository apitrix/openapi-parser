package openapi31

import (
	"encoding/json"
	"testing"
)

func TestLicense_MarshalJSON_AllFields(t *testing.T) {
	l := NewLicense("MIT", "MIT", "https://opensource.org/licenses/MIT")
	got, err := json.Marshal(l)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"name":"MIT","identifier":"MIT","url":"https://opensource.org/licenses/MIT"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestLicense_MarshalJSON_IdentifierOnly(t *testing.T) {
	l := NewLicense("Apache-2.0", "Apache-2.0", "")
	got, err := json.Marshal(l)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"name":"Apache-2.0","identifier":"Apache-2.0"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestLicense_MarshalJSON_OmitsEmpty(t *testing.T) {
	l := NewLicense("MIT", "", "")
	got, err := json.Marshal(l)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"name":"MIT"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
