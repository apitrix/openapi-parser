package openapi30

import (
	"testing"

	"github.com/apitrix/openapi-parser/models/testutil"
)

// allTypes lists all OpenAPI 3.0 model types for conformance testing.
var allTypes = []interface{}{
	OpenAPI{}, Info{}, Contact{}, License{},
	Server{}, ServerVariable{}, Components{},
	PathItem{}, Operation{}, Parameter{}, Header{},
	RequestBody{}, MediaType{}, Encoding{},
	Response{}, Responses{}, Schema{}, Discriminator{}, XML{},
	Example{}, Link{}, Callback{}, Tag{}, ExternalDocumentation{},
	SecurityScheme{}, OAuthFlows{}, OAuthFlow{},
}

// schemaNameMappings maps schema definition names to Go type names where they differ.
var schemaNameMappings = map[string]string{
	"ImplicitOAuthFlow":          "OAuthFlow",
	"PasswordOAuthFlow":          "OAuthFlow",
	"ClientCredentialsFlow":      "OAuthFlow",
	"AuthorizationCodeOAuthFlow": "OAuthFlow",
}

func TestSchemaConformance(t *testing.T) {
	testutil.RunSchemaConformance(t, testutil.SchemaConformanceConfig{
		SchemaPath:   "testdata/openapi-3.0-schema.json",
		Types:        allTypes,
		NameMappings: schemaNameMappings,
	})
}
