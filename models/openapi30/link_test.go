package openapi30

import (
	"encoding/json"
	"testing"
)

func TestLink_MarshalJSON_OperationRef(t *testing.T) {
	// Arrange
	l := NewLink("#/paths/~1users~1{userId}/get", "", nil, nil, "Get user by ID", nil)

	// Act
	got, err := json.Marshal(l)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if _, ok := result["operationRef"]; !ok {
		t.Error("expected 'operationRef' key")
	}
	if _, ok := result["description"]; !ok {
		t.Error("expected 'description' key")
	}
}

func TestLink_MarshalJSON_OperationId(t *testing.T) {
	// Arrange
	l := NewLink("", "getUser", map[string]interface{}{"userId": "$response.body#/id"}, nil, "", nil)

	// Act
	got, err := json.Marshal(l)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if _, ok := result["operationId"]; !ok {
		t.Error("expected 'operationId' key")
	}
	if _, ok := result["parameters"]; !ok {
		t.Error("expected 'parameters' key")
	}
}
