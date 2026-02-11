package openapi31

import (
	"encoding/json"
	"testing"
)

func TestDiscriminator_MarshalJSON_AllFields(t *testing.T) {
	d := NewDiscriminator("petType", map[string]string{"dog": "#/components/schemas/Dog"})
	got, err := json.Marshal(d)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if _, ok := result["propertyName"]; !ok {
		t.Error("expected 'propertyName' key")
	}
	if _, ok := result["mapping"]; !ok {
		t.Error("expected 'mapping' key")
	}
}

func TestDiscriminator_MarshalJSON_PropertyNameOnly(t *testing.T) {
	d := NewDiscriminator("petType", nil)
	got, err := json.Marshal(d)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"propertyName":"petType"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
