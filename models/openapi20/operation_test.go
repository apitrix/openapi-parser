package openapi20

import (
	"encoding/json"
	"testing"
)

func TestOperation_MarshalJSON_AllFields(t *testing.T) {
	// Arrange
	resp := NewResponses(
		&ResponseRef{Value: NewResponse("OK", nil, nil, nil)},
		nil,
	)
	op := NewOperation(
		[]string{"pets"}, "List pets", "Returns all pets", nil,
		"listPets", []string{"application/json"}, []string{"application/json"},
		nil, resp, nil, false, nil,
	)

	// Act
	got, err := json.Marshal(op)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var m map[string]interface{}
	if err := json.Unmarshal(got, &m); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if m["operationId"] != "listPets" {
		t.Errorf("operationId = %v, want listPets", m["operationId"])
	}
	if m["summary"] != "List pets" {
		t.Errorf("summary = %v, want List pets", m["summary"])
	}
}

func TestOperation_MarshalJSON_Minimal(t *testing.T) {
	// Arrange
	resp := NewResponses(
		&ResponseRef{Value: NewResponse("OK", nil, nil, nil)},
		nil,
	)
	op := NewOperation(nil, "", "", nil, "", nil, nil, nil, resp, nil, false, nil)

	// Act
	got, err := json.Marshal(op)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var m map[string]interface{}
	if err := json.Unmarshal(got, &m); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if m["responses"] == nil {
		t.Error("responses should not be nil")
	}
	if m["operationId"] != nil {
		t.Errorf("operationId should be omitted, got %v", m["operationId"])
	}
}

func TestOperation_MarshalJSON_Deprecated(t *testing.T) {
	// Arrange
	resp := NewResponses(nil, nil)
	op := NewOperation(nil, "", "", nil, "", nil, nil, nil, resp, nil, true, nil)

	// Act
	got, err := json.Marshal(op)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var m map[string]interface{}
	if err := json.Unmarshal(got, &m); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if m["deprecated"] != true {
		t.Errorf("deprecated = %v, want true", m["deprecated"])
	}
}
