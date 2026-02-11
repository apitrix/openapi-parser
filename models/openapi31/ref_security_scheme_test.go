package openapi31

import (
	"encoding/json"
	"testing"
)

func TestSecuritySchemeRef_MarshalJSON_Ref(t *testing.T) {
	ref := NewSecuritySchemeRef("#/components/securitySchemes/BearerAuth")
	got, err := json.Marshal(ref)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"$ref":"#/components/securitySchemes/BearerAuth"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestSecuritySchemeRef_MarshalJSON_InlineValue(t *testing.T) {
	ss := NewSecurityScheme("http", "", "", "", "bearer", "JWT", "", nil)
	ref := &SecuritySchemeRef{Value: ss}
	got, err := json.Marshal(ref)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if _, ok := result["type"]; !ok {
		t.Error("expected 'type' key from inline security scheme")
	}
	if _, ok := result["scheme"]; !ok {
		t.Error("expected 'scheme' key from inline security scheme")
	}
}

func TestSecuritySchemeRef_MarshalJSON_NilValue(t *testing.T) {
	ref := &SecuritySchemeRef{}
	got, err := json.Marshal(ref)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != `null` {
		t.Errorf("got %s, want null", got)
	}
}
