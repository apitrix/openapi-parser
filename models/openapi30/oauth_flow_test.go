package openapi30

import (
	"encoding/json"
	"testing"
)

func TestOAuthFlow_MarshalJSON_AllFields(t *testing.T) {
	// Arrange
	f := NewOAuthFlow("https://auth.example.com", "https://token.example.com", "https://refresh.example.com", map[string]string{"read": "Read access"})

	// Act
	got, err := json.Marshal(f)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"authorizationUrl":"https://auth.example.com","tokenUrl":"https://token.example.com","refreshUrl":"https://refresh.example.com","scopes":{"read":"Read access"}}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestOAuthFlow_MarshalJSON_OmitsEmpty(t *testing.T) {
	// Arrange
	f := NewOAuthFlow("", "https://token.example.com", "", nil)

	// Act
	got, err := json.Marshal(f)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"tokenUrl":"https://token.example.com"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
