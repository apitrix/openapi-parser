package openapi20

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for externaldocs.go - parseExternalDocs
// =============================================================================

// --- Basic ExternalDocs ---

func TestParseExternalDocs_Basic(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
externalDocs:
  description: "Find more information"
  url: "https://example.com/docs"
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	require.NotNil(t, result.Document.ExternalDocs())
	assert.Equal(t, "https://example.com/docs", result.Document.ExternalDocs().URL())
	assert.Equal(t, "Find more information", result.Document.ExternalDocs().Description())
}

// --- ExternalDocs URL Only ---

func TestParseExternalDocs_URLOnly(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
externalDocs:
  url: "https://example.com/docs"
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	require.NotNil(t, result.Document.ExternalDocs())
	assert.Equal(t, "https://example.com/docs", result.Document.ExternalDocs().URL())
	assert.Empty(t, result.Document.ExternalDocs().Description())
}

// --- ExternalDocs on Operation ---

func TestParseExternalDocs_OnOperation(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      externalDocs:
        description: "Get pets documentation"
        url: "https://example.com/pets"
      responses:
        "200":
          description: "OK"
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	ed := result.Document.Paths().Items()["/pets"].Get().ExternalDocs()
	require.NotNil(t, ed)
	assert.Equal(t, "https://example.com/pets", ed.URL())
}

// --- ExternalDocs on Tag ---

func TestParseExternalDocs_OnTag(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
tags:
  - name: pets
    externalDocs:
      url: "https://example.com/pets"
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	ed := result.Document.Tags()[0].ExternalDocs()
	require.NotNil(t, ed)
	assert.Equal(t, "https://example.com/pets", ed.URL())
}

// --- ExternalDocs on Schema ---

func TestParseExternalDocs_OnSchema(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
definitions:
  Pet:
    type: object
    externalDocs:
      url: "https://example.com/pet-schema"
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	ed := result.Document.Definitions()["Pet"].Value().ExternalDocs()
	require.NotNil(t, ed)
	assert.Equal(t, "https://example.com/pet-schema", ed.URL())
}

// --- ExternalDocs Extensions ---

func TestParseExternalDocs_Extensions(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
externalDocs:
  url: "https://example.com/docs"
  x-version: "2.0"
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	require.NotNil(t, result.Document.ExternalDocs().VendorExtensions)
	assert.Equal(t, "2.0", result.Document.ExternalDocs().VendorExtensions["x-version"])
}
