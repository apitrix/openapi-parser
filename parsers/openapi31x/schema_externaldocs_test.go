package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for schema_externaldocs.go - ParseExternalDocs method on schemas
// =============================================================================

// --- Basic ExternalDocs ---

func TestParseSchemaExternalDocs_Basic(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Pet:
      type: object
      externalDocs:
        url: "https://example.com/pet-docs"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	extDocs := doc.Components.Schemas["Pet"].Value.ExternalDocs
	require.NotNil(t, extDocs)
	assert.Equal(t, "https://example.com/pet-docs", extDocs.URL)
}

// --- With Description ---

func TestParseSchemaExternalDocs_Description(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Pet:
      type: object
      externalDocs:
        description: "Find more info about pets here"
        url: "https://example.com/pet-docs"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	extDocs := doc.Components.Schemas["Pet"].Value.ExternalDocs
	assert.Equal(t, "Find more info about pets here", extDocs.Description)
}

// --- Missing ExternalDocs ---

func TestParseSchemaExternalDocs_Missing(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Pet:
      type: object
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	extDocs := doc.Components.Schemas["Pet"].Value.ExternalDocs
	assert.Nil(t, extDocs)
}

// --- Extensions ---

func TestParseSchemaExternalDocs_Extensions(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  schemas:
    Pet:
      type: object
      externalDocs:
        url: "https://example.com/docs"
        x-custom: "value"
`
	doc, err := Parse([]byte(yaml))
	require.NoError(t, err)
	extDocs := doc.Components.Schemas["Pet"].Value.ExternalDocs
	require.NotNil(t, extDocs.Extensions)
	assert.Equal(t, "value", extDocs.Extensions["x-custom"])
}
