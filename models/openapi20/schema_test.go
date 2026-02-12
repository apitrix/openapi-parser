package openapi20

import (
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
		t.Fatalf("unexpected error: %v", err)
	}
	expected := `{"type":"string","format":"email","description":"A name"}`
	if string(got) != expected {
		t.Errorf("got %s, want %s", got, expected)
	}
}

func TestSchema_MarshalJSON_AdditionalPropertiesBool(t *testing.T) {
	// Arrange
	f := false
	s := NewSchema(SchemaFields{
		Type:                        "object",
		AdditionalPropertiesAllowed: &f,
	})

	// Act
	got, err := json.Marshal(s)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := `{"type":"object","additionalProperties":false}`
	if string(got) != expected {
		t.Errorf("got %s, want %s", got, expected)
	}
}

func TestSchema_MarshalJSON_AdditionalPropertiesSchema(t *testing.T) {
	// Arrange
	addProps := &SchemaRef{}
	addProps.SetValue(NewSchema(SchemaFields{Type: "string"}))
	s := NewSchema(SchemaFields{
		Type:                 "object",
		AdditionalProperties: addProps,
	})

	// Act
	got, err := json.Marshal(s)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := `{"type":"object","additionalProperties":{"type":"string"}}`
	if string(got) != expected {
		t.Errorf("got %s, want %s", got, expected)
	}
}

func TestSchema_MarshalJSON_AllOf(t *testing.T) {
	// Arrange
	inlineRef := &SchemaRef{}
	inlineRef.SetValue(NewSchema(SchemaFields{Type: "object"}))
	s := NewSchema(SchemaFields{
		AllOf: []*SchemaRef{
			NewSchemaRef("#/definitions/Base"),
			inlineRef,
		},
	})

	// Act
	got, err := json.Marshal(s)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := `{"allOf":[{"$ref":"#/definitions/Base"},{"type":"object"}]}`
	if string(got) != expected {
		t.Errorf("got %s, want %s", got, expected)
	}
}

func TestSchema_MarshalJSON_WithNumericConstraints(t *testing.T) {
	// Arrange
	max := float64(100)
	min := float64(1)
	s := NewSchema(SchemaFields{
		Type:    "integer",
		Maximum: &max,
		Minimum: &min,
	})

	// Act
	got, err := json.Marshal(s)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var m map[string]interface{}
	if err := json.Unmarshal(got, &m); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if m["maximum"] != float64(100) {
		t.Errorf("maximum = %v, want 100", m["maximum"])
	}
	if m["minimum"] != float64(1) {
		t.Errorf("minimum = %v, want 1", m["minimum"])
	}
}

func TestSchema_MarshalJSON_WithDiscriminator(t *testing.T) {
	// Arrange
	s := NewSchema(SchemaFields{
		Type:          "object",
		Discriminator: "petType",
	})

	// Act
	got, err := json.Marshal(s)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := `{"type":"object","discriminator":"petType"}`
	if string(got) != expected {
		t.Errorf("got %s, want %s", got, expected)
	}
}

func TestSchema_MarshalJSON_ReadOnly(t *testing.T) {
	// Arrange
	s := NewSchema(SchemaFields{Type: "string", ReadOnly: true})

	// Act
	got, err := json.Marshal(s)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := `{"type":"string","readOnly":true}`
	if string(got) != expected {
		t.Errorf("got %s, want %s", got, expected)
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
		t.Fatalf("unexpected error: %v", err)
	}
	expected := `{"type":"string","x-go-type":"MyString"}`
	if string(got) != expected {
		t.Errorf("got %s, want %s", got, expected)
	}
}
