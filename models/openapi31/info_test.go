package openapi31

import (
	"encoding/json"
	"testing"
)

func TestInfo_MarshalJSON_AllFields(t *testing.T) {
	c := NewContact("John", "", "")
	l := NewLicense("MIT", "", "")
	i := NewInfo("My API", "A summary", "A description", "https://tos.example.com", "1.0.0", c, l)
	got, err := json.Marshal(i)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"title", "summary", "description", "termsOfService", "contact", "license", "version"} {
		if _, ok := result[key]; !ok {
			t.Errorf("expected %q key", key)
		}
	}
}

func TestInfo_MarshalJSON_OmitsEmpty(t *testing.T) {
	i := NewInfo("My API", "", "", "", "1.0.0", nil, nil)
	got, err := json.Marshal(i)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"title":"My API","version":"1.0.0"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestInfo_MarshalJSON_SummaryField(t *testing.T) {
	// summary is 3.1-specific
	i := NewInfo("My API", "Brief summary", "", "", "1.0.0", nil, nil)
	got, err := json.Marshal(i)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"title":"My API","summary":"Brief summary","version":"1.0.0"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
