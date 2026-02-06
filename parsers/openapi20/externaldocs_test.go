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
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	require.NotNil(t, doc.ExternalDocs)
	assert.Equal(t, "https://example.com/docs", doc.ExternalDocs.URL)
	assert.Equal(t, "Find more information", doc.ExternalDocs.Description)
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
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	require.NotNil(t, doc.ExternalDocs)
	assert.Equal(t, "https://example.com/docs", doc.ExternalDocs.URL)
	assert.Empty(t, doc.ExternalDocs.Description)
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
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	ed := doc.Paths.Items["/pets"].Get.ExternalDocs
	require.NotNil(t, ed)
	assert.Equal(t, "https://example.com/pets", ed.URL)
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
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	ed := doc.Tags[0].ExternalDocs
	require.NotNil(t, ed)
	assert.Equal(t, "https://example.com/pets", ed.URL)
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
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	ed := doc.Definitions["Pet"].Value.ExternalDocs
	require.NotNil(t, ed)
	assert.Equal(t, "https://example.com/pet-schema", ed.URL)
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
	doc, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	require.NotNil(t, doc.ExternalDocs.Extensions)
	assert.Equal(t, "2.0", doc.ExternalDocs.Extensions["x-version"])
}
