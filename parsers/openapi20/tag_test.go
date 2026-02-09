package openapi20

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for tag.go - parseTag, parseTags
// =============================================================================

// --- Basic Tag ---

func TestParseTag_Basic(t *testing.T) {
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
	require.Len(t, result.Document.Tags, 1)
	assert.Equal(t, "pets", result.Document.Tags[0].Name)
	assert.Equal(t, "Pet operations", result.Document.Tags[0].Description)
}

// --- Multiple Tags ---

func TestParseTag_Multiple(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
tags:
  - name: pets
    description: "Pet operations"
  - name: store
    description: "Store operations"
  - name: user
    description: "User operations"
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	require.Len(t, result.Document.Tags, 3)
	assert.Equal(t, "pets", result.Document.Tags[0].Name)
	assert.Equal(t, "store", result.Document.Tags[1].Name)
	assert.Equal(t, "user", result.Document.Tags[2].Name)
}

// --- Tag with ExternalDocs ---

func TestParseTag_WithExternalDocs(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
tags:
  - name: pets
    externalDocs:
      description: "Find more info"
      url: "https://example.com/pets"
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	require.NotNil(t, result.Document.Tags[0].ExternalDocs)
	assert.Equal(t, "https://example.com/pets", result.Document.Tags[0].ExternalDocs.URL)
	assert.Equal(t, "Find more info", result.Document.Tags[0].ExternalDocs.Description)
}

// --- Tag Name Only ---

func TestParseTag_NameOnly(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
tags:
  - name: pets
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	require.Len(t, result.Document.Tags, 1)
	assert.Equal(t, "pets", result.Document.Tags[0].Name)
	assert.Empty(t, result.Document.Tags[0].Description)
}

// --- Tag Extensions ---

func TestParseTag_Extensions(t *testing.T) {
	// Arrange
	yaml := `swagger: "2.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
tags:
  - name: pets
    x-internal: true
    x-display-name: "Pet Store"
`

	// Act
	result, err := Parse([]byte(yaml))

	// Assert
	require.NoError(t, err)
	require.NotNil(t, result.Document.Tags[0].VendorExtensions)
	assert.Equal(t, true, result.Document.Tags[0].VendorExtensions["x-internal"])
	assert.Equal(t, "Pet Store", result.Document.Tags[0].VendorExtensions["x-display-name"])
}
