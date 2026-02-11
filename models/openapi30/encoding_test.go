package openapi30

import (
	"encoding/json"
	"testing"
)

func TestEncoding_MarshalJSON_AllFields(t *testing.T) {
	// Arrange
	explode := true
	e := NewEncoding("application/json", nil, "form", &explode, true)

	// Act
	got, err := json.Marshal(e)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"contentType", "style", "explode", "allowReserved"} {
		if _, ok := result[key]; !ok {
			t.Errorf("expected %q key", key)
		}
	}
}

func TestEncoding_MarshalJSON_Empty(t *testing.T) {
	// Arrange
	e := NewEncoding("", nil, "", nil, false)

	// Act
	got, err := json.Marshal(e)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
