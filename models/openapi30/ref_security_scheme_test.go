package openapi30

import (
	"encoding/json"
	"testing"
)

func TestSecuritySchemeRef_MarshalJSON_Ref(t *testing.T) {
	// Arrange
	ref := NewSecuritySchemeRef("#/components/securitySchemes/BearerAuth")

	// Act
	got, err := json.Marshal(ref)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"$ref":"#/components/securitySchemes/BearerAuth"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestSecuritySchemeRef_MarshalJSON_InlineValue(t *testing.T) {
	// Arrange
	ss := NewSecurityScheme("http", "", "", "", "bearer", "JWT", nil, "")
	ref := &SecuritySchemeRef{}
	ref.SetValue(ss)

	// Act
	got, err := json.Marshal(ref)

	// Assert
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
	if _, ok := result["scheme"]; !ok {
		t.Error("expected 'scheme' key")
	}
}
