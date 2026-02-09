package openapi31x

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Tests for ref_security_scheme.go - security scheme reference parsing
// =============================================================================

// --- Basic Reference ---

func TestParseRefSecurityScheme_Basic(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  securitySchemes:
    apiKey:
      $ref: '#/components/securitySchemes/SharedApiKey'
    SharedApiKey:
      type: apiKey
      in: header
      name: X-API-Key
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	ref := result.Document.Components.SecuritySchemes["apiKey"]
	assert.Equal(t, "#/components/securitySchemes/SharedApiKey", ref.Ref)
}

// --- Multiple Schemes ---

func TestParseRefSecurityScheme_Multiple(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  securitySchemes:
    apiKey:
      type: apiKey
      in: header
      name: X-API-Key
    bearer:
      type: http
      scheme: bearer
    oauth2:
      type: oauth2
      flows:
        clientCredentials:
          tokenUrl: https://example.com/token
          scopes: {}
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	assert.Len(t, result.Document.Components.SecuritySchemes, 3)
	assert.Equal(t, "apiKey", result.Document.Components.SecuritySchemes["apiKey"].Value.Type)
	assert.Equal(t, "http", result.Document.Components.SecuritySchemes["bearer"].Value.Type)
	assert.Equal(t, "oauth2", result.Document.Components.SecuritySchemes["oauth2"].Value.Type)
}

// --- Mixed Inline and Reference ---

func TestParseRefSecurityScheme_Mixed(t *testing.T) {
	yaml := `openapi: "3.1.0"
info:
  title: "Test"
  version: "1.0"
paths: {}
components:
  securitySchemes:
    inlineScheme:
      type: apiKey
      in: header
      name: X-API-Key
`
	result, err := Parse([]byte(yaml))
	require.NoError(t, err)
	scheme := result.Document.Components.SecuritySchemes["inlineScheme"].Value
	assert.Equal(t, "apiKey", scheme.Type)
	assert.Equal(t, "header", scheme.In)
	assert.Equal(t, "X-API-Key", scheme.Name)
}
