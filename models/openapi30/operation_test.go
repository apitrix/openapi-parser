package openapi30

import (
	"encoding/json"
	"testing"
)

func TestOperation_MarshalJSON_AllFields(t *testing.T) {
	// Arrange
	defaultRef := &RefResponse{}
	defaultRef.SetValue(NewResponse("OK", nil, nil, nil))
	resp := NewResponses(
		defaultRef,
		nil,
	)
	op := NewOperation(
		[]string{"pets"}, "List pets", "Returns all pets", nil,
		"listPets", nil, nil, resp, nil, false, nil, nil,
	)

	// Act
	got, err := json.Marshal(op)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"tags", "summary", "description", "operationId", "responses"} {
		if _, ok := result[key]; !ok {
			t.Errorf("expected %q key", key)
		}
	}
}

func TestOperation_MarshalJSON_Minimal(t *testing.T) {
	// Arrange
	op := NewOperation(nil, "", "", nil, "", nil, nil, nil, nil, false, nil, nil)

	// Act
	got, err := json.Marshal(op)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
