package openapi20

import (
	"openapi-parser/models/shared"
	"encoding/json"
	"testing"
)

func TestResponse_MarshalJSON_AllFields(t *testing.T) {
	// Arrange
	schema := &shared.Ref[Schema]{}
	schema.SetValue(NewSchema(SchemaFields{Type: "object"}))
	headers := map[string]*Header{
		"X-Rate-Limit": NewHeader(HeaderFields{Type: "integer"}),
	}
	examples := map[string]interface{}{
		"application/json": map[string]interface{}{"id": 1},
	}
	r := NewResponse("Success", schema, headers, examples)

	// Act
	got, err := json.Marshal(r)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var m map[string]interface{}
	if err := json.Unmarshal(got, &m); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if m["description"] != "Success" {
		t.Errorf("description = %v, want Success", m["description"])
	}
	if m["schema"] == nil {
		t.Error("schema should not be nil")
	}
	if m["headers"] == nil {
		t.Error("headers should not be nil")
	}
}

func TestResponse_MarshalJSON_DescriptionOnly(t *testing.T) {
	// Arrange
	r := NewResponse("Not Found", nil, nil, nil)

	// Act
	got, err := json.Marshal(r)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := `{"description":"Not Found"}`
	if string(got) != expected {
		t.Errorf("got %s, want %s", got, expected)
	}
}
