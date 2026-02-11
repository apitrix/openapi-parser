package openapi30

import (
	"encoding/json"
	"testing"
)

func TestPathItem_MarshalJSON_WithOperations(t *testing.T) {
	// Arrange
	getOp := NewOperation(nil, "Get pet", "", nil, "getPet", nil, nil, nil, nil, false, nil, nil)
	postOp := NewOperation(nil, "Create pet", "", nil, "createPet", nil, nil, nil, nil, false, nil, nil)
	pi := NewPathItem("", "Pet operations", "", getOp, nil, postOp, nil, nil, nil, nil, nil, nil, nil)

	// Act
	got, err := json.Marshal(pi)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"summary", "get", "post"} {
		if _, ok := result[key]; !ok {
			t.Errorf("expected %q key", key)
		}
	}
	// Unused HTTP methods should be omitted
	for _, key := range []string{"put", "delete", "options", "head", "patch", "trace"} {
		if _, ok := result[key]; ok {
			t.Errorf("expected %q to be omitted", key)
		}
	}
}

func TestPathItem_MarshalJSON_WithRef(t *testing.T) {
	// Arrange
	pi := NewPathItem("#/components/pathItems/Shared", "", "", nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)

	// Act
	got, err := json.Marshal(pi)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"$ref":"#/components/pathItems/Shared"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
