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
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "api.example.com", result.Document.Host())
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
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "/api/v1", result.Document.BasePath())
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
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	assert.Equal(t, []string{"https", "http"}, result.Document.Schemes())
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
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	assert.Equal(t, []string{"application/json"}, result.Document.Consumes())
	assert.Equal(t, []string{"application/json", "application/xml"}, result.Document.Produces())
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
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	require.NotNil(t, result.Document.Definitions())
	require.Contains(t, result.Document.Definitions(), "Pet")
	assert.Equal(t, "object", result.Document.Definitions()["Pet"].Value.Type())
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
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	require.NotNil(t, result.Document.Definitions()["PetList"].Value.Items())
	assert.Equal(t, "#/definitions/Pet", result.Document.Definitions()["PetList"].Value.Items().Ref)
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
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	require.NotNil(t, result.Document.Parameters())
	require.Contains(t, result.Document.Parameters(), "limitParam")
	assert.Equal(t, "limit", result.Document.Parameters()["limitParam"].Value.Name())
	assert.Equal(t, "query", result.Document.Parameters()["limitParam"].Value.In())
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
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	require.NotNil(t, result.Document.Responses())
	require.Contains(t, result.Document.Responses(), "NotFound")
	assert.Equal(t, "Resource not found", result.Document.Responses()["NotFound"].Value.Description())
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
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	require.NotNil(t, result.Document.SecurityDefinitions())
	require.Contains(t, result.Document.SecurityDefinitions(), "api_key")
	assert.Equal(t, "apiKey", result.Document.SecurityDefinitions()["api_key"].Type())
	assert.Equal(t, "X-API-Key", result.Document.SecurityDefinitions()["api_key"].Name())
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
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	require.Len(t, result.Document.Security(), 2)
	assert.Contains(t, result.Document.Security()[0], "api_key")
	assert.Equal(t, []string{"read", "write"}, result.Document.Security()[1]["oauth2"])
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
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	require.Len(t, result.Document.Tags(), 1)
	assert.Equal(t, "pets", result.Document.Tags()[0].Name())
	assert.Equal(t, "Pet operations", result.Document.Tags()[0].Description())
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
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	require.NotNil(t, result.Document.ExternalDocs())
	assert.Equal(t, "https://example.com/docs", result.Document.ExternalDocs().URL())
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
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	require.NotNil(t, result.Document.VendorExtensions)
	assert.Equal(t, "custom value", result.Document.VendorExtensions["x-custom-field"])
	assert.Equal(t, true, result.Document.VendorExtensions["x-internal"])
}
