package openapi31

import (
	"encoding/json"
	"testing"
)

func TestSecurityScheme_MarshalJSON_HTTP(t *testing.T) {
	ss := NewSecurityScheme("http", "Bearer auth", "", "", "bearer", "JWT", "", nil)
	got, err := json.Marshal(ss)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if _, ok := result["type"]; !ok {
		t.Error("expected 'type' key (mapped from schemeType)")
	}
	if _, ok := result["scheme"]; !ok {
		t.Error("expected 'scheme' key")
	}
	if _, ok := result["bearerFormat"]; !ok {
		t.Error("expected 'bearerFormat' key")
	}
}

func TestSecurityScheme_MarshalJSON_OpenIDConnect(t *testing.T) {
	ss := NewSecurityScheme("openIdConnect", "", "", "", "", "", "https://auth.example.com/.well-known/openid-configuration", nil)
	got, err := json.Marshal(ss)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if _, ok := result["type"]; !ok {
		t.Error("expected 'type' key")
	}
	if _, ok := result["openIdConnectUrl"]; !ok {
		t.Error("expected 'openIdConnectUrl' key (mapped from openIDConnectURL)")
	}
}

func TestSecurityScheme_MarshalJSON_APIKey(t *testing.T) {
	ss := NewSecurityScheme("apiKey", "", "X-API-Key", "header", "", "", "", nil)
	got, err := json.Marshal(ss)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"type", "name", "in"} {
		if _, ok := result[key]; !ok {
			t.Errorf("expected %q key", key)
		}
	}
}
