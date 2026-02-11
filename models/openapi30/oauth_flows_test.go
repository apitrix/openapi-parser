package openapi30

import (
	"encoding/json"
	"testing"
)

func TestOAuthFlows_MarshalJSON_AllFields(t *testing.T) {
	// Arrange
	implicit := NewOAuthFlow("https://auth.example.com", "", "", nil)
	password := NewOAuthFlow("", "https://token.example.com", "", nil)
	flows := NewOAuthFlows(implicit, password, nil, nil)

	// Act
	got, err := json.Marshal(flows)

	// Assert
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
		t.Error("unexpected 'clientCredentials' key (should be omitted)")
	}
}

func TestOAuthFlows_MarshalJSON_Empty(t *testing.T) {
	// Arrange
	flows := NewOAuthFlows(nil, nil, nil, nil)

	// Act
	got, err := json.Marshal(flows)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
