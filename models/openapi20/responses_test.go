package openapi20

import (
	"encoding/json"
	"testing"
)

func TestResponses_MarshalJSON_WithDefaultAndCodes(t *testing.T) {
	// Arrange
	def := &ResponseRef{Value: NewResponse("Unexpected error", nil, nil, nil)}
	codes := map[string]*ResponseRef{
		"200": {Value: NewResponse("OK", nil, nil, nil)},
		"404": {Value: NewResponse("Not Found", nil, nil, nil)},
	}
	r := NewResponses(def, codes)

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
	if m["default"] == nil {
		t.Error("default response should be present")
	}
	if m["200"] == nil {
		t.Error("200 response should be present")
	}
	if m["404"] == nil {
		t.Error("404 response should be present")
	}
}

func TestResponses_MarshalJSON_CodeOnly(t *testing.T) {
	// Arrange
	codes := map[string]*ResponseRef{
		"200": {Value: NewResponse("OK", nil, nil, nil)},
	}
	r := NewResponses(nil, codes)

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
	if m["default"] != nil {
		t.Error("default should not be present")
	}
	if m["200"] == nil {
		t.Error("200 response should be present")
	}
}
