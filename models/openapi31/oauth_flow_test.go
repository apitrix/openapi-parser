package openapi31

import (
	"encoding/json"
	"testing"
)

func TestOAuthFlow_MarshalJSON_AllFields(t *testing.T) {
	f := NewOAuthFlow("https://auth.example.com/authorize", "https://auth.example.com/token",
		"https://auth.example.com/refresh", map[string]string{"read:pets": "Read pets"})
	got, err := json.Marshal(f)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"authorizationUrl", "tokenUrl", "refreshUrl", "scopes"} {
		if _, ok := result[key]; !ok {
			t.Errorf("expected %q key", key)
		}
	}
}

func TestOAuthFlow_MarshalJSON_ScopesOnly(t *testing.T) {
	f := NewOAuthFlow("", "", "", map[string]string{"admin": "Admin access"})
	got, err := json.Marshal(f)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if _, ok := result["scopes"]; !ok {
		t.Error("expected 'scopes' key")
	}
}
