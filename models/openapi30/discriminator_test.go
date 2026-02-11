package openapi30

import (
	"encoding/json"
	"testing"
)

func TestDiscriminator_MarshalJSON_AllFields(t *testing.T) {
	// Arrange
	d := NewDiscriminator("petType", map[string]string{"dog": "#/components/schemas/Dog"})

	// Act
	got, err := json.Marshal(d)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"propertyName":"petType","mapping":{"dog":"#/components/schemas/Dog"}}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestDiscriminator_MarshalJSON_NoMapping(t *testing.T) {
	// Arrange
	d := NewDiscriminator("type", nil)

	// Act
	got, err := json.Marshal(d)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"propertyName":"type"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
