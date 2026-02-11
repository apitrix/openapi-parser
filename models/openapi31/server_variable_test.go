package openapi31

import (
	"encoding/json"
	"testing"
)

func TestServerVariable_MarshalJSON_AllFields(t *testing.T) {
	sv := NewServerVariable([]string{"v1", "v2"}, "v1", "API version")
	got, err := json.Marshal(sv)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"enum":["v1","v2"],"default":"v1","description":"API version"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestServerVariable_MarshalJSON_DefaultOnly(t *testing.T) {
	sv := NewServerVariable(nil, "v1", "")
	got, err := json.Marshal(sv)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"default":"v1"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
