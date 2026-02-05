package openapi30

import (
	"encoding/json"
	"os"
	"reflect"
	"strings"
	"testing"
)

// allTypes returns all OpenAPI 3.0 model types for reflection-based discovery.
// This is the ONLY place types are listed - the test itself is schema-driven.
var allTypes = []interface{}{
	OpenAPI{}, Info{}, Contact{}, License{},
	Server{}, ServerVariable{}, Components{},
	PathItem{}, Operation{}, Parameter{}, Header{},
	RequestBody{}, MediaType{}, Encoding{},
	Response{}, Responses{}, Schema{}, Discriminator{}, XML{},
	Example{}, Link{}, Callback{}, Tag{}, ExternalDocumentation{},
	SecurityScheme{}, OAuthFlows{}, OAuthFlow{},
}

// schemaNameToGoType maps schema names to Go type names where they differ.
// OAuth flows: schema has 4 types, Go uses 1 unified type.
var schemaNameToGoType = map[string]string{
	"ImplicitOAuthFlow":          "OAuthFlow",
	"PasswordOAuthFlow":          "OAuthFlow",
	"ClientCredentialsFlow":      "OAuthFlow",
	"AuthorizationCodeOAuthFlow": "OAuthFlow",
}

// TestSchemaConformance validates that Go struct fields match JSON schema properties.
// Schema-driven: iterates through schema definitions and validates each against Go types.
func TestSchemaConformance(t *testing.T) {
	// Load schema
	data, err := os.ReadFile("testdata/openapi-3.0-schema.json")
	if err != nil {
		t.Fatalf("Failed to read schema file: %v", err)
	}

	var schema map[string]interface{}
	if err := json.Unmarshal(data, &schema); err != nil {
		t.Fatalf("Failed to parse schema: %v", err)
	}

	definitions, ok := schema["definitions"].(map[string]interface{})
	if !ok {
		t.Fatal("Schema missing 'definitions'")
	}

	// Build type map from allTypes using reflection
	typeMap := make(map[string]reflect.Type)
	for _, instance := range allTypes {
		t := reflect.TypeOf(instance)
		typeMap[t.Name()] = t
	}

	// Iterate through schema definitions
	for schemaName, def := range definitions {
		defMap, ok := def.(map[string]interface{})
		if !ok {
			continue
		}

		props, ok := defMap["properties"].(map[string]interface{})
		if !ok {
			continue
		}

		// Map schema name to Go type name
		goTypeName := schemaName
		if mapped, ok := schemaNameToGoType[schemaName]; ok {
			goTypeName = mapped
		}

		// Look up Go type by name
		goType, ok := typeMap[goTypeName]
		if !ok {
			t.Errorf("❌ Schema %q: no matching Go type (expected %q)", schemaName, goTypeName)
			continue
		}

		// Extract JSON fields from Go type via reflection
		goFields := extractJSONFields(goType)

		// Validate each schema property exists in Go struct
		for propName := range props {
			if !goFields[propName] {
				t.Errorf("❌ Schema %q.%s: missing in Go type %s", schemaName, propName, goTypeName)
			}
		}
	}
}

// extractJSONFields uses reflection to get all JSON field names from a struct.
func extractJSONFields(t reflect.Type) map[string]bool {
	fields := make(map[string]bool)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return fields
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Handle embedded structs (skip Node - Go-specific)
		if field.Anonymous {
			if field.Name != "Node" {
				for k, v := range extractJSONFields(field.Type) {
					fields[k] = v
				}
			}
			continue
		}

		// Get JSON tag
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			continue
		}

		tagName := strings.Split(jsonTag, ",")[0]
		if tagName != "" && tagName != "-" {
			fields[tagName] = true
		}
	}

	return fields
}
