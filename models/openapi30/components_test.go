package openapi30

import (
	"openapi-parser/models/shared"
	"encoding/json"
	"testing"
)

func TestComponents_MarshalJSON_WithSchemas(t *testing.T) {
	// Arrange
	petRef := &shared.Ref[Schema]{}
	petRef.SetValue(NewSchema(SchemaFields{Type: "object"}))
	schemas := map[string]*shared.Ref[Schema]{
		"Pet": petRef,
	}
	c := NewComponents(schemas, nil, nil, nil, nil, nil, nil, nil, nil)

	// Act
	got, err := json.Marshal(c)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if _, ok := result["schemas"]; !ok {
		t.Error("expected 'schemas' key")
	}
	// Empty maps should be omitted
	for _, key := range []string{"responses", "parameters", "examples", "requestBodies", "headers", "securitySchemes", "links", "callbacks"} {
		if _, ok := result[key]; ok {
			t.Errorf("expected %q to be omitted (nil map)", key)
		}
	}
}

func TestComponents_MarshalJSON_Empty(t *testing.T) {
	// Arrange
	c := NewComponents(nil, nil, nil, nil, nil, nil, nil, nil, nil)

	// Act
	got, err := json.Marshal(c)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestComponents_MarshalJSON_Multiple(t *testing.T) {
	// Arrange
	petRef := &shared.Ref[Schema]{}
	petRef.SetValue(NewSchema(SchemaFields{Type: "object"}))
	schemas := map[string]*shared.Ref[Schema]{
		"Pet": petRef,
	}
	bearerRef := &shared.Ref[SecurityScheme]{}
	bearerRef.SetValue(NewSecurityScheme("http", "", "", "", "bearer", "", nil, ""))
	secSchemes := map[string]*shared.Ref[SecurityScheme]{
		"BearerAuth": bearerRef,
	}
	c := NewComponents(schemas, nil, nil, nil, nil, nil, secSchemes, nil, nil)

	// Act
	got, err := json.Marshal(c)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if _, ok := result["schemas"]; !ok {
		t.Error("expected 'schemas' key")
	}
	if _, ok := result["securitySchemes"]; !ok {
		t.Error("expected 'securitySchemes' key")
	}
}
