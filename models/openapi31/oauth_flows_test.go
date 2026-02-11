package openapi31

import (
	"encoding/json"
	"testing"
)

func TestOAuthFlows_MarshalJSON_AllFlows(t *testing.T) {
	implicit := NewOAuthFlow("https://auth.example.com/authorize", "", "", map[string]string{"read": "Read"})
	password := NewOAuthFlow("", "https://auth.example.com/token", "", map[string]string{"write": "Write"})
	flows := NewOAuthFlows(implicit, password, nil, nil)
	got, err := json.Marshal(flows)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if _, ok := result["implicit"]; !ok {
		t.Error("expected 'implicit' key")
	}
	if _, ok := result["password"]; !ok {
		t.Error("expected 'password' key")
	}
	if _, ok := result["clientCredentials"]; ok {
		t.Error("expected 'clientCredentials' to be omitted")
	}
}

func TestOAuthFlows_MarshalJSON_Empty(t *testing.T) {
	flows := NewOAuthFlows(nil, nil, nil, nil)
	got, err := json.Marshal(flows)
	if err != nil {
		t.Fatal(err)
	}
	want := `{}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
