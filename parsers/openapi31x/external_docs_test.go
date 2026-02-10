package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for externaldocs.go - parseExternalDocs function
// =============================================================================

// --- Basic ExternalDocs ---

func TestParseExternalDocs_Basic(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
externalDocs:
  url: "https://example.com/docs"
paths: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, result.Document.ExternalDocs())
	assert.Equal(t, "https://example.com/docs", result.Document.ExternalDocs().URL())
}

// --- With Description ---

func TestParseExternalDocs_WithDescription(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
externalDocs:
  description: "Find more documentation here"
  url: "https://example.com/docs"
paths: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Equal(t, "Find more documentation here", result.Document.ExternalDocs().Description())
}

// --- Extensions ---

func TestParseExternalDocs_Extensions(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
externalDocs:
  url: "https://example.com/docs"
  x-custom: "value"
  x-internal: true
paths: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	ext := result.Document.ExternalDocs().VendorExtensions
	require.NotNil(t, ext)
	assert.Equal(t, "value", ext["x-custom"])
}

// --- Operation-Level ExternalDocs ---

func TestParseExternalDocs_OperationLevel(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      externalDocs:
        description: "Pet documentation"
        url: "https://example.com/pets"
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	extDocs := result.Document.Paths().Items()["/pets"].Get().ExternalDocs()
	require.NotNil(t, extDocs)
	assert.Equal(t, "Pet documentation", extDocs.Description())
}

// --- Tag-Level ExternalDocs ---

func TestParseExternalDocs_TagLevel(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
tags:
  - name: pets
    externalDocs:
      url: "https://example.com/pet-docs"
paths: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, result.Document.Tags()[0].ExternalDocs())
	assert.Equal(t, "https://example.com/pet-docs", result.Document.Tags()[0].ExternalDocs().URL())
}

// --- Missing ExternalDocs ---

func TestParseExternalDocs_Missing(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Nil(t, result.Document.ExternalDocs())
}
