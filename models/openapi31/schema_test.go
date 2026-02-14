package openapi31

import (
	"openapi-parser/models/shared"
	"encoding/json"
	"testing"
)

func TestSchema_MarshalJSON_BasicType(t *testing.T) {
	s := NewSchema(SchemaFields{Type: SchemaType{Single: "string"}, Description: "A name", Format: "email"})
	got, err := json.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"type":"string","description":"A name","format":"email"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestSchema_MarshalJSON_TypeArray(t *testing.T) {
	s := NewSchema(SchemaFields{Type: SchemaType{Array: []string{"string", "null"}}})
	got, err := json.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"type":["string","null"]}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestSchema_MarshalJSON_AdditionalPropertiesAsBoolean(t *testing.T) {
	allowed := false
	s := NewSchema(SchemaFields{Type: SchemaType{Single: "object"}, AdditionalPropertiesAllowed: &allowed})
	got, err := json.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"type":"object","additionalProperties":false}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestSchema_MarshalJSON_Const(t *testing.T) {
	s := NewSchema(SchemaFields{Const: "fixed_value"})
	got, err := json.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"const":"fixed_value"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestSchema_MarshalJSON_IfThenElse(t *testing.T) {
	ifSchema := &shared.RefWithMeta[Schema]{}
	ifSchema.SetValue(NewSchema(SchemaFields{Type: SchemaType{Single: "string"}}))
	thenSchema := &shared.RefWithMeta[Schema]{}
	thenSchema.SetValue(NewSchema(SchemaFields{Format: "email"}))
	elseSchema := &shared.RefWithMeta[Schema]{}
	elseSchema.SetValue(NewSchema(SchemaFields{Format: "uri"}))
	s := NewSchema(SchemaFields{If: ifSchema, Then: thenSchema, Else: elseSchema})
	got, err := json.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"if", "then", "else"} {
		if _, ok := result[key]; !ok {
			t.Errorf("expected %q key", key)
		}
	}
}

func TestSchema_MarshalJSON_PrefixItems(t *testing.T) {
	item1 := &shared.RefWithMeta[Schema]{}
	item1.SetValue(NewSchema(SchemaFields{Type: SchemaType{Single: "string"}}))
	item2 := &shared.RefWithMeta[Schema]{}
	item2.SetValue(NewSchema(SchemaFields{Type: SchemaType{Single: "integer"}}))
	s := NewSchema(SchemaFields{PrefixItems: []*shared.RefWithMeta[Schema]{item1, item2}})
	got, err := json.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if _, ok := result["prefixItems"]; !ok {
		t.Error("expected 'prefixItems' key")
	}
}

func TestSchema_MarshalJSON_EmptySchema(t *testing.T) {
	s := NewSchema(SchemaFields{})
	got, err := json.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}
	want := `{}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestSchema_MarshalJSON_WithExtensions(t *testing.T) {
	s := NewSchema(SchemaFields{Type: SchemaType{Single: "string"}})
	s.VendorExtensions = map[string]interface{}{"x-go-type": "MyString"}
	got, err := json.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"type":"string","x-go-type":"MyString"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestSchema_MarshalJSON_AllOf(t *testing.T) {
	ref1 := shared.NewRefWithMeta[Schema]("#/components/schemas/Base")
	ref2 := shared.NewRefWithMeta[Schema]("#/components/schemas/Extended")
	s := NewSchema(SchemaFields{AllOf: []*shared.RefWithMeta[Schema]{ref1, ref2}})
	got, err := json.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"allOf":[{"$ref":"#/components/schemas/Base"},{"$ref":"#/components/schemas/Extended"}]}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestSchema_MarshalJSON_Required(t *testing.T) {
	s := NewSchema(SchemaFields{Type: SchemaType{Single: "object"}, Required: []string{"name", "age"}})
	got, err := json.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"required":["name","age"],"type":"object"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestSchema_MarshalJSON_NumericConstraints(t *testing.T) {
	max := 100.0
	min := 0.0
	mult := 5.0
	s := NewSchema(SchemaFields{
		Type:       SchemaType{Single: "integer"},
		Maximum:    &max,
		Minimum:    &min,
		MultipleOf: &mult,
	})
	got, err := json.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"multipleOf":5,"maximum":100,"minimum":0,"type":"integer"}`
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
