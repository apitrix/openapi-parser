package openapi30

import (
	"encoding/json"
	"testing"
)

func TestHeader_MarshalJSON_AllFields(t *testing.T) {
	// Arrange
	schema := &SchemaRef{Value: NewSchema(SchemaFields{Type: "integer"})}
	explode := true
	h := NewHeader("Rate limit", true, false, false, "simple", &explode, false, schema, nil, nil, nil)

	// Act
	got, err := json.Marshal(h)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"description", "required", "style", "explode", "schema"} {
		if _, ok := result[key]; !ok {
			t.Errorf("expected %q key", key)
		}
	}
}

func TestHeader_MarshalJSON_Minimal(t *testing.T) {
	// Arrange
	h := NewHeader("", false, false, false, "", nil, false, nil, nil, nil, nil)

	// Act
	got, err := json.Marshal(h)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
