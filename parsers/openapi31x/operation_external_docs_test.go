package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for operation_externaldocs.go - parseOperationExternalDocs function
// =============================================================================

// --- Basic ExternalDocs ---

func TestParseOperationExternalDocs_Basic(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      externalDocs:
        url: "https://example.com/pets-docs"
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	extDocs := result.Document.Paths().Items()["/pets"].Get().ExternalDocs()
	require.NotNil(t, extDocs)
	assert.Equal(t, "https://example.com/pets-docs", extDocs.URL())
}

// --- With Description ---

func TestParseOperationExternalDocs_WithDescription(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      externalDocs:
        description: "Find more info about pets"
        url: "https://example.com/pets-docs"
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	extDocs := result.Document.Paths().Items()["/pets"].Get().ExternalDocs()
	assert.Equal(t, "Find more info about pets", extDocs.Description())
}

// --- Missing ExternalDocs ---

func TestParseOperationExternalDocs_Missing(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Nil(t, result.Document.Paths().Items()["/pets"].Get().ExternalDocs())
}

// --- With Extensions ---

func TestParseOperationExternalDocs_Extensions(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths:
  /pets:
    get:
      externalDocs:
        url: "https://example.com/docs"
        x-custom: "value"
      responses:
        "200":
          description: "OK"
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	ext := result.Document.Paths().Items()["/pets"].Get().ExternalDocs().VendorExtensions
	require.NotNil(t, ext)
	assert.Equal(t, "value", ext["x-custom"])
}
