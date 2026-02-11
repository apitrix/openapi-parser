package openapi31

import (
	"encoding/json"
	"testing"
)

func TestOpenAPI_MarshalJSON_Minimal(t *testing.T) {
	info := NewInfo("Test API", "", "", "", "1.0.0", nil, nil)
	doc := NewOpenAPI("3.1.0", info)
	got, err := json.Marshal(doc)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if _, ok := result["openapi"]; !ok {
		t.Error("expected 'openapi' key")
	}
	if _, ok := result["info"]; !ok {
		t.Error("expected 'info' key")
	}
}

func TestOpenAPI_MarshalJSON_WithWebhooks(t *testing.T) {
	info := NewInfo("Test API", "", "", "", "1.0.0", nil, nil)
	doc := NewOpenAPI("3.1.0", info)
	doc.SetProperty("webhooks", map[string]*PathItemRef{
		"newPet": {Ref: "#/components/pathItems/NewPet"},
	})
	got, err := json.Marshal(doc)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if _, ok := result["webhooks"]; !ok {
		t.Error("expected 'webhooks' key (3.1-specific)")
	}
}

func TestOpenAPI_MarshalJSON_JsonSchemaDialect(t *testing.T) {
	info := NewInfo("Test API", "", "", "", "1.0.0", nil, nil)
	doc := NewOpenAPI("3.1.0", info)
	doc.SetProperty("jsonSchemaDialect", "https://json-schema.org/draft/2020-12/schema")
	got, err := json.Marshal(doc)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(got, &result); err != nil {
		t.Fatal(err)
	}
	if _, ok := result["jsonSchemaDialect"]; !ok {
		t.Error("expected 'jsonSchemaDialect' key (3.1-specific)")
	}
}
