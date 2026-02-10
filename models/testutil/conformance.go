// Package testutil provides shared testing utilities for models packages.
package testutil

import (
	"encoding/json"
	"os"
	"reflect"
	"strings"
	"testing"
	"unicode"
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
// Supports both exported struct fields (with JSON tags) and unexported fields (with getter methods).
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

		// Extract accessible fields from Go type via reflection
		// This checks both exported struct fields (JSON tags) AND getter methods
		goFields := ExtractAccessibleFields(goType)

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

// ExtractAccessibleFields uses reflection to get all JSON field names from a struct,
// checking both exported struct fields (via JSON tags) and getter methods (for private fields).
func ExtractAccessibleFields(t reflect.Type) map[string]bool {
	fields := make(map[string]bool)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return fields
	}

	// First: check exported struct fields with JSON tags (legacy approach)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Handle embedded structs (skip Node - Go-specific)
		if field.Anonymous {
			if field.Name != "Node" {
				for k, v := range ExtractAccessibleFields(field.Type) {
					fields[k] = v
				}
			}
			continue
		}

		// Get JSON tag from exported fields
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			continue
		}

		tagName := strings.Split(jsonTag, ",")[0]
		if tagName != "" && tagName != "-" {
			fields[tagName] = true
		}
	}

	// Second: check for getter methods (for private fields with readonly pattern)
	// We check both value and pointer receivers
	ptrType := reflect.PointerTo(t)
	for i := 0; i < ptrType.NumMethod(); i++ {
		method := ptrType.Method(i)

		// Only consider zero-parameter methods (getters)
		mType := method.Type
		// First param is the receiver, so getters have 1 param (just receiver) and 1 result
		if mType.NumIn() != 1 || mType.NumOut() != 1 {
			continue
		}

		// Skip non-exported methods
		if !method.IsExported() {
			continue
		}

		// Convert method name to JSON property name
		jsonName := methodToJSONPropertyName(method.Name)
		if jsonName != "" {
			fields[jsonName] = true
		}
	}

	return fields
}

// ExtractJSONFields is the legacy compatibility wrapper. Use ExtractAccessibleFields for
// comprehensive checking that includes getter methods.
func ExtractJSONFields(t reflect.Type) map[string]bool {
	return ExtractAccessibleFields(t)
}

// methodToJSONPropertyName converts a Go getter method name to its expected JSON schema
// property name. E.g., "Title" → "title", "ExternalDocs" → "externalDocs",
// "OperationID" → "operationId", "OpenAPIVersion" → "openapi"
func methodToJSONPropertyName(name string) string {
	// Special case mappings for names that don't follow simple camelCase rules
	switch name {
	case "OpenAPIVersion":
		return "openapi"
	case "SwaggerVersion":
		return "swagger"
	case "OperationID":
		return "operationId"
	case "Ref":
		return "$ref"
	case "Default":
		return "default"
	case "DefaultVal":
		return "" // internal helper, not a JSON property
	case "AuthorizationURL":
		return "authorizationUrl"
	case "TokenURL":
		return "tokenUrl"
	case "RefreshURL":
		return "refreshUrl"
	case "OpenIDConnectURL":
		return "openIdConnectUrl"
	// Skip infrastructure methods
	case "SetProperty":
		return ""
	}

	if len(name) == 0 {
		return ""
	}

	// Convert PascalCase to camelCase (first letter lowercase)
	runes := []rune(name)

	// Handle consecutive uppercase (e.g., "URL" → "url", "XML" → "xml")
	// Find the run of uppercase letters
	i := 0
	for i < len(runes) && unicode.IsUpper(runes[i]) {
		i++
	}

	if i == 0 {
		return name // already lowercase
	}

	if i == 1 {
		// Single uppercase letter: just lowercase it
		runes[0] = unicode.ToLower(runes[0])
	} else if i == len(runes) {
		// All uppercase: lowercase everything
		for j := range runes {
			runes[j] = unicode.ToLower(runes[j])
		}
	} else {
		// Multiple uppercase followed by lowercase: lowercase all but last uppercase
		// e.g., "URLBase" → "urlBase", but "ExternalDocs" → "externalDocs"
		for j := 0; j < i-1; j++ {
			runes[j] = unicode.ToLower(runes[j])
		}
	}

	return string(runes)
}
