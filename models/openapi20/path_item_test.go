package openapi20

import (
	"openapi-parser/models/shared"
	"encoding/json"
	"testing"
)

func TestPathItem_MarshalJSON_WithOperations(t *testing.T) {
	// Arrange
	getOp := NewOperation(nil, "Get pets", "", nil, "getPets", nil, nil, nil, nil, nil, false, nil)
	postOp := NewOperation(nil, "Create pet", "", nil, "createPet", nil, nil, nil, nil, nil, false, nil)
	pi := NewPathItem("", getOp, nil, postOp, nil, nil, nil, nil, nil)

	// Act
	got, err := json.Marshal(pi)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var m map[string]interface{}
	if err := json.Unmarshal(got, &m); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if m["get"] == nil {
		t.Error("get should not be nil")
	}
	if m["post"] == nil {
		t.Error("post should not be nil")
	}
	if m["put"] != nil {
		t.Error("put should be omitted")
	}
}

func TestPathItem_MarshalJSON_WithRef(t *testing.T) {
	// Arrange
	pi := NewPathItem("#/paths/shared", nil, nil, nil, nil, nil, nil, nil, nil)

	// Act
	got, err := json.Marshal(pi)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := `{"$ref":"#/paths/shared"}`
	if string(got) != expected {
		t.Errorf("got %s, want %s", got, expected)
	}
}

func TestPathItem_MarshalJSON_WithParameters(t *testing.T) {
	// Arrange
	idParam := &shared.Ref[Parameter]{}
	idParam.SetValue(NewParameter(ParameterFields{Name: "id", In: "path", Type: "string"}))
	params := []*shared.Ref[Parameter]{
		idParam,
	}
	pi := NewPathItem("", nil, nil, nil, nil, nil, nil, nil, params)

	// Act
	got, err := json.Marshal(pi)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var m map[string]interface{}
	if err := json.Unmarshal(got, &m); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if m["parameters"] == nil {
		t.Error("parameters should not be nil")
	}
}
