package openapi30x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for openapi.go - Parse function and OpenAPI document parsing
// =============================================================================

// --- Basic Document ---

func TestParseOpenAPI_Basic(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test API"
  version: "1.0.0"
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, doc)
	assert.Equal(t, "3.0.3", doc.OpenAPI)
	assert.Equal(t, "Test API", doc.Info.Title)
}

// --- Different OpenAPI Versions ---

func TestParseOpenAPI_Version303(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Equal(t, "3.0.3", doc.OpenAPI)
}

func TestParseOpenAPI_Version300(t *testing.T) {
	yaml := `openapi: "3.0.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Equal(t, "3.0.0", doc.OpenAPI)
}

// --- Complete Document ---

func TestParseOpenAPI_Complete(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Pet Store"
  version: "1.0.0"
  description: "A pet store API"
  contact:
    name: "Support"
    email: "support@example.com"
  license:
    name: "MIT"
servers:
  - url: https://api.example.com
tags:
  - name: pets
    description: "Pet operations"
paths:
  /pets:
    get:
      tags:
        - pets
      summary: "List pets"
      responses:
        "200":
          description: "OK"
components:
  schemas:
    Pet:
      type: object
  securitySchemes:
    apiKey:
      type: apiKey
      in: header
      name: X-API-Key
security:
  - apiKey: []
externalDocs:
  url: https://docs.example.com
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.NotNil(t, doc.Info)
	assert.Len(t, doc.Servers, 1)
	assert.Len(t, doc.Tags, 1)
	assert.NotEmpty(t, doc.Paths.Items)
	assert.NotNil(t, doc.Components)
	assert.Len(t, doc.Security, 1)
	assert.NotNil(t, doc.ExternalDocs)
}

// --- Extensions ---

func TestParseOpenAPI_Extensions(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
x-custom: "value"
x-internal: true
x-tags:
  - internal
  - beta
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, doc.Extensions)
	assert.Equal(t, "value", doc.Extensions["x-custom"])
	assert.Equal(t, true, doc.Extensions["x-internal"])
}

// --- Node Source ---

func TestParseOpenAPI_NodeSource(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Greater(t, doc.NodeSource.Start.Line, 0)
}

// --- JSON Bytes ---

func TestParseOpenAPI_JSON(t *testing.T) {
	json := `{
  "openapi": "3.0.3",
  "info": {
    "title": "Test API",
    "version": "1.0.0"
  },
  "paths": {}
}`
	doc, err := Parse([]byte(json))
	require.NoError(t, err)
	assert.Equal(t, "3.0.3", doc.OpenAPI)
	assert.Equal(t, "Test API", doc.Info.Title)
}

// --- Empty Paths ---

func TestParseOpenAPI_EmptyPaths(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "Test"
  version: "1.0"
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Empty(t, doc.Paths.Items)
}

// --- Error Cases ---

func TestParseOpenAPI_InvalidYAML(t *testing.T) {
	yaml := `{{{ invalid yaml`
	_, err := Parse([]byte(yaml))
	assert.Error(t, err)
}

func TestParseOpenAPI_MissingOpenAPI(t *testing.T) {
	yaml := `info:
  title: "Test"
  version: "1.0"
paths: {}
`
	_, err := Parse([]byte(yaml))
	assert.Error(t, err)
}

// --- Minimal Document ---

func TestParseOpenAPI_Minimal(t *testing.T) {
	yaml := `openapi: "3.0.3"
info:
  title: "a"
  version: "1"
paths: {}
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Equal(t, "a", doc.Info.Title)
	assert.Equal(t, "1", doc.Info.Version)
}
