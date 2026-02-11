package openapi20

import (
	"encoding/json"
	"testing"
)

func TestXML_MarshalJSON_AllFields(t *testing.T) {
	// Arrange
	x := NewXML("animal", "http://example.com/schema", "ex", true, true)

	// Act
	got, err := json.Marshal(x)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := `{"name":"animal","namespace":"http://example.com/schema","prefix":"ex","attribute":true,"wrapped":true}`
	if string(got) != expected {
		t.Errorf("got %s, want %s", got, expected)
	}
}

func TestXML_MarshalJSON_NameOnly(t *testing.T) {
	// Arrange
	x := NewXML("item", "", "", false, false)

	// Act
	got, err := json.Marshal(x)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := `{"name":"item"}`
	if string(got) != expected {
		t.Errorf("got %s, want %s", got, expected)
	}
}
