package openapi30

import (
	"encoding/json"
	"testing"
)

func TestXML_MarshalJSON_AllFields(t *testing.T) {
	// Arrange
	x := NewXML("animal", "http://example.com/schema", "ns", true, true)

	// Act
	got, err := json.Marshal(x)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"name":"animal","namespace":"http://example.com/schema","prefix":"ns","attribute":true,"wrapped":true}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestXML_MarshalJSON_OmitsFalseBoolsAndEmptyStrings(t *testing.T) {
	// Arrange
	x := NewXML("animal", "", "", false, false)

	// Act
	got, err := json.Marshal(x)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"name":"animal"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
