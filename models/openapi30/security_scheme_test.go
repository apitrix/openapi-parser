package openapi30

import (
	"encoding/json"
	"testing"
)

func TestSecurityScheme_MarshalJSON_HTTPBearer(t *testing.T) {
	// Arrange
	ss := NewSecurityScheme("http", "Bearer auth", "", "", "bearer", "JWT", nil, "")

	// Act
	got, err := json.Marshal(ss)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"type", "description", "scheme", "bearerFormat"} {
		if _, ok := result[key]; !ok {
			t.Errorf("expected %q key", key)
		}
	}
	// "type" JSON key maps from secType field
	if string(result["type"]) != `"http"` {
		t.Errorf("type = %s, want %q", result["type"], "http")
	}
}

func TestSecurityScheme_MarshalJSON_APIKey(t *testing.T) {
	// Arrange
	ss := NewSecurityScheme("apiKey", "", "X-API-Key", "header", "", "", nil, "")

	// Act
	got, err := json.Marshal(ss)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"type":"apiKey","name":"X-API-Key","in":"header"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
