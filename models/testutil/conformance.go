// Package testutil provides shared testing utilities for models packages.
package testutil

import (
	"encoding/json"
	"os"
	"reflect"
	"strings"
	"testing"
)

// SchemaConformanceConfig configures the schema conformance test.
type SchemaConformanceConfig struct {
	// SchemaPath is the path to the JSON schema file
	SchemaPath string
	// Types is the list of Go type instances to validate
	Types []interface{}
	// NameMappings maps schema definition names to Go type names where they differ
	NameMappings map[string]string
	// PropertyExclusions lists properties to skip per definition (e.g. schema bugs)
	PropertyExclusions map[string][]string
}

// RunSchemaConformance validates that Go struct fields match JSON schema properties.
// Schema-driven: iterates through schema definitions and validates each against Go types.
func RunSchemaConformance(t *testing.T, cfg SchemaConformanceConfig) {
	t.Helper()

	// Arrange
	data, err := os.ReadFile(cfg.SchemaPath)
	if err != nil {
		t.Fatalf("Failed to read schema file: %v", err)
	}

	var schema map[string]interface{}
	if err := json.Unmarshal(data, &schema); err != nil {
		t.Fatalf("Failed to parse schema: %v", err)
	}

	definitions, ok := schema["definitions"].(map[string]interface{})
	if !ok {
		// JSON Schema 2020-12 uses "$defs" instead of "definitions"
		definitions, ok = schema["$defs"].(map[string]interface{})
		if !ok {
			t.Fatal("Schema missing 'definitions' or '$defs'")
		}
	}

	// Build type map using reflection
	typeMap := make(map[string]reflect.Type)
	for _, instance := range cfg.Types {
		rt := reflect.TypeOf(instance)
		typeMap[rt.Name()] = rt
	}

	// Act & Assert: iterate through schema definitions
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
		if cfg.NameMappings != nil {
			if mapped, ok := cfg.NameMappings[schemaName]; ok {
				if mapped == "" {
					// Explicitly skipped definition
					continue
				}
				goTypeName = mapped
			}
		}

		// Look up Go type by name
		goType, ok := typeMap[goTypeName]
		if !ok {
			t.Errorf("❌ Schema %q: no matching Go type (expected %q)", schemaName, goTypeName)
			continue
		}

		// Extract JSON fields from Go type via reflection
		goFields := ExtractJSONFields(goType)

		// Validate each schema property exists in Go struct
		for propName := range props {
			// Check property exclusions
			if excluded := cfg.PropertyExclusions[schemaName]; len(excluded) > 0 {
				skip := false
				for _, ex := range excluded {
					if ex == propName {
						skip = true
						break
					}
				}
				if skip {
					continue
				}
			}
			if !goFields[propName] {
				t.Errorf("❌ Schema %q.%s: missing in Go type %s", schemaName, propName, goTypeName)
			}
		}
	}
}

// ExtractJSONFields uses reflection to get all JSON field names from a struct.
func ExtractJSONFields(t reflect.Type) map[string]bool {
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
				for k, v := range ExtractJSONFields(field.Type) {
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
