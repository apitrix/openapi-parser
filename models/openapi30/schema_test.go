package openapi30

import (
	"openapi-parser/models/shared"
	"encoding/json"
	"testing"
)

func TestSchema_MarshalJSON_BasicType(t *testing.T) {
	// Arrange
	s := NewSchema(SchemaFields{Type: "string", Description: "A name", Format: "email"})

	// Act
	got, err := json.Marshal(s)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"type":"string","description":"A name","format":"email"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestSchema_MarshalJSON_AdditionalPropertiesAsBoolean(t *testing.T) {
	// Arrange
	allowed := false
	s := NewSchema(SchemaFields{Type: "object", AdditionalPropertiesAllowed: &allowed})

	// Act
	got, err := json.Marshal(s)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	// false bool for additionalProperties should NOT be omitted because it's in a *bool context
	want := `{"type":"object","additionalProperties":false}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestSchema_MarshalJSON_AdditionalPropertiesAsSchema(t *testing.T) {
	// Arrange
	propSchema := &shared.Ref[Schema]{}
	propSchema.SetValue(NewSchema(SchemaFields{Type: "string"}))
	s := NewSchema(SchemaFields{Type: "object", AdditionalProperties: propSchema})

	// Act
	got, err := json.Marshal(s)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if _, ok := result["additionalProperties"]; !ok {
		t.Error("expected 'additionalProperties' key")
	}
	if _, ok := result["type"]; !ok {
		t.Error("expected 'type' key")
	}
}

func TestSchema_MarshalJSON_Required(t *testing.T) {
	// Arrange
	s := NewSchema(SchemaFields{Type: "object", Required: []string{"name", "age"}})

	// Act
	got, err := json.Marshal(s)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"required":["name","age"],"type":"object"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestSchema_MarshalJSON_Nullable(t *testing.T) {
	// Arrange
	s := NewSchema(SchemaFields{Type: "string", Nullable: true})

	// Act
	got, err := json.Marshal(s)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"type":"string","nullable":true}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestSchema_MarshalJSON_AllOf(t *testing.T) {
	// Arrange
	ref1 := shared.NewRef[Schema]("#/components/schemas/Base")
	ref2 := shared.NewRef[Schema]("#/components/schemas/Extended")
	s := NewSchema(SchemaFields{AllOf: []*shared.Ref[Schema]{ref1, ref2}})

	// Act
	got, err := json.Marshal(s)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"allOf":[{"$ref":"#/components/schemas/Base"},{"$ref":"#/components/schemas/Extended"}]}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestSchema_MarshalJSON_EmptySchema(t *testing.T) {
	// Arrange
	s := NewSchema(SchemaFields{})

	// Act
	got, err := json.Marshal(s)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestSchema_MarshalJSON_WithExtensions(t *testing.T) {
	// Arrange
	s := NewSchema(SchemaFields{Type: "string"})
	s.VendorExtensions = map[string]interface{}{"x-go-type": "MyString"}

	// Act
	got, err := json.Marshal(s)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"type":"string","x-go-type":"MyString"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestSchema_MarshalJSON_NumericConstraints(t *testing.T) {
	// Arrange
	max := 100.0
	min := 0.0
	mult := 5.0
	s := NewSchema(SchemaFields{
		Type:       "integer",
		Maximum:    &max,
		Minimum:    &min,
		MultipleOf: &mult,
	})

	// Act
	got, err := json.Marshal(s)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"multipleOf":5,"maximum":100,"minimum":0,"type":"integer"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
