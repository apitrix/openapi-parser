package openapi30

import (
	"encoding/json"
	"testing"
)

func TestParameter_MarshalJSON_AllFields(t *testing.T) {
	// Arrange
	schema := &SchemaRef{}
	schema.SetValue(NewSchema(SchemaFields{Type: "integer"}))
	p := NewParameter("limit", "query", "Max items", false, false, false, "", nil, false, schema, nil, nil, nil)

	// Act
	got, err := json.Marshal(p)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"name", "in", "description", "schema"} {
		if _, ok := result[key]; !ok {
			t.Errorf("expected %q key", key)
		}
	}
	// false booleans should be omitted
	for _, key := range []string{"required", "deprecated", "allowEmptyValue"} {
		if _, ok := result[key]; ok {
			t.Errorf("expected %q to be omitted (false bool)", key)
		}
	}
}

func TestParameter_MarshalJSON_RequiredParam(t *testing.T) {
	// Arrange
	p := NewParameter("id", "path", "", true, false, false, "", nil, false, nil, nil, nil, nil)

	// Act
	got, err := json.Marshal(p)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	want := `{"name":"id","in":"path","required":true}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
