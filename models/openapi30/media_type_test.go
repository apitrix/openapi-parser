package openapi30

import (
	"encoding/json"
	"testing"
)

func TestMediaType_MarshalJSON_WithSchema(t *testing.T) {
	// Arrange
	schema := &SchemaRef{}
	schema.SetValue(NewSchema(SchemaFields{Type: "object"}))
	mt := NewMediaType(schema, nil, nil, nil)

	// Act
	got, err := json.Marshal(mt)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if _, ok := result["schema"]; !ok {
		t.Error("expected 'schema' key")
	}
}

func TestMediaType_MarshalJSON_Empty(t *testing.T) {
	// Arrange
	mt := NewMediaType(nil, nil, nil, nil)

	// Act
	got, err := json.Marshal(mt)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
