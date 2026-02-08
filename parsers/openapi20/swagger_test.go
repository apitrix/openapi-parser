package openapi20

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for swagger.go - parseSwagger function
// =============================================================================

// --- Root Object Simple Properties ---

func TestParseSwagger_Host(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
host: "api.example.com"
paths: {}
`

	// Act
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "api.example.com", doc.Host)
}

func TestParseSwagger_BasePath(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
basePath: "/api/v1"
paths: {}
`

	// Act
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "/api/v1", doc.BasePath)
}

func TestParseSwagger_Schemes(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
schemes:
  - https
  - http
paths: {}
`

	// Act
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	assert.Equal(t, []string{"https", "http"}, doc.Schemes)
}

func TestParseSwagger_ConsumesProduces(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
consumes:
  - application/json
produces:
  - application/json
  - application/xml
paths: {}
`

	// Act
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	assert.Equal(t, []string{"application/json"}, doc.Consumes)
	assert.Equal(t, []string{"application/json", "application/xml"}, doc.Produces)
}

// --- Definitions ---

func TestParseSwagger_Definitions(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
definitions:
  Pet:
    type: object
    properties:
      name:
        type: string
`

	// Act
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	require.NotNil(t, doc.Definitions)
	require.Contains(t, doc.Definitions, "Pet")
	assert.Equal(t, "object", doc.Definitions["Pet"].Value.Type)
}

func TestParseSwagger_DefinitionsWithRef(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
definitions:
  Pet:
    type: object
  PetList:
    type: array
    items:
      $ref: "#/definitions/Pet"
`

	// Act
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	require.NotNil(t, doc.Definitions["PetList"].Value.Items)
	assert.Equal(t, "#/definitions/Pet", doc.Definitions["PetList"].Value.Items.Ref)
}

// --- Global Parameters ---

func TestParseSwagger_Parameters(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
parameters:
  limitParam:
    name: limit
    in: query
    type: integer
`

	// Act
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	require.NotNil(t, doc.Parameters)
	require.Contains(t, doc.Parameters, "limitParam")
	assert.Equal(t, "limit", doc.Parameters["limitParam"].Value.Name)
	assert.Equal(t, "query", doc.Parameters["limitParam"].Value.In)
}

// --- Global Responses ---

func TestParseSwagger_GlobalResponses(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
responses:
  NotFound:
    description: "Resource not found"
`

	// Act
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	require.NotNil(t, doc.Responses)
	require.Contains(t, doc.Responses, "NotFound")
	assert.Equal(t, "Resource not found", doc.Responses["NotFound"].Value.Description)
}

// --- SecurityDefinitions ---

func TestParseSwagger_SecurityDefinitions(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
securityDefinitions:
  api_key:
    type: apiKey
    name: X-API-Key
    in: header
`

	// Act
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	require.NotNil(t, doc.SecurityDefinitions)
	require.Contains(t, doc.SecurityDefinitions, "api_key")
	assert.Equal(t, "apiKey", doc.SecurityDefinitions["api_key"].Type)
	assert.Equal(t, "X-API-Key", doc.SecurityDefinitions["api_key"].Name)
}

// --- Security ---

func TestParseSwagger_Security(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
security:
  - api_key: []
  - oauth2:
      - read
      - write
`

	// Act
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	require.Len(t, doc.Security, 2)
	assert.Contains(t, doc.Security[0], "api_key")
	assert.Equal(t, []string{"read", "write"}, doc.Security[1]["oauth2"])
}

// --- Tags ---

func TestParseSwagger_Tags(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
tags:
  - name: pets
    description: "Pet operations"
`

	// Act
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	require.Len(t, doc.Tags, 1)
	assert.Equal(t, "pets", doc.Tags[0].Name)
	assert.Equal(t, "Pet operations", doc.Tags[0].Description)
}

// --- ExternalDocs ---

func TestParseSwagger_ExternalDocs(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
externalDocs:
  description: "Find more info here"
  url: "https://example.com/docs"
`

	// Act
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	require.NotNil(t, doc.ExternalDocs)
	assert.Equal(t, "https://example.com/docs", doc.ExternalDocs.URL)
}

// --- Extensions ---

func TestParseSwagger_Extensions(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
x-custom-field: "custom value"
x-internal: true
`

	// Act
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	require.NotNil(t, doc.VendorExtensions)
	assert.Equal(t, "custom value", doc.VendorExtensions["x-custom-field"])
	assert.Equal(t, true, doc.VendorExtensions["x-internal"])
}
