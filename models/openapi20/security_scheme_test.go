package openapi20

import (
	"encoding/json"
	"testing"
)

func TestSecurityScheme_MarshalJSON_APIKey(t *testing.T) {
	// Arrange
	ss := NewSecurityScheme("apiKey", "API key auth", "X-API-Key", "header", "", "", "", nil)

	// Act
	got, err := json.Marshal(ss)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := `{"type":"apiKey","description":"API key auth","name":"X-API-Key","in":"header"}`
	if string(got) != expected {
		t.Errorf("got %s, want %s", got, expected)
	}
}

func TestSecurityScheme_MarshalJSON_OAuth2(t *testing.T) {
	// Arrange
	scopes := map[string]string{"read:pets": "Read pets", "write:pets": "Write pets"}
	ss := NewSecurityScheme("oauth2", "", "", "", "implicit", "https://example.com/auth", "", scopes)

	// Act
	got, err := json.Marshal(ss)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// Verify key fields present
	var m map[string]interface{}
	if err := json.Unmarshal(got, &m); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if m["type"] != "oauth2" {
		t.Errorf("type = %v, want oauth2", m["type"])
	}
	if m["flow"] != "implicit" {
		t.Errorf("flow = %v, want implicit", m["flow"])
	}
	if m["authorizationUrl"] != "https://example.com/auth" {
		t.Errorf("authorizationUrl = %v, want https://example.com/auth", m["authorizationUrl"])
	}
}

func TestSecurityScheme_MarshalJSON_Basic(t *testing.T) {
	// Arrange
	ss := NewSecurityScheme("basic", "Basic HTTP auth", "", "", "", "", "", nil)

	// Act
	got, err := json.Marshal(ss)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := `{"type":"basic","description":"Basic HTTP auth"}`
	if string(got) != expected {
		t.Errorf("got %s, want %s", got, expected)
	}
}
