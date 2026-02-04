package openapi30

import (
	_ "embed"
	"encoding/json"
	"sort"
	"strings"
	"testing"
)

//go:embed testdata/openapi-3.0-schema.json
var openapiSchemaJSON []byte

// JSONSchema represents a simplified JSON Schema structure
type JSONSchema struct {
	Type                 string                 `json:"type"`
	Properties           map[string]*JSONSchema `json:"properties"`
	PatternProperties    map[string]*JSONSchema `json:"patternProperties"`
	AdditionalProperties interface{}            `json:"additionalProperties"`
	Definitions          map[string]*JSONSchema `json:"definitions"`
	Ref                  string                 `json:"$ref"`
	OneOf                []*JSONSchema          `json:"oneOf"`
	AllOf                []*JSONSchema          `json:"allOf"`
}

// expectedProperties defines what properties we expect to parse for each OpenAPI type.
// This serves as our source of truth that we compare against the official schema.
var expectedProperties = map[string][]string{
	"Schema": {
		"title", "description", "type", "format", "pattern",
		"multipleOf", "maximum", "minimum", "exclusiveMaximum", "exclusiveMinimum",
		"maxLength", "minLength", "maxItems", "minItems",
		"maxProperties", "minProperties", "uniqueItems",
		"nullable", "readOnly", "writeOnly", "deprecated",
		"required", "enum", "default", "example",
		"allOf", "oneOf", "anyOf", "not", "items",
		"properties", "additionalProperties",
		"discriminator", "xml", "externalDocs",
	},
	"Info": {
		"title", "description", "termsOfService", "version",
		"contact", "license",
	},
	"Contact": {
		"name", "url", "email",
	},
	"License": {
		"name", "url",
	},
	"Server": {
		"url", "description", "variables",
	},
	"ServerVariable": {
		"enum", "default", "description",
	},
	"Components": {
		"schemas", "responses", "parameters", "examples",
		"requestBodies", "headers", "securitySchemes", "links", "callbacks",
	},
	"PathItem": {
		"$ref", "summary", "description", "servers", "parameters",
		// HTTP methods handled separately via patternProperties
	},
	"Operation": {
		"tags", "summary", "description", "externalDocs", "operationId",
		"parameters", "requestBody", "responses", "callbacks",
		"deprecated", "security", "servers",
	},
	"ExternalDocumentation": {
		"description", "url",
	},
	"Parameter": {
		"name", "in", "description", "required", "deprecated",
		"allowEmptyValue", "style", "explode", "allowReserved",
		"schema", "example", "examples", "content",
	},
	"RequestBody": {
		"description", "content", "required",
	},
	"MediaType": {
		"schema", "example", "examples", "encoding",
	},
	"Encoding": {
		"contentType", "headers", "style", "explode", "allowReserved",
	},
	"Response": {
		"description", "headers", "content", "links",
	},
	"Callback": {
		// Callbacks are maps of expression -> PathItem (no fixed properties)
	},
	"Example": {
		"summary", "description", "value", "externalValue",
	},
	"Link": {
		"operationRef", "operationId", "parameters", "requestBody",
		"description", "server",
	},
	"Header": {
		"description", "required", "deprecated", "allowEmptyValue",
		"style", "explode", "allowReserved",
		"schema", "example", "examples", "content",
	},
	"Tag": {
		"name", "description", "externalDocs",
	},
	"Discriminator": {
		"propertyName", "mapping",
	},
	"XML": {
		"name", "namespace", "prefix", "attribute", "wrapped",
	},
	"OAuthFlows": {
		"implicit", "password", "clientCredentials", "authorizationCode",
	},
	// Note: OAuthFlow types have different properties per flow type
	// We handle them all in a single parser with all possible properties
	"OAuthFlow": {
		"authorizationUrl", "tokenUrl", "refreshUrl", "scopes",
	},
}

func TestSpecCompliance(t *testing.T) {
	var schema JSONSchema
	if err := json.Unmarshal(openapiSchemaJSON, &schema); err != nil {
		t.Fatalf("failed to parse OpenAPI schema: %v", err)
	}

	for typeName, expectedProps := range expectedProperties {
		t.Run(typeName, func(t *testing.T) {
			// Find the definition in the schema
			def, ok := schema.Definitions[typeName]
			if !ok {
				// Try alternative names
				altNames := map[string]string{
					"ExternalDocumentation": "ExternalDocumentation",
					"OAuthFlow":             "ImplicitOAuthFlow", // Use one of the flow types
				}
				if altName, hasAlt := altNames[typeName]; hasAlt {
					def, ok = schema.Definitions[altName]
				}
			}

			if !ok {
				t.Logf("Definition %s not found in schema (may be combined type)", typeName)
				return
			}

			// Collect all properties from the official schema
			schemaProps := collectProperties(def, &schema)

			// Compare
			expectedSet := toSet(expectedProps)
			schemaSet := toSet(schemaProps)

			// Find missing (in schema but not in our expected list)
			var missing []string
			for prop := range schemaSet {
				if !expectedSet[prop] {
					missing = append(missing, prop)
				}
			}

			// Find extra (in our expected list but not in schema)
			var extra []string
			for prop := range expectedSet {
				if !schemaSet[prop] {
					extra = append(extra, prop)
				}
			}

			sort.Strings(missing)
			sort.Strings(extra)

			if len(missing) > 0 {
				t.Errorf("%s: properties in official schema but missing from our parser: %v", typeName, missing)
			}
			if len(extra) > 0 {
				t.Logf("%s: properties we handle that aren't in schema (extensions?): %v", typeName, extra)
			}
		})
	}
}

// collectProperties extracts property names from a JSON Schema definition
func collectProperties(def *JSONSchema, root *JSONSchema) []string {
	if def == nil {
		return nil
	}

	var props []string

	// Direct properties
	for name := range def.Properties {
		props = append(props, name)
	}

	// Handle oneOf (for polymorphic types like SecurityScheme)
	for _, sub := range def.OneOf {
		// Resolve $ref if present
		resolved := resolveRef(sub, root)
		props = append(props, collectProperties(resolved, root)...)
	}

	// Handle allOf (for composed types)
	for _, sub := range def.AllOf {
		resolved := resolveRef(sub, root)
		// Skip constraint schemas (ExampleXORExamples, etc.)
		if resolved != nil && resolved.Properties != nil {
			props = append(props, collectProperties(resolved, root)...)
		}
	}

	return props
}

// resolveRef resolves a $ref to its definition
func resolveRef(schema *JSONSchema, root *JSONSchema) *JSONSchema {
	if schema == nil {
		return nil
	}
	if schema.Ref == "" {
		return schema
	}

	// Parse ref like "#/definitions/Example"
	ref := schema.Ref
	if strings.HasPrefix(ref, "#/definitions/") {
		defName := strings.TrimPrefix(ref, "#/definitions/")
		if def, ok := root.Definitions[defName]; ok {
			return def
		}
	}
	return schema
}

func toSet(items []string) map[string]bool {
	set := make(map[string]bool)
	for _, item := range items {
		set[item] = true
	}
	return set
}
