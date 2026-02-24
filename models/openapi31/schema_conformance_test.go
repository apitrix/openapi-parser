package openapi31

import (
	"testing"

	"github.com/apitrix/openapi-parser/models/testutil"
)

// allTypes lists all OpenAPI 3.1 model types for conformance testing.
var allTypes = []interface{}{
	OpenAPI{}, Info{}, Contact{}, License{},
	Server{}, ServerVariable{}, Components{},
	PathItem{}, Operation{}, Parameter{}, Header{},
	RequestBody{}, MediaType{}, Encoding{},
	Response{}, Responses{}, Schema{}, Discriminator{}, XML{},
	Example{}, Link{}, Callback{}, Tag{}, ExternalDocumentation{},
	SecurityScheme{}, OAuthFlows{}, OAuthFlow{},
}

// schemaNameMappings maps JSON Schema $defs names (kebab-case) to Go type names.
// Use empty string "" to skip definitions with no matching Go struct.
// Used for both 3.1 and 3.2 schemas.
var schemaNameMappings = map[string]string{
	// kebab-case to PascalCase
	"server-variable":        "ServerVariable",
	"path-item":              "PathItem",
	"external-documentation": "ExternalDocumentation",
	"media-type":             "MediaType",
	"request-body":           "RequestBody",
	"oauth-flows":            "OAuthFlows",
	"security-scheme":        "SecurityScheme",
	"security-requirement":   "SecurityRequirement",
	"discriminator":          "Discriminator",

	// PascalCase names that match directly
	"info":       "Info",
	"contact":    "Contact",
	"license":    "License",
	"server":     "Server",
	"operation":  "Operation",
	"parameter":  "Parameter",
	"header":     "Header",
	"encoding":   "Encoding",
	"response":   "Response",
	"responses":  "Responses",
	"schema":     "Schema",
	"example":    "Example",
	"link":       "Link",
	"tag":        "Tag",
	"components": "Components",

	// Abstract, union, or map types with no single Go struct - skip
	"reference":                    "",
	"examples":                     "",
	"callbacks":                    "",
	"content":                      "",
	"paths":                        "",
	"specification-extensions":     "",
	"map-of-strings":               "",
	"path-item-or-reference":       "",
	"response-or-reference":        "",
	"parameter-or-reference":       "",
	"example-or-reference":         "",
	"request-body-or-reference":    "",
	"header-or-reference":          "",
	"security-scheme-or-reference": "",
	"link-or-reference":            "",
	"callbacks-or-reference":       "",
}

func TestSchemaConformance(t *testing.T) {
	testutil.RunSchemaConformance(t, testutil.SchemaConformanceConfig{
		SchemaPath:   "testdata/openapi-3.1-schema.json",
		Types:        allTypes,
		NameMappings: schemaNameMappings,
		PropertyExclusions: map[string][]string{
			// The official 3.1 JSON schema has "body" referencing "server"
			// but the OAS spec itself names this property "server".
			"link": {"body"},
		},
	})
}

func TestSchemaConformance32(t *testing.T) {
	testutil.RunSchemaConformance(t, testutil.SchemaConformanceConfig{
		SchemaPath:   "testdata/openapi-3.2-schema.json",
		Types:        allTypes,
		NameMappings: schemaNameMappings,
		PropertyExclusions: map[string][]string{
			"link": {"body"},
		},
	})
}
