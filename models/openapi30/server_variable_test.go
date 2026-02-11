package openapi30

import (
	"encoding/json"
	"testing"
)

func TestServerVariable_MarshalJSON_AllFields(t *testing.T) {
	// Arrange
	sv := NewServerVariable("8080", "Server port", []string{"8080", "443"})

	// Act
	got, err := json.Marshal(sv)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"enum":["8080","443"],"default":"8080","description":"Server port"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestServerVariable_MarshalJSON_DefaultOnly(t *testing.T) {
	// Arrange
	sv := NewServerVariable("prod", "", nil)

	// Act
	got, err := json.Marshal(sv)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"default":"prod"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
