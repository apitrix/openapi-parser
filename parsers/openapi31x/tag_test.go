package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for tag.go - parseTag function
// =============================================================================

// --- Basic Tag ---

func TestParseTag_Basic(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
tags:
  - name: pets
paths: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.Len(t, result.Document.Tags(), 1)
	assert.Equal(t, "pets", result.Document.Tags()[0].Name())
}

// --- With Description ---

func TestParseTag_WithDescription(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
tags:
  - name: pets
    description: "Everything about your Pets"
paths: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Equal(t, "Everything about your Pets", result.Document.Tags()[0].Description())
}

// --- Multiple Tags ---

func TestParseTag_Multiple(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
tags:
  - name: pets
    description: "Pet operations"
  - name: users
    description: "User operations"
  - name: orders
    description: "Order operations"
  - name: store
    description: "Store operations"
paths: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, result.Document.Tags(), 4)
}

// --- External Docs ---

func TestParseTag_ExternalDocs(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
tags:
  - name: pets
    externalDocs:
      description: "Find out more"
      url: "https://example.com/pets"
paths: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	require.NotNil(t, result.Document.Tags()[0].ExternalDocs())
	assert.Equal(t, "Find out more", result.Document.Tags()[0].ExternalDocs().Description())
	assert.Equal(t, "https://example.com/pets", result.Document.Tags()[0].ExternalDocs().URL())
}

// --- Extensions ---

func TestParseTag_Extensions(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
tags:
  - name: pets
    x-displayName: "Pet Operations"
    x-order: 1
paths: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	ext := result.Document.Tags()[0].VendorExtensions
	require.NotNil(t, ext)
	assert.Equal(t, "Pet Operations", ext["x-displayName"])
}

// --- Complete Tag ---

func TestParseTag_Complete(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
tags:
  - name: pets
    description: "Endpoints for managing pets"
    externalDocs:
      description: "Pet docs"
      url: "https://example.com/pets"
    x-order: 1
paths: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	tag := result.Document.Tags()[0]
	assert.Equal(t, "pets", tag.Name())
	assert.NotEmpty(t, tag.Description())
	require.NotNil(t, tag.ExternalDocs())
	require.NotNil(t, tag.VendorExtensions)
}

// --- Empty Tags ---

func TestParseTag_Empty(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
tags: []
paths: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Empty(t, result.Document.Tags())
}

// --- Missing Tags ---

func TestParseTag_Missing(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Empty(t, result.Document.Tags())
}
