package openapi20

import (
	"testing"

	"github.com/apitrix/openapi-parser/models/testutil"
)

// allTypes lists all Swagger 2.0 model types for conformance testing.
var allTypes = []interface{}{
	Swagger{}, Info{}, Contact{}, License{},
	Paths{}, PathItem{}, Operation{},
	Parameter{}, Items{}, Schema{}, XML{},
	Responses{}, Response{}, Header{},
	SecurityScheme{}, Tag{}, ExternalDocs{},
}

func TestSchemaConformance(t *testing.T) {
	testutil.RunSchemaConformance(t, testutil.SchemaConformanceConfig{
		SchemaPath: "testdata/swagger-2.0-schema.json",
		Types:      allTypes,
	})
}
